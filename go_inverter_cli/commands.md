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

### Run Docker Container (with device and MQTT config mounted to /app/mqtt.json)
```bash
docker run --rm -it --platform linux/386 --device=/dev/hidraw4 -v /home/fish/Software/Development/github/Home-Assistant/docker-voltronic-homeassistant-master/config/mqtt.json:/app/mqtt.json go-inverter-cli -device /dev/hidraw4 -interval 5s
```

## Testing Commands

### Subscribe to MQTT Topic (using mosquitto_sub)
```bash
mosquitto_sub -h 192.168.31.243 -p 1883 -u mqqt-user -P Venita69 -t homeassistant/voltronic/state -v
```
