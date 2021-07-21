package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/goburrow/modbus"
)

// Modbus TCP constants
const (
	ipAddr   = "192.168.1.20:502"
	modbusId = 1
)

// HT20 32-bit register addresses (contain floats in IEEE-754 standard. Order of bytes: B3 B2 B1 B0)
const (
	TemperatureAddr = 7500
	HumidityAddr    = 7501
	DewPointAddr    = 7502
)

func main() {

	// Handler
	handler := modbus.NewTCPClientHandler(ipAddr)
	handler.Timeout = 10 * time.Second
	handler.SlaveId = modbusId

	err := handler.Connect()
	defer handler.Close()
	if err != nil {
		println("Error connecting handler:", err.Error())
	}

	// Client
	client := modbus.NewClient(handler)
	fmt.Println(client)

	// Read Measurements
	bsTemp, err := client.ReadInputRegisters(TemperatureAddr, 2)
	if err != nil {
		fmt.Println("Error reading temperature", err.Error())
	}

	var temperature float32
	var reader = bytes.NewReader(bsTemp)

	if readErr := binary.Read(reader, binary.BigEndian, &temperature); readErr != nil {
		fmt.Println("Error reading binary data")
	}

	temperature = float32(temperature)

	// Print
	fmt.Println("HT20 readings:")
	fmt.Println()
	fmt.Printf("Temperature (ºC): \t %v \n", temperature)
	fmt.Printf("Relative Humidity (%%): \t ????\n")
	fmt.Printf("Dew Point (ºC): \t ????\n")

}
