package transfer

import (
	"github.com/giskook/gotcp"
)

type DownstreamCallback struct{}

func (this *DownstreamCallback) OnConnect(c *gotcp.Conn) bool {
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
	NewConnsDownstream().Add(conn)

	return true
}

func (this *DownstreamCallback) OnClose(c *gotcp.Conn) {
	conn := c.GetExtraData().(*Conn)
	conn.Close()
	NewConnsDownstream().Remove(conn)
}

func (this *DownstreamCallback) OnMessage(c *gotcp.Conn, p gotcp.Packet) bool {
	return true
}
