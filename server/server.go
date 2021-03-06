package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Server struct {
	Port           string
	WelcomeMessage string
}

func (server *Server) RunServer() {
	listener, err := net.Listen("tcp", ":"+server.Port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port", server.Port)

	// Closing the connection once the program is finished
	defer listener.Close()

	g := NewGreeter(server.WelcomeMessage)

	rpcServer := rpc.NewServer()

	rpcServer.Register(g)

	// Infinite loop to run the server
	for {
		// Waits for connection
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go rpcServer.ServeConn(conn)
	}
}
