package server

import (
	dpf "github.com/sachaservan/vdpf"
)

type Database struct {
	Data      []uint64
	DpfServer *dpf.Dpf
}

func newDatabase(data []uint64) *Database {
	newDatabase := new(Database)
	newDatabase.Data = data

	return newDatabase
}

func (d *Database) Init(key dpf.PrfKey, reply *int) error {
	d.DpfServer = dpf.ServerDPFInitialize(key)
	return nil
}

func (d *Database) BatchEval(key *dpf.DPFKey, reply *uint64) error {
	list := d.DpfServer.BatchEval(key, d.Data)
	var sum uint64 = 0
	for i := range d.Data {
		sum += d.Data[i] * uint64(list[i])
	}
	*reply = sum
	return nil
}
