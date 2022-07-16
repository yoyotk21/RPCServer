package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"server/common"
)

func main() {
	// Tcp connection on port :8000
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening...")

	// Closing the connection once the program is finished
	defer listener.Close()

	g := common.NewGreeter("Welcome")

	server := rpc.NewServer()

	server.Register(g)

	// Infinite loop to run the server
	for {
		// Waits for connection
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go server.ServeConn(conn)
	}

}
