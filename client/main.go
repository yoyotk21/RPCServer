package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	// Connecting to server
	conn, err := rpc.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	g := new(string)
	name := "Joel"

	err = conn.Call("Greeter.SayHi",name, g)
	if err  != nil {
		log.Fatal(err)
	}
	fmt.Println(*g)

}
