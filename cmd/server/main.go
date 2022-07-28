package main

import (
	"flag"

	"github.com/yoyotk21/RPCServer/server"
)

func main() {
	s := new(server.Server)

	flag.StringVar(&s.Port, "p", "8000", "the port")
	flag.Parse()

	s.Data = []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10};

	s.RunServer()
}
