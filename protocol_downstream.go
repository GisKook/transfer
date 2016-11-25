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

	buffer := smconn.GetBuffer()
	conn := c.GetRawConn()
	for {
		if smconn.ReadMore {
			data := make([]byte, 2048)
			readLengh, err := conn.Read(data)
			log.Printf("<Down IN>  %x\n", data[0:readLengh])
			if err != nil {
				return nil, err
			}

			if readLengh == 0 {
				return nil, gotcp.ErrConnClosing
			}
			buffer.Write(data[0:readLengh])

			//	return &DownstreamPacket{
			//		buf: data[0:readLengh],
			//	}, nil
		}
		cmdid, pkglen := CheckProtocol(buffer)
		log.Printf("protocol id %d\n", cmdid)

		pkgbyte := make([]byte, pkglen)
		buffer.Read(pkgbyte)
		switch cmdid {
		case PROTOCOL_OK:
			return &DownstreamPacket{
				buf: DelPackage(pkgbyte),
			}, nil
			smconn.ReadMore = false
		case PROTOCOL_ILLEGAL:
			smconn.ReadMore = true
		case PROTOCOL_HALF_PACK:
			smconn.ReadMore = true
		}
	}
}
