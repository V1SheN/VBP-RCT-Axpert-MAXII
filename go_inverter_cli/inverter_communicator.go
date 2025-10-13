package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

const ( 
	// devicePath = "/dev/hidraw4" // This will eventually come from config
	bufferSize = 64 // Typical HID report size
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
			if (crc & 0x8000) != 0 {
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
	// Read the response
	// Inverter responses are typically terminated with a carriage return (CR)
	// and sometimes a line feed (LF). We'll read until CR.
	responseBuffer := make([]byte, bufferSize)
	var response []byte

	readStartTime := time.Now()
	for {
		// Add a timeout for reading the full response to prevent infinite loops
		if time.Since(readStartTime) > 5*time.Second { // 5-second timeout for response
			return "", fmt.Errorf("response read timeout after 5 seconds")
		}

		n, err := ic.deviceFile.Read(responseBuffer)
		if err != nil {
			if err == io.EOF {
				return "", fmt.Errorf("EOF reached while reading response: %w", err)
			} else if os.IsTimeout(err) {
				// This might happen if the device is slow, but we have a higher-level timeout
				fmt.Println("Communicator: Read operation timed out (no data yet).")
				time.Sleep(100 * time.Millisecond) // Wait a bit before retrying read
				continue
			} else {
				return "", fmt.Errorf("error reading from device: %w", err)
			}
		}

					if n > 0 {
					fmt.Printf("Communicator: Read %d bytes: %x\n", n, responseBuffer[:n])
								response = append(response, responseBuffer[:n]...)
								fmt.Printf("Communicator: Current response buffer (hex): %x, length: %d\n", response, len(response))
					
								// Search for carriage return to signify end of response
								crIndex := bytes.IndexByte(response, '\r')
								if crIndex != -1 { // If CR is found
									// The full response is up to and including the CR
									fullResponseWithCR := response[:crIndex+1]
									fmt.Printf("Communicator: Full response received (hex): %x\n", fullResponseWithCR)
					
									// The response should be DATA<CRC_LOW><CRC_HIGH><CR>
									// We need at least 3 bytes for CRC (2) and CR (1)
									if len(fullResponseWithCR) < 3 {
										return "", fmt.Errorf("response too short to contain data, CRC, and CR")
									}
					
									// Remove the trailing CR
									responseWithoutCR := fullResponseWithCR[:len(fullResponseWithCR)-1]
									fmt.Printf("Communicator: Response without CR (hex): %x\n", responseWithoutCR)
					
									// Extract CRC (last two bytes before CR)
									if len(responseWithoutCR) < 2 {
										return "", fmt.Errorf("response too short to contain data and CRC")
									}
									receivedCRC := responseWithoutCR[len(responseWithoutCR)-2:]
									dataPart := responseWithoutCR[:len(responseWithoutCR)-2]
									fmt.Printf("Communicator: Data part (hex): %x, Received CRC (hex): %x\n", dataPart, receivedCRC)
					
									// Calculate CRC for the data part
									calculatedCRC := calculateCRC(dataPart)
									fmt.Printf("Communicator: Calculated CRC (hex): %x\n", calculatedCRC)
					
									if !bytes.Equal(receivedCRC, calculatedCRC) {
										return "", fmt.Errorf("CRC mismatch: received %x, calculated %x for data %s (hex: %x)", receivedCRC, calculatedCRC, dataPart, dataPart)
									}
					
									// Trim any trailing null bytes or non-printable characters from the dataPart
									// This is done after CRC validation to ensure CRC was calculated on the correct data.
									for len(dataPart) > 0 && (dataPart[len(dataPart)-1] == 0 || dataPart[len(dataPart)-1] < 32) {
										dataPart = dataPart[:len(dataPart)-1]
									}
									fmt.Printf("Communicator: Cleaned data part (hex): %x\n", dataPart)
					
									// Return the data part as a string
									return string(dataPart), nil
								}
		}
		// Prevent busy-waiting if device is slow or sends partial responses
		time.Sleep(10 * time.Millisecond)
	}
}
