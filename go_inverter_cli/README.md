# Go Inverter CLI

This application is a command-line interface (CLI) for monitoring and controlling Voltronic Axpert MAX II inverters. It communicates with the inverter via a HID device, polls for data, and publishes it to an MQTT broker.

## Features

- Polls the inverter for general status (`QPIGS`), rating information (`QPIRI`), and warning status (`QPIWS`).
- Publishes inverter data to an MQTT broker.
- Allows sending commands to the inverter via MQTT.

## Configuration

The application is configured using the `mqtt.json` file. This file should be placed in the same directory as the executable.

```json
{
    "server": "<your_mqtt_broker_ip>",
    "port": "1883",
    "topic": "homeassistant",
    "devicename": "voltronic",
    "username": "<your_mqtt_username>",
    "password": "<your_mqtt_password>",
    "clientid": "voltronic_cli"
}
```

## Usage

1.  **Build the application:**
    ```bash
    go build
    ```

2.  **Run the application:**
    ```bash
    ./go_inverter_cli
    ```

## MQTT Commands

You can send commands to the inverter by publishing messages to a specific MQTT topic.

-   **Command Topic:** `homeassistant/voltronic/cmd` (This is constructed from the `topic` and `devicename` in your `mqtt.json` file).
-   **Payload Format:** The message payload should be a JSON object with a `command` and a `value` field.

### Set Output Source Priority (POP)

To set the output source priority, use the `pop` command.

-   **Command:** `pop`
-   **Values:**
    -   `uti`: Utility first (Utility -> Solar -> Battery)
    -   `sol`: Solar first (Solar -> Utility -> Battery)
    -   `sbu`: SBU (Solar -> Battery -> Utility)

**Example:**

To set the output source priority to SBU, publish the following message to the command topic:

```bash
mosquitto_pub -h <your_mqtt_broker_ip> -t "homeassistant/voltronic/cmd" -m '{"command": "pop", "value": "sbu"}'
```

The result of the command will be published to the `homeassistant/voltronic/cmd/result` topic.
