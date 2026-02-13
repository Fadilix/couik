package main

import (
	"log"

	"github.com/fadilix/couik/pkg/network"
)

func main() {
	server := network.NewServer()
	log.Println("Starting server on port 4217...")
	server.Start()
}
