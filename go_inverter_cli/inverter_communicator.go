package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

const (
	bufferSize = 256 // Increased buffer size to accommodate longer responses
)

// InverterCommunicator handles low-level communication with the inverter device.
type InverterCommunicator struct {
	deviceFile *os.File
	devicePath string
}

// NewInverterCommunicator creates a new communicator instance.
func NewInverterCommunicator(path string) *InverterCommunicator {
	return &InverterCommunicator{
		devicePath: path,
	}
}

// OpenDevice opens the hidraw device file.
func (ic *InverterCommunicator) OpenDevice() error {
	var err error
	ic.deviceFile, err = os.OpenFile(ic.devicePath, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("error opening device %s: %w", ic.devicePath, err)
	}
	return nil
}

// CloseDevice closes the hidraw device file.
func (ic *InverterCommunicator) CloseDevice() error {
	if ic.deviceFile != nil {
		return ic.deviceFile.Close()
	}
	return nil
}

func calculateCRC(data []byte) []byte {
	var crc uint16 = 0x0000
	for _, b := range data {
		crc ^= uint16(b) << 8
		for i := 0; i < 8; i++ {
			if (crc&0x8000) != 0 {
				crc = (crc << 1) ^ 0x1021
			} else {
				crc <<= 1
			}
		}
	}
	return []byte{byte((crc >> 8) & 0xFF), byte(crc & 0xFF)}
}

// SendCommand sends a command to the inverter and reads its response.
func (ic *InverterCommunicator) SendCommand(command string) (string, error) {
	if ic.deviceFile == nil {
		return "", fmt.Errorf("device not open")
	}

	// --- Pre-read Flush (Aggressive Best Effort) ---
	// Repeatedly read and discard any lingering data until no more is found,
	// or a short overall flush timeout is reached.
	flushBuf := make([]byte, bufferSize)
	flushStartTime := time.Now()
	flushedBytes := 0 // Add counter
	for {
		if time.Since(flushStartTime) > 200*time.Millisecond { // Overall flush timeout
			fmt.Printf("Communicator: Pre-read flush timeout reached after %dms, %d bytes flushed.\n", time.Since(flushStartTime).Milliseconds(), flushedBytes)
			break
		}
		_ = ic.deviceFile.SetReadDeadline(time.Now().Add(20 * time.Millisecond)) // Short deadline for each read
		n, err := ic.deviceFile.Read(flushBuf)
		_ = ic.deviceFile.SetReadDeadline(time.Time{})      // Clear the deadline

		if err != nil {
			// If it's a timeout error, assume buffer is clear
			if os.IsTimeout(err) {
				fmt.Printf("Communicator: Pre-read flush completed (read timeout). %d bytes flushed.\n", flushedBytes)
				break
			}
			// Other errors, just break
			fmt.Printf("Communicator: Error during pre-read flush: %v. %d bytes flushed.\n", err, flushedBytes)
			break
		}
		if n == 0 {
			// No data read, assume buffer is clear
			fmt.Printf("Communicator: Pre-read flush completed (no data). %d bytes flushed.\n", flushedBytes)
			break
		}
		flushedBytes += n // Increment counter
		// Data was read, continue flushing
	}

	// Commands need to be terminated with CRC and a carriage return (CR)
	cmdWithCRC := []byte(command)
	crc := calculateCRC(cmdWithCRC)
	cmdBytes := append(cmdWithCRC, crc...)
	cmdBytes = append(cmdBytes, '\r')

	// Write the command
	fmt.Printf("Communicator: Sending command '%s' (bytes: %x)\n", command, cmdBytes)
	_, err := ic.deviceFile.Write(cmdBytes)
	if err != nil {
		return "", fmt.Errorf("error writing command to device: %w", err)
	}

	fmt.Println("Communicator: Command sent. Waiting for response...")
	responseBuffer := make([]byte, bufferSize)
	var response []byte

	readStartTime := time.Now()
	for {
		// This inner timeout is not perfectly reliable with blocking I/O, 
		// but it's a safety net. The main timeout is now in main.go
		if time.Since(readStartTime) > 5*time.Second { 
			return "", fmt.Errorf("internal read timeout after 5 seconds")
		}

		n, err := ic.deviceFile.Read(responseBuffer)
		if err != nil {
			if err == io.EOF {
				return "", fmt.Errorf("EOF reached while reading response: %w", err)
			} else {
				return "", fmt.Errorf("error reading from device: %w", err)
			}
		}

		if n > 0 {
			fmt.Printf("Communicator: Read %d bytes: %x\n", n, responseBuffer[:n])
			response = append(response, responseBuffer[:n]...)
			fmt.Printf("Communicator: Current response buffer (hex): %x, length: %d\n", response, len(response))

			crIndex := bytes.IndexByte(response, '\r')
			if crIndex != -1 { // If CR is found
				fullResponseWithCR := response[:crIndex+1]

				if len(fullResponseWithCR) < 3 {
					return "", fmt.Errorf("response too short to contain data, CRC, and CR")
			}

				responseWithoutCR := fullResponseWithCR[:len(fullResponseWithCR)-1]
			receivedCRC := responseWithoutCR[len(responseWithoutCR)-2:]
				dataPart := responseWithoutCR[:len(responseWithoutCR)-2]
				calculatedCRC := calculateCRC(dataPart)

				if !bytes.Equal(receivedCRC, calculatedCRC) {
					return "", fmt.Errorf("CRC mismatch: received %x, calculated %x for data %s (hex: %x)", receivedCRC, calculatedCRC, dataPart, dataPart)
				}

				for len(dataPart) > 0 && (dataPart[len(dataPart)-1] == 0 || dataPart[len(dataPart)-1] < 32) {
					dataPart = dataPart[:len(dataPart)-1]
				}
				return string(dataPart), nil
			}
		}
		// Prevent busy-waiting if device is slow or sends partial responses
		time.Sleep(50 * time.Millisecond)
	}
}
