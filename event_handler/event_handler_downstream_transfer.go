package event_handler

import (
	"github.com/giskook/gotcp"
	"github.com/giskook/transfer/conn"
	"github.com/giskook/transfer/pkg"
)

func event_handler_down_tranfer(c *gotcp.Conn, p *pkg.TransparentTransmissionPacket) {
	connection := c.GetExtraData().(*conn.Conn)
	if connection != nil {
		peer_id := connection.PeerID
		peer_conn := conn.NewConnsUpstream().GetConn(peer_id)
		if peer_conn != nil {
			peer_conn.SendToTerm(p)
		}
	}
}
