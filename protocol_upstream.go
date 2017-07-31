package transfer

import (
	"github.com/giskook/gotcp"
	"github.com/giskook/transfer/conn"
	"github.com/giskook/transfer/pkg"
	"github.com/giskook/transfer/protocol"
	"log"
	"sync"
)

type UpstreamProtocol struct {
}

func (this *UpstreamProtocol) ReadPacket(c *gotcp.Conn) (gotcp.Packet, error) {
	smconn := c.GetExtraData().(*conn.Conn)
	var once sync.Once
	once.Do(smconn.UpdateReadflag)

	buffer := smconn.GetBuffer()
	conn := c.GetRawConn()
	for {
		if smconn.ReadMore {
			data := make([]byte, 2048)
			readLengh, err := conn.Read(data)
			log.Printf("<Up IN>  %x\n", data[0:readLengh])
			if err != nil {
				return nil, err
			}

			if readLengh == 0 {
				return nil, gotcp.ErrConnClosing
			}
			buffer.Write(data[0:readLengh])

			//		return &DownstreamPacket{
			//			buf: data[0:readLengh],
			//		}, nil
		}
		cmdid, pkglen := protocol.CheckProtocol(buffer)
		log.Printf("protocol id %d\n", cmdid)

		pkgbyte := make([]byte, pkglen)
		buffer.Read(pkgbyte)
		switch cmdid {
		case protocol.PROTOCOL_UP_REQ_REGISTER:
			p := protocol.ParseUpRegister(pkgbyte)
			smconn.ReadMore = false

			return pkg.NewTransparentTransmissionPakcet(cmdid, p), nil
		case protocol.PROTOCOL_UP_REQ_CANCEL:
			p := protocol.ParseUpCancel(pkgbyte)
			smconn.ReadMore = false

			return pkg.NewTransparentTransmissionPakcet(cmdid, p), nil

		case protocol.PROTOCOL_UP_TRANSFER:
			p := protocol.ParseUpTransfer(pkgbyte)
			smconn.ReadMore = false
			smconn.RecvByteCount += uint32(pkglen)

			return pkg.NewTransparentTransmissionPakcet(cmdid, p), nil

		case protocol.PROTOCOL_ILLEGAL:
			smconn.ReadMore = true
		case protocol.PROTOCOL_HALF_PACK:
			smconn.ReadMore = true
		}
	}
}
