#!/bin/bash

# Exit immediately if a command exits with a non-zero status.
set -e

# Define variables
IMAGE_NAME="go-inverter-cli"
DEVICE_PATH="/dev/hidraw5"
MQTT_CONFIG_HOST_PATH="./mqtt.json"
MQTT_CONFIG_CONTAINER_PATH="/app/mqtt.json"
POLLING_INTERVAL="5s"
PLATFORM="linux/386"

# Navigate to the script's directory
SCRIPT_DIR=$(dirname "$(readlink -f "$0")")
cd "$SCRIPT_DIR"

echo "--- Building Docker image: $IMAGE_NAME for platform $PLATFORM ---"
docker build --platform "$PLATFORM" -t "$IMAGE_NAME" .

echo "--- Running Docker container: $IMAGE_NAME ---"
docker run --rm --platform "$PLATFORM" --device="$DEVICE_PATH" \
  "$IMAGE_NAME" -device "$DEVICE_PATH" -interval "$POLLING_INTERVAL" -debug
