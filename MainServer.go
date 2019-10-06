package main

import (
	"HotShower/HandleClient"
	"fmt"
	"net"
	"os"
)

const (
	HOST          = "127.0.0.1"
	PORT          = "8080"
	PROTOCOL_TYPE = "tcp"
)

func main() {
	fmt.Println("HandleClient Starts")
	// listen on all interfaces
	ln, err := net.Listen(PROTOCOL_TYPE, HOST + ":" + PORT)
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer ln.Close()

	fmt.Println("Listening on " + HOST + ":" + PORT)
	// accept connection on port
	for {
		fmt.Println("Waiting for new connection...")
		// accept a connection
		conn, err := ln.Accept()
		HandleClient.ReceiveRequest(conn)
		if err != nil {
			fmt.Println("Error connecting to user:", err.Error())
			os.Exit(1)
		}
	}
}
