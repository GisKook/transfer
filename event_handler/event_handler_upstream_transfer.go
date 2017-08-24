package event_handler

import (
	"github.com/giskook/gotcp"
	"github.com/giskook/transfer/conn"
	"github.com/giskook/transfer/pkg"
	"github.com/giskook/transfer/protocol"
	"log"
)

func event_handler_up_transfer(c *gotcp.Conn, p *pkg.TransparentTransmissionPacket) {
	connection := c.GetExtraData().(*conn.Conn)
	if connection != nil {
		peer_id := connection.PeerID
		log.Printf("peer id %d\n", peer_id)
		peer_conn := conn.NewConnsDownstream().GetConn(peer_id)
		if peer_conn != nil {
			err := peer_conn.SendToTerm(p)
			if err == nil {
				tt_pkg := p.Packet.(*protocol.UpTransferPacket)
				peer_conn.SendByteCount += tt_pkg.Len()
			}
		} else {
			log.Println("peer connection is nil")
		}
	} else {
		log.Println("connection is nil")
	}
}
