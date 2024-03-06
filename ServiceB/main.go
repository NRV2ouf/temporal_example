package main

import (
	"log"

	"go.temporal.io/sdk/client"
)

/*
	TO REMOVE ?
*/

func main() {
	// Create a client
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()
}
