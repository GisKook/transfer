package event_handler

import (
	"github.com/giskook/gotcp"
	"github.com/giskook/transfer/conn"
	"github.com/giskook/transfer/pkg"
	//"github.com/giskook/transfer/protocol"
)

func event_handler_up_req_cancel(c *gotcp.Conn, p *pkg.TransparentTransmissionPacket) {
	connection := c.GetExtraData().(*conn.Conn)
	if connection != nil {
		//cancel_pkg := p.Packet.(*protocol.UpCancelPacket)
		connection.SendToTerm(p)
		c.Close()
	}
}
