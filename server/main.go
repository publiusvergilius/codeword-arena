package main

import (
	"bufio"
	"fmt"
	"net"
)

func main () {

	fmt.Println("Starting Codeword Arena TCP server")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	defer listener.Close()
	fmt.Println("TCP server listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		
		go HandleConnection(conn)
	}

}

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed")
		}

		fmt.Printf("Received: %s", message)

		conn.Write([]byte("Echo: " + message))
	}

}



