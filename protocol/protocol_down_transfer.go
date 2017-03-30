package protocol

import (
	"bytes"
	"github.com/giskook/transfer/base"
)

type DownTransferPacket struct {
	RouterRegisterID uint64
	SerialID         uint16
	Value            []byte
}

func (p *DownTransferPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_UP_TRANSFER, p.RouterRegisterID, p.SerialID)
	writer.Write(p.Value)

	base.WriteLength(&writer)
	base.WriteWord(&writer, CRC_ISO13239(writer.Bytes()[1:], uint16(writer.Len()-1)))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseDownTransfer(buffer []byte) *DownTransferPacket {
	reader, length, _, tid, serial := ParseHeader(buffer)
	value := make([]byte, length-PROTOCOL_COMMON_LEN)
	reader.Read(value)

	return &DownTransferPacket{
		RouterRegisterID: tid,
		SerialID:         serial,
		Value:            value,
	}

}
