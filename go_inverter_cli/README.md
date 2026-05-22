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
    "clientid": "voltronic_cli",
    "devices": ["/dev/hidraw4", "/dev/hidraw5", "/dev/hidraw6", "/dev/hidraw7", "/dev/hidraw8"]
}
```

The application will attempt to open the devices in the `devices` list sequentially. If none are found or work, it will attempt auto-discovery of any `hidraw` device belonging to the `plugdev` group. If no device is found, it will retry every 30 seconds.

## Usage

1.  **Deploy to Raspberry Pi:**
    From your development machine, run:
    ```bash
    rsync -avz /home/fish/Software/Development/github/VBP-RCT-Axpert-MAXII/go_inverter_cli/ fish@192.168.31.218:/opt/go_inverter_cli/
    ```

2.  **Build and Run (on Raspberry Pi):**
    SSH into your Raspberry Pi and run:
    ```bash
    cd /opt/go_inverter_cli/ && sudo ./build.sh
    ```

## Local Development (Optional)

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

when POP00 send {"command": "pop", "value": "uti"}
 when POP01 send {"command": "pop", "value": "sol"}
when pop02 send {"command": "pop", "value": "sbu"}
**Example:**

To set the output source priority to SBU, publish the following message to the command topic:

```bash
mosquitto_pub -h <your_mqtt_broker_ip> -t "homeassistant/voltronic/cmd" -m '{"command": "pop", "value": "sbu"}'
```

### Set Charger Source Priority (PCP)

To set the charger source priority, use the `pcp` command.

-   **Command:** `pcp`
-   **Values:**
    -   `uti`: Utility first
    -   `sol`: Solar first
    -   `soluti`: Solar and Utility
    -   `onlysol`: Only Solar

**Example:**

To set the charger source priority to Solar and Utility, publish the following message to the command topic:

```bash
mosquitto_pub -h <your_mqtt_broker_ip> -t "homeassistant/voltronic/cmd" -m '{"command": "pcp", "value": "soluti"}'
```

The result of the command will be published to the `homeassistant/voltronic/cmd/result` topic.
