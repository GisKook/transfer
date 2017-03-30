package event_handler

import (
	"github.com/giskook/gotcp"
	"github.com/giskook/transfer/conf"
	"github.com/giskook/transfer/conn"
	"github.com/giskook/transfer/pkg"
	"github.com/giskook/transfer/protocol"
	"log"
)

type UpstreamCallback struct{}

func (this *UpstreamCallback) OnConnect(c *gotcp.Conn) bool {
	checkinterval := conf.GetConfiguration().ConnCheckInterval
	readlimit := conf.GetConfiguration().ReadLimit
	writelimit := conf.GetConfiguration().WriteLimit
	config := &conn.ConnConfig{
		ConnCheckInterval: uint16(checkinterval),
		ReadLimit:         uint16(readlimit),
		WriteLimit:        uint16(writelimit),
	}
	_conn := conn.NewConn(c, config)

	c.PutExtraData(_conn)

	_conn.Do()
	conn.NewConnsUpstream().Add(_conn)

	return true
}

func (this *UpstreamCallback) OnClose(c *gotcp.Conn) {
	_conn := c.GetExtraData().(*conn.Conn)
	_conn.Close()
	conn.NewConnsUpstream().Remove(_conn)
}

func (this *UpstreamCallback) OnMessage(c *gotcp.Conn, p gotcp.Packet) bool {
	tt_pkg := p.(*pkg.TransparentTransmissionPacket)
	switch tt_pkg.Type {
	case protocol.PROTOCOL_UP_REQ_REGISTER:
		log.Println("PROTOCOL_UP_REQ_REGISTER")
		event_handler_up_req_register(c, tt_pkg)
	case protocol.PROTOCOL_UP_REQ_CANCEL:
		log.Println("PROTOCOL_UP_REQ_CANCEL")
		event_handler_up_req_cancel(c, tt_pkg)
	case protocol.PROTOCOL_UP_TRANSFER:
		log.Println("PROTOCOL_UP_TRANSFER")
		event_handler_up_transfer(c, tt_pkg)
	}
	return true
}
