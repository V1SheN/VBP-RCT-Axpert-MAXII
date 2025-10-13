package main

import (
	"encoding/json"
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
	flag.Parse()

	devicePath := *devicePtr
	pollingInterval := *intervalPtr

	fmt.Println("Starting Go Inverter CLI...")

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
		// --- QPIGS Command ---
		fmt.Println("\nSending QPIGS command...")
		rawResponse, err := communicator.SendCommand("QPIGS")
		if err != nil {
			fmt.Printf("Error sending QPIGS command: %v\n", err)
		} else {
			fmt.Printf("Raw QPIGS response: %s\n", rawResponse)
			fmt.Printf("--- Raw QPIGS Response (Length: %d) ---\n%s\n--- End Raw QPIGS Response ---\n", len(rawResponse), rawResponse)

			qpigSData, err := parser.ParseQPIGSResponse(rawResponse)
			if err != nil {
				fmt.Printf("Error parsing QPIGS response: %v\n", err)
			} else {
				fmt.Printf("Parsed QPIGS Data: %+v\n", qpigSData)
				jsonData, err := json.MarshalIndent(qpigSData, "", "  ")
				if err != nil {
					fmt.Printf("Error marshaling QPIGS data to JSON: %v\n", err)
				} else {
					fmt.Printf("JSON Data: %s\n", jsonData)
					err = publisher.PublishData(qpigSData)
					if err != nil {
						fmt.Printf("Error publishing QPIGS data to MQTT: %v\n", err)
					}
				}
			}
		}

		// --- QPIRI Command ---
		fmt.Println("\nSending QPIRI command...")
		rawResponse, err = communicator.SendCommand("QPIRI")
		if err != nil {
			fmt.Printf("Error sending QPIRI command: %v\n", err)
		} else {
			fmt.Printf("Raw QPIRI response: %s\n", rawResponse)
			fmt.Printf("--- Raw QPIRI Response (Length: %d) ---\n%s\n--- End Raw QPIRI Response ---\n", len(rawResponse), rawResponse)

			qpiriData, err := parser.ParseQPIRIResponse(rawResponse)
			if err != nil {
				fmt.Printf("Error parsing QPIRI response: %v\n", err)
			} else {
				fmt.Printf("Parsed QPIRI Data: %+v\n", qpiriData)
				jsonData, err := json.MarshalIndent(qpiriData, "", "  ")
				if err != nil {
					fmt.Printf("Error marshaling QPIRI data to JSON: %v\n", err)
				} else {
					fmt.Printf("JSON Data: %s\n", jsonData)
					// Publish to a different topic for QPIRI
					err = publisher.PublishData(qpiriData)
					if err != nil {
						fmt.Printf("Error publishing QPIRI data to MQTT: %v\n", err)
					}
				}
			}
		}

		// --- QPIGS2 Command ---
		fmt.Println("\nSending QPIGS2 command...")
		rawResponse, err = communicator.SendCommand("QPIGS2")
		if err != nil {
			fmt.Printf("Error sending QPIGS2 command: %v\n", err)
		} else {
			fmt.Printf("Raw QPIGS2 response: %s\n", rawResponse)
			fmt.Printf("--- Raw QPIGS2 Response (Length: %d) ---\n%s\n--- End Raw QPIGS2 Response ---\n", len(rawResponse), rawResponse)

			qpigs2Data, err := parser.ParseQPIGS2Response(rawResponse)
			if err != nil {
				fmt.Printf("Error parsing QPIGS2 response: %v\n", err)
			} else {
				fmt.Printf("Parsed QPIGS2 Data: %+v\n", qpigs2Data)
				jsonData, err := json.MarshalIndent(qpigs2Data, "", "  ")
				if err != nil {
					fmt.Printf("Error marshaling QPIGS2 data to JSON: %v\n", err)
				} else {
					fmt.Printf("JSON Data: %s\n", jsonData)
					// Publish to a different topic for QPIGS2
					err = publisher.PublishData(qpigs2Data)
					if err != nil {
						fmt.Printf("Error publishing QPIGS2 data to MQTT: %v\n", err)
					}
				}
			}
		}

		// --- QPIWS Command ---
		fmt.Println("\nSending QPIWS command...")
		rawResponse, err = communicator.SendCommand("QPIWS")
		if err != nil {
			fmt.Printf("Error sending QPIWS command: %v\n", err)
		} else {
			fmt.Printf("Raw QPIWS response: %s\n", rawResponse)
			fmt.Printf("--- Raw QPIWS Response (Length: %d) ---\n%s\n--- End Raw QPIWS Response ---\n", len(rawResponse), rawResponse)

			qpiwsData, err := parser.ParseQPIWSResponse(rawResponse)
			if err != nil {
				fmt.Printf("Error parsing QPIWS response: %v\n", err)
			} else {
				fmt.Printf("Parsed QPIWS Data: %+v\n", qpiwsData)
				jsonData, err := json.MarshalIndent(qpiwsData, "", "  ")
				if err != nil {
					fmt.Printf("Error marshaling QPIWS data to JSON: %v\n", err)
				} else {
					fmt.Printf("JSON Data: %s\n", jsonData)
					// Publish to a different topic for QPIWS
					err = publisher.PublishData(qpiwsData)
					if err != nil {
						fmt.Printf("Error publishing QPIWS data to MQTT: %v\n", err)
					}
				}
			}
		}

		time.Sleep(pollingInterval) // Wait for the next poll
	}
			}
			