package main

import (
	"fmt"
	"strconv"
	"strings"
)

// InverterParser handles parsing raw inverter responses into structured data.
type InverterParser struct {
	// Add any necessary fields here, e.g., for caching or specific parsing rules
}

// NewInverterParser creates a new parser instance.
func NewInverterParser() *InverterParser {
	return &InverterParser{}
}

// QPIGSData holds the parsed data from the QPIGS command.
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
	// Add more fields as needed based on data_format.md
}

// QPIRIData holds the parsed data from the QPIRI command.
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

// ParseQPIGSResponse parses the raw string response from the QPIGS command.
func (ip *InverterParser) ParseQPIGSResponse(rawResponse string) (*QPIGSData, error) {
	// Remove the leading '(' and trailing CRC/CR if present
	cleanedResponse := strings.TrimPrefix(rawResponse, "(")
	// Assuming the response ends with CRC and CR, we need to remove them.
	// For now, let's just trim the last char if it's CR, and assume CRC is handled elsewhere or not present in this string.
	cleanedResponse = strings.TrimSuffix(cleanedResponse, "\r")

	// The protocol document shows a space-separated list of values.
	parts := strings.Fields(cleanedResponse)

	// Based on data_format.md, QPIGS has 21 fields before N/A ones.
	// We need to be careful with the exact number of fields and their types.
	// This is a simplified parsing and needs to be robustified.
	// Example raw data: (229.8 49.8 229.8 49.8 0781 0583 009 396 00.00 000 000 0034 00.0 000.0 00.00 00000 00010000 00 00 00000 010

	if len(parts) < 21 { // Minimum expected fields
		return nil, fmt.Errorf("QPIGS response has too few fields: %d", len(parts))
	}

	data := &QPIGSData{}
	var err error

	// Parse each field according to data_format.md
	// This is a direct mapping and needs careful indexing.
	// Field 1: Grid Voltage (BBB.B)
	data.GridVoltage, err = strconv.ParseFloat(parts[0], 64)
	if err != nil { return nil, fmt.Errorf("error parsing GridVoltage: %w", err) }

	// Field 2: Grid Frequency (CC.C)
	data.GridFrequency, err = strconv.ParseFloat(parts[1], 64)
	if err != nil { return nil, fmt.Errorf("error parsing GridFrequency: %w", err) }

	// Field 3: AC Output Voltage (DDD.D)
	data.ACOutputVoltage, err = strconv.ParseFloat(parts[2], 64)
	if err != nil { return nil, fmt.Errorf("error parsing ACOutputVoltage: %w", err) }

	// Field 4: AC Output Frequency (EE.E)
	data.ACOutputFrequency, err = strconv.ParseFloat(parts[3], 64)
	if err != nil { return nil, fmt.Errorf("error parsing ACOutputFrequency: %w", err) }

	// Field 5: AC Output Apparent Power (FFFF)
	data.ACOutputApparentPower, err = strconv.Atoi(parts[4])
	if err != nil { return nil, fmt.Errorf("error parsing ACOutputApparentPower: %w", err) }

	// Field 6: AC Output Active Power (GGGG)
	data.ACOutputActivePower, err = strconv.Atoi(parts[5])
	if err != nil { return nil, fmt.Errorf("error parsing ACOutputActivePower: %w", err) }

	// Field 7: Output Load Percent (HHH)
	data.OutputLoadPercent, err = strconv.Atoi(parts[6])
	if err != nil { return nil, fmt.Errorf("error parsing OutputLoadPercent: %w", err) }

	// Field 8: BUS Voltage (III)
	data.BUSVoltage, err = strconv.Atoi(parts[7])
	if err != nil { return nil, fmt.Errorf("error parsing BUSVoltage: %w", err) }

	// Field 9: Battery Voltage (JJ.JJ)
	data.BatteryVoltage, err = strconv.ParseFloat(parts[8], 64)
	if err != nil { return nil, fmt.Errorf("error parsing BatteryVoltage: %w", err) }

	// Field 10: Battery Charging Current (KKK)
	data.BatteryChargingCurrent, err = strconv.Atoi(parts[9])
	if err != nil { return nil, fmt.Errorf("error parsing BatteryChargingCurrent: %w", err) }

	// Field 11: Battery Capacity (OOO)
	data.BatteryCapacity, err = strconv.Atoi(parts[10])
	if err != nil { return nil, fmt.Errorf("error parsing BatteryCapacity: %w", err) }

	// Field 12: Inverter Heat Sink Temp. (TTTT)
	data.InverterHeatSinkTemp, err = strconv.Atoi(parts[11])
	if err != nil { return nil, fmt.Errorf("error parsing InverterHeatSinkTemp: %w", err) }

	// Field 13: PV1 Input Current (EE.E)
	data.PV1InputCurrent, err = strconv.ParseFloat(parts[12], 64)
	if err != nil { return nil, fmt.Errorf("error parsing PV1InputCurrent: %w", err) }

	// Field 14: PV1 Input Voltage (UUU.U)
	data.PV1InputVoltage, err = strconv.ParseFloat(parts[13], 64)
	if err != nil { return nil, fmt.Errorf("error parsing PV1InputVoltage: %w", err) }

	// Field 15: Battery Voltage from SCC (WW.WW)
	data.BatteryVoltageFromSCC, err = strconv.ParseFloat(parts[14], 64)
	if err != nil { return nil, fmt.Errorf("error parsing BatteryVoltageFromSCC: %w", err) }

	// Field 16: Battery Discharge Current (PPPPP)
	data.BatteryDischargeCurrent, err = strconv.Atoi(parts[15])
	if err != nil { return nil, fmt.Errorf("error parsing BatteryDischargeCurrent: %w", err) }

	// Field 17: Device Status 1 (b7..b0)
	data.DeviceStatus1 = parts[16]

	// Field 18: Battery V offset for fans on (QQ)
	data.BatteryVOffsetForFansOn, err = strconv.Atoi(parts[17])
	if err != nil { return nil, fmt.Errorf("error parsing BatteryVOffsetForFansOn: %w", err) }

	// Field 19: EEPROM Version (VV)
	data.EEPROMVersion, err = strconv.Atoi(parts[18])
	if err != nil { return nil, fmt.Errorf("error parsing EEPROMVersion: %w", err) }

	// Field 20: PV1 Charging Power (MMMMM)
	data.PV1ChargingPower, err = strconv.Atoi(parts[19])
	if err != nil { return nil, fmt.Errorf("error parsing PV1ChargingPower: %w", err) }

	// Field 21: Device Status 2 (b10b9b8)
	data.DeviceStatus2 = parts[20]

	return data, nil
}

// ParseQPIWSResponse parses the raw string response from the QPIWS command.
func (ip *InverterParser) ParseQPIWSResponse(rawResponse string) (*QPIWSData, error) {
	cleanedResponse := strings.TrimPrefix(rawResponse, "(")
	cleanedResponse = strings.TrimSuffix(cleanedResponse, "\r")

	// QPIWS response is a single bit string
	if len(cleanedResponse) == 0 {
		return nil, fmt.Errorf("QPIWS response is empty")
	}

	data := &QPIWSData{
		WarningFlags: cleanedResponse,
	}

	return data, nil
}
type QPIWSData struct {
	WarningFlags string // Raw bit string for now
	// Individual flags can be parsed as needed
}

// ParseQPIGS2Response parses the raw string response from the QPIGS2 command.
func (ip *InverterParser) ParseQPIGS2Response(rawResponse string) (*QPIGS2Data, error) {
	cleanedResponse := strings.TrimPrefix(rawResponse, "(")
	cleanedResponse = strings.TrimSuffix(cleanedResponse, "\r")
	parts := strings.Fields(cleanedResponse)

	if len(parts) < 3 { // Based on data_format.md
		return nil, fmt.Errorf("QPIGS2 response has too few fields: %d", len(parts))
	}

	data := &QPIGS2Data{}
	var err error

	data.PV2InputCurrent, err = strconv.ParseFloat(parts[0], 64)
	if err != nil { return nil, fmt.Errorf("error parsing PV2InputCurrent: %w", err) }
	data.PV2InputVoltage, err = strconv.ParseFloat(parts[1], 64)
	if err != nil { return nil, fmt.Errorf("error parsing PV2InputVoltage: %w", err) }
	data.PV2ChargingPower, err = strconv.Atoi(parts[2])
	if err != nil { return nil, fmt.Errorf("error parsing PV2ChargingPower: %w", err) }

	return data, nil
}

// QPIGS2Data holds the parsed data from the QPIGS2 command.
type QPIGS2Data struct {
	PV2InputCurrent  float64
	PV2InputVoltage  float64
	PV2ChargingPower int
}

// ParseQPIRIResponse parses the raw string response from the QPIRI command.
func (ip *InverterParser) ParseQPIRIResponse(rawResponse string) (*QPIRIData, error) {
	cleanedResponse := strings.TrimPrefix(rawResponse, "(")
	cleanedResponse = strings.TrimSuffix(cleanedResponse, "\r")
	parts := strings.Fields(cleanedResponse)

	if len(parts) < 28 { // Based on data_format.md
		return nil, fmt.Errorf("QPIRI response has too few fields: %d", len(parts))
	}

	data := &QPIRIData{}
	var err error

	// Parse each field according to data_format.md
	data.GridRatingVoltage, err = strconv.ParseFloat(parts[0], 64)
	if err != nil { return nil, fmt.Errorf("error parsing GridRatingVoltage: %w", err) }
	data.GridRatingCurrent, err = strconv.ParseFloat(parts[1], 64)
	if err != nil { return nil, fmt.Errorf("error parsing GridRatingCurrent: %w", err) }
	data.ACOutputRatingVoltage, err = strconv.ParseFloat(parts[2], 64)
	if err != nil { return nil, fmt.Errorf("error parsing ACOutputRatingVoltage: %w", err) }
	data.ACOutputRatingFrequency, err = strconv.ParseFloat(parts[3], 64)
	if err != nil { return nil, fmt.Errorf("error parsing ACOutputRatingFrequency: %w", err) }
	data.ACOutputRatingCurrent, err = strconv.ParseFloat(parts[4], 64)
	if err != nil { return nil, fmt.Errorf("error parsing ACOutputRatingCurrent: %w", err) }
	data.ACOutputRatingApparentPower, err = strconv.Atoi(parts[5])
	if err != nil { return nil, fmt.Errorf("error parsing ACOutputRatingApparentPower: %w", err) }
	data.ACOutputRatingActivePower, err = strconv.Atoi(parts[6])
	if err != nil { return nil, fmt.Errorf("error parsing ACOutputRatingActivePower: %w", err) }
	data.BatteryRatingVoltage, err = strconv.ParseFloat(parts[7], 64)
	if err != nil { return nil, fmt.Errorf("error parsing BatteryRatingVoltage: %w", err) }
	data.BatteryRechargeVoltage, err = strconv.ParseFloat(parts[8], 64)
	if err != nil { return nil, fmt.Errorf("error parsing BatteryRechargeVoltage: %w", err) }
	data.BatteryUnderVoltage, err = strconv.ParseFloat(parts[9], 64)
	if err != nil { return nil, fmt.Errorf("error parsing BatteryUnderVoltage: %w", err) }
	data.BatteryBulkVoltage, err = strconv.ParseFloat(parts[10], 64)
	if err != nil { return nil, fmt.Errorf("error parsing BatteryBulkVoltage: %w", err) }
	data.BatteryFloatVoltage, err = strconv.ParseFloat(parts[11], 64)
	if err != nil { return nil, fmt.Errorf("error parsing BatteryFloatVoltage: %w", err) }
	data.BatteryType, err = strconv.Atoi(parts[12])
	if err != nil { return nil, fmt.Errorf("error parsing BatteryType: %w", err) }
	data.MaxACChargingCurrent, err = strconv.Atoi(parts[13])
	if err != nil { return nil, fmt.Errorf("error parsing MaxACChargingCurrent: %w", err) }
	data.MaxChargingCurrent, err = strconv.Atoi(parts[14])
	if err != nil { return nil, fmt.Errorf("error parsing MaxChargingCurrent: %w", err) }
	data.InputVoltageRange, err = strconv.Atoi(parts[15])
	if err != nil { return nil, fmt.Errorf("error parsing InputVoltageRange: %w", err) }
	data.OutputSourcePriority, err = strconv.Atoi(parts[16])
	if err != nil { return nil, fmt.Errorf("error parsing OutputSourcePriority: %w", err) }
	data.ChargerSourcePriority, err = strconv.Atoi(parts[17])
	if err != nil { return nil, fmt.Errorf("error parsing ChargerSourcePriority: %w", err) }
	data.ParallelMaxNumber, err = strconv.Atoi(parts[18])
	if err != nil { return nil, fmt.Errorf("error parsing ParallelMaxNumber: %w", err) }
	data.MachineType, err = strconv.Atoi(parts[19])
	if err != nil { return nil, fmt.Errorf("error parsing MachineType: %w", err) }
	data.Topology, err = strconv.Atoi(parts[20])
	if err != nil { return nil, fmt.Errorf("error parsing Topology: %w", err) }
	data.OutputMode, err = strconv.Atoi(parts[21])
	if err != nil { return nil, fmt.Errorf("error parsing OutputMode: %w", err) }
	data.BatteryRedischargeVoltage, err = strconv.ParseFloat(parts[22], 64)
	if err != nil { return nil, fmt.Errorf("error parsing BatteryRedischargeVoltage: %w", err) }
	data.PVOKConditionForParallel, err = strconv.Atoi(parts[23])
	if err != nil { return nil, fmt.Errorf("error parsing PVOKConditionForParallel: %w", err) }
	data.PVPowerBalance, err = strconv.Atoi(parts[24])
	if err != nil { return nil, fmt.Errorf("error parsing PVPowerBalance: %w", err) }
	data.MaxChargingTimeAtCVStage, err = strconv.Atoi(parts[25])
	if err != nil { return nil, fmt.Errorf("error parsing MaxChargingTimeAtCVStage: %w", err) }
	data.OperationLogic, err = strconv.Atoi(parts[26])
	if err != nil { return nil, fmt.Errorf("error parsing OperationLogic: %w", err) }
	data.MaxDischargingCurrent, err = strconv.Atoi(parts[27])
	if err != nil { return nil, fmt.Errorf("error parsing MaxDischargingCurrent: %w", err) }

	return data, nil
}
