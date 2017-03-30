package pkg

import (
	"github.com/giskook/gotcp"
)

type TransparentTransmissionPacket struct {
	Type   uint16
	Packet gotcp.Packet
}

func (this *TransparentTransmissionPacket) Serialize() []byte {
	return this.Packet.Serialize()
}

func NewTransparentTransmissionPakcet(Type uint16, Packet gotcp.Packet) *TransparentTransmissionPacket {
	return &TransparentTransmissionPacket{
		Type:   Type,
		Packet: Packet,
	}
}
