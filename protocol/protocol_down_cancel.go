package protocol

import (
	"bytes"
	"github.com/giskook/transfer/base"
)

type DownCancelPacket struct {
	RegisterID uint64
	SerialID   uint16
	Result     uint8
}

func (p *DownCancelPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_DOWN_REP_CANCEL, p.RegisterID, p.SerialID)
	writer.WriteByte(p.Result)
	base.WriteLength(&writer)
	base.WriteWord(&writer, CRC_ISO13239(writer.Bytes()[1:], uint16(writer.Len()-1)))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseDownCancel(buffer []byte) *DownCancelPacket {
	_, _, _, tid, serial := ParseHeader(buffer)

	return &DownCancelPacket{
		RegisterID: tid,
		SerialID:   serial,
	}
}
