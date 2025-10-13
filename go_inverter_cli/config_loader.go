package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// LoadMQTTConfig reads the MQTT configuration from the specified file path.
func LoadMQTTConfig(filePath string) (MQTTConfig, error) {
	var config MQTTConfig

	file, err := os.ReadFile(filePath)
	if err != nil {
		return config, fmt.Errorf("failed to read MQTT config file %s: %w", filePath, err)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return config, fmt.Errorf("failed to unmarshal MQTT config from %s: %w", filePath, err)
	}

	return config, nil
}
