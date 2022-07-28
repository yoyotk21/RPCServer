package client

import (
	"log"
	"net/rpc"
)

type Client struct {
	Name string
	Port string
	Addr string
}

func (client *Client) GetMessage() string {
	// Connecting to server
	conn, err := rpc.Dial("tcp", client.Addr+":"+client.Port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	g := new(string)

	err = conn.Call("Greeter.SayHi", client.Name, g)
	if err != nil {
		log.Fatal(err)
	}

	return *g
}
