package protocol

import (
	"bytes"
	"github.com/giskook/transfer/base"
)

type UpCancelPacket struct {
	ID       uint64
	SerialID uint16
}

func (p *UpCancelPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_UP_REP_CANCEL, p.ID, p.SerialID)
	writer.WriteByte(0)
	base.WriteLength(&writer)
	base.WriteWord(&writer, CRC_ISO13239(writer.Bytes()[1:], uint16(writer.Len()-1)))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseUpCancel(buffer []byte) *UpCancelPacket {
	_, _, _, tid, serial := ParseHeader(buffer)

	return &UpCancelPacket{
		ID:       tid,
		SerialID: serial,
	}
}
