package main

import (
	"log"
	"net"
	"os"
)
func main () {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println("", err)
		os.Exit(1)
	}
	defer conn.Close()

	message := "Ola, servidor!\n"

	_, err = conn.Write([]byte(message))

	if err != nil {
		log.Println("Error on sending: ", err)
		return
	}

	log.Println("Message sent")
}
