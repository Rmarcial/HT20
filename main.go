package main

import (
	"fmt"
	"time"

	"github.com/goburrow/modbus"
)

// Modbus TCP constants
const (
	Ip = "192.168.1.19"
	Id = 1
)

// HT20 16-bit register addresses
const (
	TemperatureAddr = 6000
	HumidityAddr    = 6002
	DewPointAddr    = 6004
)

func main() {

	handler := modbus.NewTCPClientHandler("localhost:502")
	handler.Timeout = 10 * time.Second
	handler.SlaveId = byte(Id)

	err := handler.Connect()
	defer handler.Close()

	client := modbus.NewClient(handler)

	fmt.Println(err)
	fmt.Println(client)

	fmt.Println("HT20 readings:")

}
