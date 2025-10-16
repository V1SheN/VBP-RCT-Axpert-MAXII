# Live Data Mapping and Processing Analysis

This document provides a mapping between the data fields processed by the Go Inverter CLI application (specifically for the QPIGS, QPIRI, QPIGS2, and QPIWS commands) and the corresponding MQTT sensor entities defined for Home Assistant. It also assesses whether 100% of the data received from the inverter is being processed.

## Go Inverter CLI (Structs)

The following structs, as defined in `inverter_parser.go`, represent the parsed data received from the inverter via their respective commands:

### QPIGSData Struct
```go
type QPIGSData struct {
	GridVoltage                 float64
	GridFrequency               float64
	ACOutputVoltage             float64
	ACOutputFrequency           float64
	ACOutputApparentPower       int
	ACOutputActivePower         int
	OutputLoadPercent           int
	BUSVoltage                  int
	BatteryVoltage              float64
	BatteryChargingCurrent      int
	BatteryCapacity             int
	InverterHeatSinkTemp        int
	PV1InputCurrent             float64
	PV1InputVoltage             float64
	BatteryVoltageFromSCC       float64
	BatteryDischargeCurrent     int
	DeviceStatus1               string // Raw bit string for now
	BatteryVOffsetForFansOn     int
	EEPROMVersion               int
	PV1ChargingPower            int
	DeviceStatus2               string // Raw bit string for now
}
```

### QPIRIData Struct
```go
type QPIRIData struct {
	GridRatingVoltage           float64
	GridRatingCurrent           float64
	ACOutputRatingVoltage       float64
	ACOutputRatingFrequency     float64
	ACOutputRatingCurrent       float64
	ACOutputRatingApparentPower int
	ACOutputRatingActivePower   int
	BatteryRatingVoltage        float64
	BatteryRechargeVoltage      float64
	BatteryUnderVoltage         float64
	BatteryBulkVoltage          float64
	BatteryFloatVoltage         float64
	BatteryType                 int
	MaxACChargingCurrent        int
	MaxChargingCurrent          int
	InputVoltageRange           int
	OutputSourcePriority        int
	ChargerSourcePriority       int
	ParallelMaxNumber           int
	MachineType                 int
	Topology                    int
	OutputMode                  int
	BatteryRedischargeVoltage   float64
	PVOKConditionForParallel    int
	PVPowerBalance              int
	MaxChargingTimeAtCVStage    int
	OperationLogic              int
	MaxDischargingCurrent       int
}
```

### QPIGS2Data Struct
```go
type QPIGS2Data struct {
	PV2InputCurrent  float64
	PV2InputVoltage  float64
	PV2ChargingPower int
}
```

### QPIWSData Struct
```go
type QPIWSData struct {
	WarningFlags string // Raw bit string for now
}
```

## Home Assistant MQTT Sensors

The following files define MQTT sensor entities in Home Assistant, each subscribing to a dedicated topic and extracting specific data points from the JSON payload published by the Go application.

### `mqtt_sensors.yaml` (for QPIGS Data)

| QPIGSData Field             | Home Assistant Sensor Name              | `value_template` Mapping              |
| :-------------------------- | :-------------------------------------- | :------------------------------------ |
| `GridVoltage`               | `Inverter Grid Voltage`                 | `value_json.GridVoltage`              |
| `GridFrequency`             | `Inverter Grid Frequency`               | `value_json.GridFrequency`            |
| `ACOutputVoltage`           | `Inverter AC Output Voltage`            | `value_json.ACOutputVoltage`          |
| `ACOutputFrequency`         | `Inverter AC Output Frequency`          | `value_json.ACOutputFrequency`        |
| `ACOutputApparentPower`     | `Inverter AC Output Apparent Power`     | `value_json.ACOutputApparentPower`    |
| `ACOutputActivePower`       | `Inverter AC Output Active Power`       | `value_json.ACOutputActivePower`      |
| `OutputLoadPercent`         | `Inverter Output Load Percent`          | `value_json.OutputLoadPercent`        |
| `BUSVoltage`                | `Inverter BUS Voltage`                  | `value_json.BUSVoltage`               |
| `BatteryVoltage`            | `Inverter Battery Voltage`              | `value_json.BatteryVoltage`           |
| `BatteryChargingCurrent`    | `Inverter Battery Charging Current`     | `value_json.BatteryChargingCurrent`   |
| `BatteryCapacity`           | `Inverter Battery Capacity`             | `value_json.BatteryCapacity`          |
| `InverterHeatSinkTemp`      | `Inverter Heat Sink Temperature`        | `value_json.InverterHeatSinkTemp`     |
| `PV1InputCurrent`           | `Inverter PV1 Input Current`            | `value_json.PV1InputCurrent`          |
| `PV1InputVoltage`           | `Inverter PV1 Input Voltage`            | `value_json.PV1InputVoltage`          |
| `BatteryVoltageFromSCC`     | `Inverter Battery Voltage From SCC`     | `value_json.BatteryVoltageFromSCC`    |
| `BatteryDischargeCurrent`   | `Inverter Battery Discharge Current`    | `value_json.BatteryDischargeCurrent`  |
| `DeviceStatus1`             | `Inverter Device Status 1`              | `value_json.DeviceStatus1`            |
| `BatteryVOffsetForFansOn`   | `Inverter Battery V Offset For Fans On` | `value_json.BatteryVOffsetForFansOn`  |
| `EEPROMVersion`             | `Inverter EEPROM Version`               | `value_json.EEPROMVersion`            |
| `PV1ChargingPower`          | `Inverter PV1 Charging Power`           | `value_json.PV1ChargingPower`         |
| `DeviceStatus2`             | `Inverter Device Status 2`              | `value_json.DeviceStatus2`            |

### `mqtt_sensors_qpiri.yaml` (for QPIRI Data)

| QPIRIData Field             | Home Assistant Sensor Name              | `value_template` Mapping              |
| :-------------------------- | :-------------------------------------- | :------------------------------------ |
| `GridRatingVoltage`         | `Inverter Grid Rating Voltage`          | `value_json.GridRatingVoltage`        |
| `GridRatingCurrent`         | `Inverter Grid Rating Current`          | `value_json.GridRatingCurrent`        |
| `ACOutputRatingVoltage`     | `Inverter AC Output Rating Voltage`     | `value_json.ACOutputRatingVoltage`    |
| `ACOutputRatingFrequency`   | `Inverter AC Output Rating Frequency`   | `value_json.ACOutputRatingFrequency`  |
| `ACOutputRatingCurrent`     | `Inverter AC Output Rating Current`     | `value_json.ACOutputRatingCurrent`    |
| `ACOutputRatingApparentPower`| `Inverter AC Output Rating Apparent Power`| `value_json.ACOutputRatingApparentPower`|
| `ACOutputRatingActivePower` | `Inverter AC Output Rating Active Power`| `value_json.ACOutputRatingActivePower`|
| `BatteryRatingVoltage`      | `Inverter Battery Rating Voltage`       | `value_json.BatteryRatingVoltage`     |
| `BatteryRechargeVoltage`    | `Inverter Battery Recharge Voltage`     | `value_json.BatteryRechargeVoltage`   |
| `BatteryUnderVoltage`       | `Inverter Battery Under Voltage`        | `value_json.BatteryUnderVoltage`      |
| `BatteryBulkVoltage`        | `Inverter Battery Bulk Voltage`         | `value_json.BatteryBulkVoltage`       |
| `BatteryFloatVoltage`       | `Inverter Battery Float Voltage`        | `value_json.BatteryFloatVoltage`      |
| `BatteryType`               | `Inverter Battery Type`                 | `value_json.BatteryType`              |
| `MaxACChargingCurrent`      | `Inverter Max AC Charging Current`      | `value_json.MaxACChargingCurrent`     |
| `MaxChargingCurrent`        | `Inverter Max Charging Current`         | `value_json.MaxChargingCurrent`       |
| `InputVoltageRange`         | `Inverter Input Voltage Range`          | `value_json.InputVoltageRange`        |
| `OutputSourcePriority`      | `Inverter Output Source Priority`       | `value_json.OutputSourcePriority`     |
| `ChargerSourcePriority`     | `Inverter Charger Source Priority`      | `value_json.ChargerSourcePriority`    |
| `ParallelMaxNumber`         | `Inverter Parallel Max Number`          | `value_json.ParallelMaxNumber`        |
| `MachineType`               | `Inverter Machine Type`                 | `value_json.MachineType`              |
| `Topology`                  | `Inverter Topology`                     | `value_json.Topology`                 |
| `OutputMode`                | `Inverter Output Mode`                  | `value_json.OutputMode`               |
| `BatteryRedischargeVoltage` | `Inverter Battery Redischarge Voltage`  | `value_json.BatteryRedischargeVoltage`|
| `PVOKConditionForParallel`  | `Inverter PV OK Condition For Parallel` | `value_json.PVOKConditionForParallel` |
| `PVPowerBalance`            | `Inverter PV Power Balance`             | `value_json.PVPowerBalance`           |
| `MaxChargingTimeAtCVStage`  | `Inverter Max Charging Time At CV Stage`| `value_json.MaxChargingTimeAtCVStage` |
| `OperationLogic`            | `Inverter Operation Logic`              | `value_json.OperationLogic`           |
| `MaxDischargingCurrent`     | `Inverter Max Discharging Current`      | `value_json.MaxDischargingCurrent`    |

### `mqtt_sensors_qpigs2.yaml` (for QPIGS2 Data)

| QPIGS2Data Field            | Home Assistant Sensor Name              | `value_template` Mapping              |
| :-------------------------- | :-------------------------------------- | :------------------------------------ |
| `PV2InputCurrent`           | `Inverter PV2 Input Current`            | `value_json.PV2InputCurrent`          |
| `PV2InputVoltage`           | `Inverter PV2 Input Voltage`            | `value_json.PV2InputVoltage`          |
| `PV2ChargingPower`          | `Inverter PV2 Charging Power`           | `value_json.PV2ChargingPower`         |

### `mqtt_sensors_qpiws.yaml` (for QPIWS Data)

| QPIWSData Field             | Home Assistant Sensor Name              | `value_template` Mapping              |
| :-------------------------- | :-------------------------------------- | :------------------------------------ |
| `WarningFlags`              | `Inverter Warning Flags`                | `value_json.WarningFlags`             |

## Analysis of Data Processing Completeness

### QPIGS Command

By comparing the `QPIGSData` struct fields with the `QPIGS` command definition in `axpert_protocol.pdf` (pages 8-10) and `data_format.md`, the following conclusions can be drawn:

1.  **Core Fields (1-21):** All 21 primary fields that are consistently present in the `QPIGS` response from the inverter are accurately captured and represented in the `QPIGSData` struct. The data types (float64, int, string) in the Go struct align with the expected format of the inverter's response.

2.  **Optional Fields (22 onwards):** The protocol documents indicate that fields from 22 onwards (e.g., `Solar feed to grid status`, `Set country regulation`, `Solar feed to grid power`, `Grid input current`) are often not present in the response from this specific inverter model. The `QPIGSData` struct intentionally does not include these fields, as they are not reliably provided by the device.

**Conclusion for QPIGS:** The Go application is processing **100% of the data that is reliably received from the inverter for the QPIGS command.** There are no missing data points that are consistently provided by the inverter and not captured by the `QPIGSData` struct.

### QPIRI Command

By comparing the `QPIRIData` struct fields with the `QPIRI` command definition in `axpert_protocol.pdf` (pages 6-7) and `data_format.md`, it is confirmed that **100% of the data fields for the QPIRI command are being processed** by the Go application and mapped to corresponding Home Assistant MQTT sensors. All 28 fields defined in the protocol for QPIRI are present in the `QPIRIData` struct and have corresponding MQTT sensors.

### QPIGS2 Command

By comparing the `QPIGS2Data` struct fields with the `QPIGS2` command definition in `axpert_protocol.pdf` (page 10) and `data_format.md`, it is confirmed that **100% of the data fields for the QPIGS2 command are being processed** by the Go application and mapped to corresponding Home Assistant MQTT sensors. All 3 fields defined in the protocol for QPIGS2 are present in the `QPIGS2Data` struct and have corresponding MQTT sensors.

### QPIWS Command

By comparing the `QPIWSData` struct fields with the `QPIWS` command definition in `axpert_protocol.pdf` (page 13) and `data_format.md`, it is confirmed that **100% of the data fields for the QPIWS command are being processed** by the Go application and mapped to corresponding Home Assistant MQTT sensors. The single `WarningFlags` field, which represents the 36-bit warning status string, is captured.

## Overall Conclusion on Data Processing Completeness

With the integration of QPIRI, QPIGS2, and QPIWS commands, the Go application is now processing **100% of the data fields defined in the protocol for all implemented inquiry commands (QPIGS, QPIRI, QPIGS2, QPIWS)** that are reliably received from the inverter. All processed fields have corresponding Home Assistant MQTT sensors configured to extract their values.

## Next Steps for Home Assistant Integration

If data is still not appearing or showing an "unknown" state in Home Assistant, the issue is highly likely to reside within the Home Assistant environment itself. The following steps should be taken by the user:

*   **Home Assistant Configuration Loading**: Ensure Home Assistant has been restarted and has successfully loaded all the `mqtt_sensors_*.yaml` files.
*   **MQTT Broker Connectivity (from Home Assistant)**: Verify that Home Assistant's MQTT integration is properly connected to the MQTT broker.
*   **Home Assistant Logs**: Check Home Assistant's logs for any errors or warnings related to MQTT or the specific sensor entities. These logs are crucial for diagnosing issues within Home Assistant's processing of MQTT messages.
*   **Home Assistant Templating Engine**: Manually test the `value_template` expressions in Home Assistant's Developer Tools -> Template editor with a sample JSON payload (like the one shown in the CLI logs) to confirm they extract values as expected.