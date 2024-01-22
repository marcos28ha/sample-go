package main

import (
	"fmt";
	"github.com/nats-io/nats.go"
)


func main() {
    fmt.Println("Miguelito Maravillas")
	nc, err := nats.Connect("nats://dialasset-nats-server-1:4222")
}