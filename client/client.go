package client

import (
	"log"
	"net/rpc"

	"github.com/yoyotk21/RPCServer/dpf"
)

type Client struct {
	Port1 string
	Port2 string
	Addr  string
}

func (client *Client) GetIndex(index uint64) uint64 {
	// Connecting to first server
	conn1, err := rpc.Dial("tcp", client.Addr+":"+client.Port1)
	if err != nil {
		log.Fatal(err)
	}
	defer conn1.Close()

	// Connecting to the second server
	conn2, err := rpc.Dial("tcp", client.Addr+":"+client.Port2)
	if err != nil {
		log.Fatal(err)
	}
	defer conn2.Close()

	dpfClient := dpf.ClientDPFInitialize()
	key1, key2 := dpfClient.GenDPFKeys(index, 64)

	var ans1, ans2 *uint64

	var tmp int
	conn1.Call("Database.Init", dpfClient.PrfKey, &tmp)
	conn2.Call("Database.Init", dpfClient.PrfKey, &tmp)
	conn1.Call("Database.BatchEval", key1, &ans1)
	conn2.Call("Database.BatchEval", key2, &ans2)

	if *ans1 > *ans2 {
		return *ans1 - *ans2
	} else {
		return *ans2 - *ans1
	}
}
