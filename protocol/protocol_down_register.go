package protocol

import (
	"bytes"
	"github.com/giskook/transfer/base"
	"log"
)

type DownRegisterPacket struct {
	RegisterID                 uint64
	SerialID                   uint16
	PeerClientID               uint64
	TransparentTransmissionKey uint32
	Status                     uint8
}

func (p *DownRegisterPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_DOWN_REP_REGISTER, p.RegisterID, p.SerialID)
	writer.WriteByte(p.Status)
	base.WriteDWord(&writer, 0)
	base.WriteLength(&writer)
	base.WriteWord(&writer, CRC_ISO13239(writer.Bytes()[1:], uint16(writer.Len()-1)))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseDownRegister(buffer []byte) *DownRegisterPacket {
	log.Println("ParseDownRegister")
	reader, _, _, tid, serial := ParseHeader(buffer)
	peer_client_id := base.ReadQuaWord(reader)
	transparent_transmission_key := base.ReadDWord(reader)
	log.Println("ParseDownRegister")

	return &DownRegisterPacket{
		RegisterID:                 tid,
		SerialID:                   serial,
		PeerClientID:               peer_client_id,
		TransparentTransmissionKey: transparent_transmission_key,
	}

}
