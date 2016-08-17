package transfer

import (
	"github.com/giskook/gotcp"
	"log"
)

type UpstreamPacket struct {
	buf []byte
}

func (this *UpstreamPacket) Serialize() []byte {
	return this.buf
}

type UpstreamProtocol struct {
}

func (this *UpstreamProtocol) ReadPacket(c *gotcp.Conn) (gotcp.Packet, error) {
	smconn := c.GetExtraData().(*Conn)
	smconn.UpdateReadflag()

	conn := c.GetRawConn()
	for {
		data := make([]byte, 2048)
		readLengh, err := conn.Read(data)
		log.Printf("<UP IN>  %x\n", data[0:readLengh])
		if err != nil {
			return nil, err
		}

		if readLengh == 0 {
			return nil, gotcp.ErrConnClosing
		}

		return &UpstreamPacket{
			buf: data[0:readLengh],
		}, nil
	}
}
