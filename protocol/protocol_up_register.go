package protocol

import (
	"bytes"
	"github.com/giskook/transfer/base"
)

type UpRegisterPacket struct {
	ID                         uint64
	SerialID                   uint16
	PeerRouterRegisterID       uint64
	TransparentTransmissionKey uint32
	Status                     uint8
}

func (p *UpRegisterPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_UP_REP_REGISTER, p.ID, p.SerialID)
	writer.WriteByte(p.Status)
	base.WriteDWord(&writer, 0)
	base.WriteLength(&writer)
	base.WriteWord(&writer, CRC_ISO13239(writer.Bytes()[1:], uint16(writer.Len()-1)))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseUpRegister(buffer []byte) *UpRegisterPacket {
	reader, _, _, tid, serial := ParseHeader(buffer)
	peer_router_register_id := base.ReadQuaWord(reader)
	transparent_transmission_key := base.ReadDWord(reader)

	return &UpRegisterPacket{
		ID:                         tid,
		SerialID:                   serial,
		PeerRouterRegisterID:       peer_router_register_id,
		TransparentTransmissionKey: transparent_transmission_key,
	}
}
