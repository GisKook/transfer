package transfer

import (
	"github.com/giskook/gotcp"
	"log"
)

type DownstreamPacket struct {
	buf []byte
}

func (this *DownstreamPacket) Serialize() []byte {
	return this.buf
}

type DownstreamProtocol struct {
}

func (this *DownstreamProtocol) ReadPacket(c *gotcp.Conn) (gotcp.Packet, error) {
	smconn := c.GetExtraData().(*Conn)
	smconn.UpdateReadflag()

	conn := c.GetRawConn()
	for {
		data := make([]byte, 2048)
		readLengh, err := conn.Read(data)
		log.Printf("<Down IN>  %x\n", data[0:readLengh])
		if err != nil {
			return nil, err
		}

		if readLengh == 0 {
			return nil, gotcp.ErrConnClosing
		}

		return &DownstreamPacket{
			buf: data[0:readLengh],
		}, nil
	}
}
