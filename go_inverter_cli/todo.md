# To-Do List: Go Inverter CLI for Voltronic Inverters

This document outlines the tasks required to develop and deploy the Go-based inverter CLI application.

## Phase 1: Core Communication & Parsing

- [x] Create `go_inverter_cli` directory.
- [x] Create `main.go` file (initial device test).
- [x] Implement device opening logic for `/dev/hidrawX`.
- [x] Implement continuous data reading loop (initial test).
- [x] Implement data printing (hexadecimal format) (initial test).
- [x] Add basic error handling (e.g., device not found, read errors) (initial test).
- [x] Add graceful shutdown on `SIGINT` (Ctrl+C) (initial test).
- [x] Refactor `main.go` to separate concerns into `inverter_communicator.go` and `inverter_parser.go`.
- [x] Implement `inverter_communicator.go`:
    - [x] Function to send commands to the inverter.
    - [x] Function to read responses from the inverter.
    - [x] Implement CRC calculation for commands (sending).
    - [x] Implement CRC validation for responses (receiving).
- [x] Refine response parsing in `SendCommand` to handle trailing null bytes/padding.
- [x] Implement `inverter_parser.go`:
    - [x] Define Go structs for `QPIGS` data based on `data_format.md` and `axpert_protocol.pdf`.
    - [x] Function to parse raw `QPIGS` response into `QPIGS` struct.

## Phase 2: Data Structuring & MQTT Integration

- [x] Implement `inverter_parser.go` (continued):
    - [x] Define Go structs for `QPIRI` data.
    - [x] Function to parse raw `QPIRI` response into `QPIRI` struct.
    - [x] Define Go structs for `QPIGS2` data.
    - [x] Function to parse raw `QPIGS2` response into `QPIGS2` struct.
    - [x] Define Go structs for `QPIWS` data.
    - [x] Function to parse raw `QPIWS` response into `QPIWS` struct.
- [x] Implement `mqtt_publisher.go`:
    - [x] Integrate `paho.mqtt.golang` library.
    - [x] Function to connect to MQTT broker.
    - [x] Function to publish structured inverter data (JSON) to MQTT topic.

## Phase 3: Configuration & Robustness

- [x] Implement `config_loader.go`:
    - [x] Function to read `config/mqtt.json`.
    - [x] Define Go struct for MQTT configuration.
- [x] Integrate configuration into MQTT client.
- [x] Enhance error handling and logging across all modules.
- [x] Implement continuous polling loop in `main.go` with configurable interval.
- [x] Add command-line argument parsing (e.g., for device path, polling interval, run-once mode).

## Phase 4: Containerization & Deployment

- [x] Create `Dockerfile` (initial version).
- [x] Implement multi-stage build in Dockerfile (initial version).
- [x] Provide `docker build` command with `--platform linux/386`.
- [x] Provide `docker run` command with `--device=/dev/hidrawX` and `--platform linux/386`.
- [x] Verify Docker image builds successfully for `linux/386`.
- [x] Verify Docker container runs successfully.
- [x] Verify Go application inside container can access and read data from `/dev/hidrawX`.
- [x] Observe and confirm data output in container logs (initial test).
- [x] Test graceful shutdown of the container (initial test).
- [x] Update `Dockerfile` to include new Go modules and dependencies.
- [x] Build and test final Docker image.
- [x] Validate end-to-end functionality (polling, parsing, MQTT publishing) in Docker container.

## Phase 5: MQTT Stability & Output Enhancements

- [x] Make MQTT ClientID dynamic to prevent collisions.
- [x] Display the JSON string of the data before publishing to MQTT.