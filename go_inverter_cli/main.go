package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Command-line arguments
	devicePtr := flag.String("device", "/dev/hidraw4", "Path to the hidraw device")
	intervalPtr := flag.Duration("interval", 2*time.Second, "Polling interval (e.g., 2s, 1m)")
	debugPtr := flag.Bool("debug", false, "Enable debug mode to query extra commands")
	flag.Parse()

	devicePath := *devicePtr
	pollingInterval := *intervalPtr
	debugMode := *debugPtr

	fmt.Println("Starting Go Inverter CLI...")
	if debugMode {
		fmt.Println("** DEBUG MODE ENABLED **")
	}

	// Initialize communicator and parser
	communicator := NewInverterCommunicator(devicePath)
	parser := NewInverterParser()

	// Open the device
	err := communicator.OpenDevice()
	if err != nil {
		fmt.Printf("Failed to open device: %v\n", err)
		if os.IsNotExist(err) {
			fmt.Println("Device not found. Make sure it exists.")
		} else if os.IsPermission(err) {
			fmt.Println("Permission denied. Try running with appropriate permissions (e.g., sudo) or check udev rules.")
		} else {
			fmt.Println("An unexpected error occurred during device opening.")
		}
		os.Exit(1)
	}
	defer communicator.CloseDevice()

	fmt.Printf("Successfully opened %s.\n", devicePath)

	// Setup signal handling for graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("\nReceived interrupt signal. Closing device and exiting.")
		communicator.CloseDevice()
		os.Exit(0)
	}()

	// Load MQTT Configuration
	mqttConfig, err := LoadMQTTConfig("/app/mqtt.json")
	if err != nil {
		fmt.Printf("Failed to load MQTT configuration: %v\n", err)
		os.Exit(1)
	}

	// Initialize MQTT Publisher
	publisher := NewMQTTPublisher(mqttConfig)
	err = publisher.Connect()
	if err != nil {
		fmt.Printf("Failed to connect to MQTT broker: %v\n", err)
		os.Exit(1)
	}
	defer publisher.Disconnect()

	// Main polling loop
	for {
		// Define a struct to hold command results
		type CommandResult struct {
			Response string
			Err      error
		}

		// --- QPIGS Command ---
		fmt.Println("\nSending QPIGS command...")
		cmdChanQPIGS := make(chan CommandResult, 1)
		go func() {
			rawResponse, err := communicator.SendCommand("QPIGS")
			cmdChanQPIGS <- CommandResult{Response: rawResponse, Err: err}
		}()

		select {
		case res := <-cmdChanQPIGS:
			if res.Err != nil {
				fmt.Printf("Error sending QPIGS command: %v\n", res.Err)
			} else {
				qpigSData, err := parser.ParseQPIGSResponse(res.Response)
				if err != nil {
					fmt.Printf("Error parsing QPIGS response: %v\n", err)
				} else {
					fmt.Printf("Parsed QPIGS Data: %+v\n", qpigSData)
					err = publisher.PublishData(qpigSData, "state")
					if err != nil {
						fmt.Printf("Error publishing QPIGS data to MQTT: %v\n", err)
					}
				}
			}
		case <-time.After(2 * time.Second):
			fmt.Println("Error: Timeout waiting for QPIGS response after 2 seconds")
		}

		time.Sleep(300 * time.Millisecond)

		// --- QPIRI Command ---
		fmt.Println("\nSending QPIRI command...")
		cmdChanQPIRI := make(chan CommandResult, 1)
		go func() {
			rawResponse, err := communicator.SendCommand("QPIRI")
			cmdChanQPIRI <- CommandResult{Response: rawResponse, Err: err}
		}()

		select {
		case res := <-cmdChanQPIRI:
			if res.Err != nil {
				fmt.Printf("Error sending QPIRI command: %v\n", res.Err)
			} else {
				qpiriData, err := parser.ParseQPIRIResponse(res.Response)
				if err != nil {
					fmt.Printf("Error parsing QPIRI response: %v\n", err)
				} else {
					fmt.Printf("Parsed QPIRI Data: %+v\n", qpiriData)
					err = publisher.PublishData(qpiriData, "rating")
					if err != nil {
						fmt.Printf("Error publishing QPIRI data to MQTT: %v\n", err)
					}
				}
			}
		case <-time.After(2 * time.Second):
			fmt.Println("Error: Timeout waiting for QPIRI response after 2 seconds")
		}

		time.Sleep(300 * time.Millisecond)

		        // --- QPIGS2 Command ---
				fmt.Println("\nSending QPIGS2 command...")
				cmdChanQPIGS2 := make(chan CommandResult, 1)
				go func() {
					rawResponse, err := communicator.SendCommand("QPIGS2")
					cmdChanQPIGS2 <- CommandResult{Response: rawResponse, Err: err}
				}()
		
				select {
				case res := <-cmdChanQPIGS2:
					if res.Err != nil {
						fmt.Printf("Error sending QPIGS2 command: %v\n", res.Err)
					} else {
						qpigs2Data, err := parser.ParseQPIGS2Response(res.Response)
						if err != nil {
							fmt.Printf("Error parsing QPIGS2 response: %v\n", err)
						} else {
							fmt.Printf("Parsed QPIGS2 Data: %+v\n", qpigs2Data)
							err = publisher.PublishData(qpigs2Data, "pv2")
							if err != nil {
								fmt.Printf("Error publishing QPIGS2 data to MQTT: %v\n", err)
							}
						}
					}
				case <-time.After(2 * time.Second):
					fmt.Println("Error: Timeout waiting for QPIGS2 response after 2 seconds")
				}
		time.Sleep(300 * time.Millisecond)

		// --- QPIWS Command ---
		fmt.Println("\nSending QPIWS command...")
		cmdChanQPIWS := make(chan CommandResult, 1)
		go func() {
			rawResponse, err := communicator.SendCommand("QPIWS")
			cmdChanQPIWS <- CommandResult{Response: rawResponse, Err: err}
		}()

		select {
		case res := <-cmdChanQPIWS:
			if res.Err != nil {
				fmt.Printf("Error sending QPIWS command: %v\n", res.Err)
			} else {
				qpiwsData, err := parser.ParseQPIWSResponse(res.Response)
				if err != nil {
					fmt.Printf("Error parsing QPIWS response: %v\n", err)
				} else {
					fmt.Printf("Parsed QPIWS Data: %+v\n", qpiwsData)
					err = publisher.PublishData(qpiwsData, "warnings")
					if err != nil {
						fmt.Printf("Error publishing QPIWS data to MQTT: %v\n", err)
					}
				}
			}
		case <-time.After(2 * time.Second):
			fmt.Println("Error: Timeout waiting for QPIWS response after 2 seconds")
		}

		// --- Debug Commands ---
		if debugMode {
			time.Sleep(300 * time.Millisecond)

			// --- QMOD Command ---
			fmt.Println("\n[DEBUG] Sending QMOD command...")
			cmdChanQMOD := make(chan CommandResult, 1)
			go func() {
				rawResponse, err := communicator.SendCommand("QMOD")
				cmdChanQMOD <- CommandResult{Response: rawResponse, Err: err}
			}()

			select {
			case res := <-cmdChanQMOD:
				if res.Err != nil {
					fmt.Printf("Error sending QMOD command: %v\n", res.Err)
				} else {
					qmodData, err := parser.ParseQMODResponse(res.Response)
					if err != nil {
						fmt.Printf("Error parsing QMOD response: %v\n", err)
					} else {
						fmt.Printf("Parsed QMOD Data: %+v\n", qmodData)
						err = publisher.PublishData(qmodData, "debug/qmod")
						if err != nil {
							fmt.Printf("Error publishing QMOD data to MQTT: %v\n", err)
						}
					}
				}
			case <-time.After(2 * time.Second):
				fmt.Println("Error: Timeout waiting for QMOD response after 2 seconds")
			}

			time.Sleep(300 * time.Millisecond)

			// --- QDI Command ---
			fmt.Println("\n[DEBUG] Sending QDI command...")
			cmdChanQDI := make(chan CommandResult, 1)
			go func() {
				rawResponse, err := communicator.SendCommand("QDI")
				cmdChanQDI <- CommandResult{Response: rawResponse, Err: err}
			}()

			select {
			case res := <-cmdChanQDI:
				if res.Err != nil {
					fmt.Printf("Error sending QDI command: %v\n", res.Err)
				} else {
					qdiData, err := parser.ParseQDIResponse(res.Response)
					if err != nil {
						fmt.Printf("Error parsing QDI response: %v\n", err)
					} else {
						fmt.Printf("Parsed QDI Data: %+v\n", qdiData)
						err = publisher.PublishData(qdiData, "debug/qdi")
						if err != nil {
							fmt.Printf("Error publishing QDI data to MQTT: %v\n", err)
						}
					}
				}
			case <-time.After(2 * time.Second):
				fmt.Println("Error: Timeout waiting for QDI response after 2 seconds")
			}
		}

		time.Sleep(pollingInterval) // Wait for the next poll
	}
}	