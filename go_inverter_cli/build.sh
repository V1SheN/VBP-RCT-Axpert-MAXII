#!/bin/bash

# Exit immediately if a command exits with a non-zero status.
set -e

# Define variables
IMAGE_NAME="go-inverter-cli"
CONTAINER_NAME="go-inverter-cli"

# Auto-detect DEVICE_PATH by looking for a hidraw device in the 'plugdev' group
echo "Searching for inverter device..."
DEVICE_PATH=$(ls -l /dev/hidraw* 2>/dev/null | grep plugdev | awk '{print $NF}' | head -n 1)

if [ -z "$DEVICE_PATH" ]; then
    echo "Error: No hidraw device with 'plugdev' group found."
    echo "Please check if the inverter is plugged in and udev rules are applied."
    exit 1
fi

echo "--- Detected Inverter Device: $DEVICE_PATH ---"

MQTT_CONFIG_HOST_PATH="./mqtt.json"
MQTT_CONFIG_CONTAINER_PATH="/app/mqtt.json"
POLLING_INTERVAL="5s"
PLATFORM="linux/386"

# Navigate to the script's directory
SCRIPT_DIR=$(dirname "$(readlink -f "$0")")
cd "$SCRIPT_DIR"

# Check if the container is running and stop it
if [ "$(docker ps -q -f name=$CONTAINER_NAME)" ]; then
    echo "Stopping running container: $CONTAINER_NAME"
    docker stop $CONTAINER_NAME
fi

# Check if the container exists and remove it
if [ "$(docker ps -a -q -f name=$CONTAINER_NAME)" ]; then
    echo "Removing existing container: $CONTAINER_NAME"
    docker rm $CONTAINER_NAME
fi

echo "--- Building Docker image: $IMAGE_NAME for platform $PLATFORM ---"
docker build --platform "$PLATFORM" -t "$IMAGE_NAME" .

echo "--- Running Docker container: $IMAGE_NAME ---"
# We mount /dev and use --privileged so the container can start even if a specific
# device is missing. The Go app will then retry and auto-discover the device.
docker run -d --restart always --name "$CONTAINER_NAME" --platform "$PLATFORM" \
  -v /dev:/dev --privileged \
  "$IMAGE_NAME" -interval "$POLLING_INTERVAL" -debug
