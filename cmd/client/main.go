package main

import (
	"flag"

	"github.com/yoyotk21/RPCServer/client"
)

func main() {
	c := new(client.Client)
	flag.StringVar(&c.Name, "name", "unknown", "the user's name")
	flag.StringVar(&c.Port, "port", "8000", "the port")
	flag.StringVar(&c.Addr, "addr", "localhost", "the adress")
	flag.Parse()
	
	message := c.GetMessage()
	println(message)
}
