# High-Level Plan: Go Inverter CLI for Voltronic Inverters

## Objective
To develop a Go-based command-line interface (CLI) application that replicates the core functionality of the existing C++ inverter-cli. This includes establishing communication with a Voltronic inverter via `/dev/hidrawX`, polling for various data points, parsing the raw responses according to the Axpert protocol, and publishing the structured data to an MQTT broker. The application will be containerized using Docker, specifically targeting a `linux/386` host architecture.

## Key Features
*   **Device Communication:** Open and maintain a connection to `/dev/hidrawX` (e.g., `/dev/hidraw4`).
*   **Command Polling:** Implement functions to send specific inquiry commands (e.g., `QPIGS`, `QPIRI`, `QPIGS2`, `QPIWS`) to the inverter.
*   **Data Parsing:** Parse raw string responses from the inverter into structured Go data types, interpreting values based on `data_format.md` and `axpert_protocol.pdf`.
*   **Data Structuring:** Organize parsed data into a coherent data model, suitable for JSON serialization.
*   **MQTT Publishing:** Connect to an MQTT broker (configured via `config/mqtt.json`) and publish the structured inverter data.
*   **Configuration Management:** Load MQTT and potentially inverter-specific settings from configuration files.
*   **Robustness:** Implement comprehensive error handling for device access, communication, parsing, and MQTT operations.
*   **Graceful Shutdown:** Ensure proper resource cleanup and application termination upon signals.
*   **Containerization:** Deployable as a Docker container for `linux/386` architecture.

## Technology Stack
*   **Language:** Go (Golang)
*   **Containerization:** Docker
*   **Device Interaction:** Standard Go `os` package for file I/O (for `/dev/hidrawX`).
*   **MQTT Client:** A Go MQTT client library (e.g., `paho.mqtt.golang`).
*   **JSON Handling:** Standard Go `encoding/json` package.
*   **Configuration:** Standard Go `encoding/json` or a dedicated config library.
*   **Base Image (Build):** `golang:1.21`
*   **Base Image (Runtime):** `alpine:latest`
*   **Target Architecture:** `linux/386`

## Architecture Overview
1.  **`main.go`:** Main application entry point, orchestrating polling, parsing, and publishing.
2.  **`inverter_communicator.go`:** Handles low-level device interaction (sending commands, reading responses).
3.  **`inverter_parser.go`:** Contains logic to parse raw inverter responses into structured Go types.
4.  **`mqtt_publisher.go`:** Manages connection to MQTT broker and publishing of data.
5.  **`config_loader.go`:** Reads and provides application configuration (MQTT settings, etc.).
6.  **Dockerfile:** Multi-stage build for `linux/386` to create a minimal, self-contained container.

## Phases
1.  **Core Communication & Parsing:** Refine device communication, implement command sending, and initial data parsing.
2.  **Data Structuring & MQTT Integration:** Develop data models, integrate MQTT client, and publish parsed data.
3.  **Configuration & Robustness:** Implement configuration loading, enhance error handling, and add logging.
4.  **Containerization & Deployment:** Finalize Dockerfile, build, and validate container deployment.

## Success Criteria
*   The Go application successfully builds for `linux/386` within Docker.
*   The Docker container starts and runs without errors.
*   The Go application inside the container can open `/dev/hidrawX`.
*   The application can send commands to the inverter and receive responses.
*   Raw inverter responses are correctly parsed into structured Go data.
*   Structured data is successfully published to the configured MQTT broker.
*   Application configuration is loaded from `config/mqtt.json`.
*   Robust error handling and graceful shutdown are implemented.