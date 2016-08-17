package transfer

import (
	"github.com/giskook/gotcp"
	"time"
)

type UpstreamCallback struct{}

func (this *UpstreamCallback) OnConnect(c *gotcp.Conn) bool {
	checkinterval := GetConfiguration().ConnCheckInterval
	readlimit := GetConfiguration().ReadLimit
	writelimit := GetConfiguration().WriteLimit
	config := &ConnConfig{
		ConnCheckInterval: uint16(checkinterval),
		ReadLimit:         uint16(readlimit),
		WriteLimit:        uint16(writelimit),
	}
	conn := NewConn(c, config)

	c.PutExtraData(conn)

	conn.Do()
	NewConnsUpstream().Add(conn)

	return true
}

func (this *UpstreamCallback) OnClose(c *gotcp.Conn) {
	conn := c.GetExtraData().(*Conn)
	conn.Close()
	NewConnsUpstream().Remove(conn)
}

func (this *UpstreamCallback) OnMessage(c *gotcp.Conn, p gotcp.Packet) bool {
	//c.AsyncWritePacket(p, time.Second)
	var conn *Conn
	for i := 0; i <= int(NewConnsDownstream().index); i++ {
		conn = nil
		conn = NewConnsDownstream().connsindex[uint32(i)]
		if conn != nil {
			conn.conn.AsyncWritePacket(p, time.Second)
		}
	}

	return true
}
