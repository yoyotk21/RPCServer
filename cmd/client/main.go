package main

import (
	"flag"
	"fmt"

	"github.com/yoyotk21/RPCServer/client"
)

func main() {
	c := new(client.Client)
	flag.StringVar(&c.Addr, "a", "localhost", "the address")
	flag.StringVar(&c.Port1, "p1", "8000", "the port of the first server")
	flag.StringVar(&c.Port2, "p2", "8080", "the port of the second server")
	var index uint64
	flag.Uint64Var(&index, "i", 0, "the index to retrieve from the database")
	flag.Parse()
	
	res := c.GetIndex(index)
	fmt.Println(res)
}
