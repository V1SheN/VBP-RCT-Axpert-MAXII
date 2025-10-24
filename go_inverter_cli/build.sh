#!/bin/bash

# Exit immediately if a command exits with a non-zero status.
set -e

# Define variables
IMAGE_NAME="go-inverter-cli"
CONTAINER_NAME="go-inverter-cli"
DEVICE_PATH="/dev/hidraw5"
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
docker run -d --rm --name go-inverter-cli --platform "$PLATFORM" --device="$DEVICE_PATH" \
  "$IMAGE_NAME" -device "$DEVICE_PATH" -interval "$POLLING_INTERVAL" -debug
