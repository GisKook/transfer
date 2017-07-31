package protocol

import (
//"bytes"
//"github.com/giskook/transfer/base"
)

type UpTransferPacket struct {
	Value []byte
}

func (p *UpTransferPacket) Serialize() []byte {
	return p.Value
}

func (p *UpTransferPacket) Len() uint32 {
	return uint32(len(p.Value))
}

//func (p *UpTransferPacket) Serialize() []byte {
//	//return p.Value
//	var writer bytes.Buffer
//	WriteHeader(&writer, 0,
//		PROTOCOL_DOWN_TRANSFER, p.ID, p.SerialID)
//	writer.Write(p.Value)
//
//	base.WriteLength(&writer)
//	base.WriteWord(&writer, CRC_ISO13239(writer.Bytes()[1:], uint16(writer.Len()-1)))
//	writer.WriteByte(PROTOCOL_END_FLAG)
//
//	return writer.Bytes()
//}

func ParseUpTransfer(buffer []byte) *UpTransferPacket {
	return &UpTransferPacket{
		Value: buffer,
	}
}

//func ParseUpTransfer(buffer []byte) *UpTransferPacket {
//	reader, length, _, tid, serial := ParseHeader(buffer)
//	value := make([]byte, length-PROTOCOL_COMMON_LEN)
//	reader.Read(value)
//
//	return &UpTransferPacket{
//		ID:       tid,
//		SerialID: serial,
//		Value:    value,
//	}
//
//}
