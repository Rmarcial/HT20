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
	ip1      = "192.168.1.19:502"
	ip2      = "192.168.1.20:502"
	modbusId = 1
	timeout  = 10 * time.Second
)

// HT20 32-bit register addresses (contain floats in IEEE-754 standard. Order of bytes: B3 B2 B1 B0)
const (
	TemperatureAddr = 7500
	HumidityAddr    = 7501
	DewpointAddr    = 7502
)

func main() {

	// ModbusTCP Handler and Client
	handler := modbus.NewTCPClientHandler(ip2)
	handler.SlaveId = modbusId
	handler.Timeout = timeout

	err := handler.Connect()
	defer handler.Close()
	if err != nil {
		println("Error connecting handler:", err.Error())
	}

	client := modbus.NewClient(handler)

	// Read Measurements
	bsTemp, err := client.ReadInputRegisters(TemperatureAddr, 1)
	if err != nil {
		fmt.Println("Error reading temperature", err.Error())
	}
	temperature := castValue(bsTemp)

	bsHum, err := client.ReadInputRegisters(HumidityAddr, 1)
	if err != nil {
		fmt.Println("Error reading humidity", err.Error())
	}
	humidity := castValue(bsHum)

	bsDew, err := client.ReadInputRegisters(DewpointAddr, 1)
	if err != nil {
		fmt.Println("Error reading dew point", err.Error())
	}
	dewpoint := castValue(bsDew)

	// Print
	fmt.Println("HT20 readings:")
	fmt.Println()
	fmt.Printf("Temperature (ºC): \t %v \n", temperature)
	fmt.Printf("Relative Humidity (%%): \t %v \n", humidity)
	fmt.Printf("Dew Point (ºC): \t %v \n", dewpoint)

}

func castValue(bs []byte) float32 {
	var value float32
	var reader = bytes.NewReader(bs)

	if readError := binary.Read(reader, binary.BigEndian, &value); readError != nil {
		fmt.Println("Error reading binary data")
	}
	return value
}
