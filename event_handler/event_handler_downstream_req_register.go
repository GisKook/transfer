package event_handler

import (
	"github.com/giskook/gotcp"
	"github.com/giskook/transfer/conn"
	"github.com/giskook/transfer/pkg"
	"github.com/giskook/transfer/protocol"
	"log"
)

func event_handler_down_req_register(c *gotcp.Conn, p *pkg.TransparentTransmissionPacket) {
	log.Println("event_handler_down_req_register")

	connection := c.GetExtraData().(*conn.Conn)
	if connection != nil {
		register_pkg := p.Packet.(*protocol.DownRegisterPacket)
		connection.ID = register_pkg.RegisterID
		connection.PeerID = register_pkg.PeerClientID
		connection.TransparentTransmissionKey = register_pkg.TransparentTransmissionKey
		connection.Mode = 1
		conn.NewConnsDownstream().SetID(register_pkg.RegisterID, connection)
		register_pkg.Status = 0
		connection.SendToTerm(p)
	}
}
