package main

import (
	"flag"

	"github.com/yoyotk21/RPCServer/server"
)

func main() {
	s := new(server.Server)

	flag.StringVar(&s.Port, "port", "8000", "the port")
	flag.StringVar(&s.WelcomeMessage, "m", "Welcome,", "the welcome message")
	flag.Parse()

	s.RunServer()
}
