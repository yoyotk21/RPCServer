package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Server struct {
	Port string
	Data []uint64
}

func (server *Server) RunServer() {
	listener, err := net.Listen("tcp", ":"+server.Port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port", server.Port)

	// Closing the connection once the program is finished
	defer listener.Close()


	database := newDatabase(server.Data)

	rpcServer := rpc.NewServer()

	rpcServer.Register(database)

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
