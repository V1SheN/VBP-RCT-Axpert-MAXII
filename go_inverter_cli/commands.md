# Go Inverter CLI - Commands Log

This document logs the essential commands used for building, running, and testing the Go Inverter CLI application.

## Automation Script

### Build and Run with `build.sh`
This script automates the Docker image build and container execution process.

```bash
./build.sh
```

**Note:** Ensure you are in the `go_inverter_cli` directory when running this script, or provide the full path to `build.sh`.

## Build Commands

### Build Docker Image (for linux/386 architecture)
```bash
docker build --platform linux/386 -t go-inverter-cli .
```

## Run Commands

### Run Docker Container (using the automatic build script)
The easiest way to run the container is using the provided `build.sh` script, which automatically detects the correct `/dev/hidrawX` device:
```bash
./build.sh
```

### Run Docker Container Manually
If you need to run it manually, replace `/dev/hidrawX` with your actual device path:
```bash
docker run --rm -it --platform linux/386 --device=/dev/hidrawX -v $(pwd)/mqtt.json:/app/mqtt.json go-inverter-cli -device /dev/hidrawX -interval 5s
```

## Testing Commands

### Subscribe to MQTT Topic (using mosquitto_sub)
```bash
mosquitto_sub -h 192.168.31.243 -p 1883 -u mqqt-user -P Venita69 -t homeassistant/voltronic/state -v
```
