package event_handler

import (
	"github.com/giskook/gotcp"
	"github.com/giskook/transfer/conn"
	"github.com/giskook/transfer/pkg"
	"github.com/giskook/transfer/protocol"
	"log"
)

func event_handler_up_req_register(c *gotcp.Conn, p *pkg.TransparentTransmissionPacket) {
	connection := c.GetExtraData().(*conn.Conn)
	if connection != nil {
		register_pkg := p.Packet.(*protocol.UpRegisterPacket)
		connection.ID = register_pkg.ID
		//connection.PeerID = register_pkg.PeerRouterRegisterID
		connection.TransparentTransmissionKey = register_pkg.TransparentTransmissionKey
		ok, router_id := conn.NewConnsDownstream().CheckKey(register_pkg.TransparentTransmissionKey)
		log.Println(ok)
		log.Println(router_id)
		conn.NewConnsUpstream().SetID(register_pkg.ID, connection)

		connection.PeerID = router_id
		if ok {
			register_pkg.Status = 0
		} else {
			register_pkg.Status = 1
		}
		connection.SendToTerm(register_pkg)
	}
}
