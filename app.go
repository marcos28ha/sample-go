package main

import (
	"fmt";
	"github.com/nats-io/nats.go"
)


func main() {
    fmt.Println("Miguelito Maravillas")
	nc, err := nats.Connect("nats://nats-1:4222")
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer nc.Close()
}
