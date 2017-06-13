package conn

import (
	"log"
	"sync/atomic"
)

type Conns struct {
	Connsindex map[uint32]*Conn
	connsuid   map[uint64]*Conn
	index      uint32
}

var conns_upstream_instance *Conns
var conns_downstream_instance *Conns

func NewConnsUpstream() *Conns {
	if conns_upstream_instance == nil {
		conns_upstream_instance = &Conns{
			Connsindex: make(map[uint32]*Conn),
			connsuid:   make(map[uint64]*Conn),
			index:      0,
		}
	}

	return conns_upstream_instance
}

func NewConnsDownstream() *Conns {
	if conns_downstream_instance == nil {
		conns_downstream_instance = &Conns{
			Connsindex: make(map[uint32]*Conn),
			connsuid:   make(map[uint64]*Conn),
			index:      0,
		}
	}

	return conns_downstream_instance
}

func (cs *Conns) Add(conn *Conn) {
	conn.index = atomic.AddUint32(&cs.index, 1)
	cs.Connsindex[conn.index] = conn
}

func (cs *Conns) SetID(gatewayid uint64, conn *Conn) {
	cs.connsuid[gatewayid] = conn
}

func (cs *Conns) GetConn(uid uint64) *Conn {
	return cs.connsuid[uid]
}

func (cs *Conns) Remove(c *Conn) {
	delete(cs.Connsindex, c.index)

	connuid, ok := cs.connsuid[c.ID]
	if ok && c.index == connuid.index {
		delete(cs.connsuid, c.ID)
	}
}

func (cs *Conns) Check(uid uint64) bool {
	conn, ok := cs.connsuid[uid]
	if ok {
		_, realok := cs.Connsindex[conn.index]

		return realok
	}
	return ok
}

func (cs *Conns) CheckKey(key uint32) (bool, uint64) {
	log.Println("---")
	log.Println(cs.connsuid)
	log.Println(cs.Connsindex)
	log.Println("---")
	for _, conn := range cs.connsuid {
		if conn.TransparentTransmissionKey == key {
			return true, conn.ID
		}
	}

	return false, 0
}

func (cs *Conns) GetCount() int {
	return len(cs.Connsindex)
}
