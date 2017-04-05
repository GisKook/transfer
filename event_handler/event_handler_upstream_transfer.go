package event_handler

import (
	"github.com/giskook/gotcp"
	"github.com/giskook/transfer/conn"
	"github.com/giskook/transfer/pkg"
	"log"
)

func event_handler_up_transfer(c *gotcp.Conn, p *pkg.TransparentTransmissionPacket) {
	connection := c.GetExtraData().(*conn.Conn)
	if connection != nil {
		peer_id := connection.PeerID
		log.Printf("peer id %d\n", peer_id)
		peer_conn := conn.NewConnsDownstream().GetConn(peer_id)
		if peer_conn != nil {
			peer_conn.SendToTerm(p)
		} else {
			log.Println("peer connection is nil")
		}
	} else {
		log.Println("connection is nil")
	}
}
