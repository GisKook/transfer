package event_handler

import (
	"github.com/giskook/gotcp"
	"github.com/giskook/transfer/conn"
	"github.com/giskook/transfer/pkg"
	"github.com/giskook/transfer/protocol"
	"log"
)

func event_handler_down_tranfer(c *gotcp.Conn, p *pkg.TransparentTransmissionPacket) {
	connection := c.GetExtraData().(*conn.Conn)
	if connection != nil {
		peer_id := connection.PeerID
		peer_conn := conn.NewConnsUpstream().GetConn(peer_id)
		if peer_conn != nil {
			err := peer_conn.SendToTerm(p)
			if err == nil {
				tt_pkg := p.Packet.(*protocol.DownTransferPacket)
				peer_conn.SendByteCount += tt_pkg.Len()
			}
		} else {
			log.Printf("peerid is nil %d\n", peer_id)
		}
	}
}
