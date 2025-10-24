package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// MQTTPublisher handles connecting to an MQTT broker and publishing messages.
type MQTTPublisher struct {
	client mqtt.Client
	config MQTTConfig
}

// MQTTConfig holds the configuration for the MQTT connection.
type MQTTConfig struct {
	Server     string `json:"server"`
	Port       string `json:"port"`
	Topic      string `json:"topic"`
	DeviceName string `json:"devicename"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	ClientID   string `json:"clientid"`	
}

// NewMQTTPublisher creates a new MQTT publisher instance.
func NewMQTTPublisher(config MQTTConfig) *MQTTPublisher {
	return &MQTTPublisher{
		config: config,
	}
}

// Connect establishes a connection to the MQTT broker.
func (mp *MQTTPublisher) Connect() error {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", mp.config.Server, mp.config.Port))

	rand.Seed(time.Now().UnixNano())
	clientID := mp.config.ClientID + "_" + strconv.Itoa(rand.Intn(100000))
	opts.SetClientID(clientID)
	opts.SetUsername(mp.config.Username)
	opts.SetPassword(mp.config.Password)
	opts.SetKeepAlive(60 * time.Second)
	opts.SetPingTimeout(5 * time.Second)
	opts.SetAutoReconnect(true)

	// Set up handlers for connection events
	opts.SetOnConnectHandler(func(client mqtt.Client) {
		fmt.Println("Connected to MQTT broker!")
	})
	opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		fmt.Printf("MQTT connection lost: %v\n", err)
	})

	mp.client = mqtt.NewClient(opts)
	if token := mp.client.Connect(); token.Wait() && token.Error() != nil {
		return fmt.Errorf("failed to connect to MQTT broker: %w", token.Error())
	}
	return nil
}

// Disconnect closes the connection to the MQTT broker.
func (mp *MQTTPublisher) Disconnect() {
	if mp.client != nil && mp.client.IsConnected() {
		mp.client.Disconnect(250) // Disconnect with 250ms grace period
		fmt.Println("Disconnected from MQTT broker.")
	}
}

// PublishData publishes structured data to a specific sub-topic.
func (mp *MQTTPublisher) PublishData(data interface{}, subTopic string) error {
	if mp.client == nil || !mp.client.IsConnected() {
		return fmt.Errorf("not connected to MQTT broker")
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data to JSON: %w", err)
	}

	topic := fmt.Sprintf("%s/%s/%s", mp.config.Topic, mp.config.DeviceName, subTopic)
	token := mp.client.Publish(topic, 1, false, payload)
	token.Wait()
	if token.Error() != nil {
		return fmt.Errorf("failed to publish message: %w", token.Error())
	}

	// fmt.Printf("Published to topic %s: %s", topic, payload)
	fmt.Printf("Published to topic %s -> ", topic)
	return nil
}
