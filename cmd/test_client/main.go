package main

import (
	"fmt"
	"log"

	"github.com/fadilix/couik/pkg/network"
)

func main() {
	client, err := network.NewClient("localhost:4217")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer client.Close()

	fmt.Println("Connected to server successfully!")

	fmt.Println("Sending Join message...")
	if err := client.SendJoin("TestUser"); err != nil {
		log.Printf("Failed to join: %v", err)
	}

	fmt.Println("Waiting for messages from server...")

	for range 3 {
		select {
		case msg := <-client.NextMessage():
			fmt.Printf("Received: Type=%s Payload=%s\n", msg.Type, string(msg.Payload))
		case err := <-client.Errors():
			log.Fatalf("Network error: %v", err)
		}
	}
}
