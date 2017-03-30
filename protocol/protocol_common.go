package protocol

import (
	"bytes"
	//"encoding/binary"
	"github.com/giskook/transfer/base"
	"log"
)

const (
	PROTOCOL_START_FLAG uint8  = 0x55
	PROTOCOL_END_FLAG   uint8  = 0xaa
	PROTOCOL_ILLEGAL    uint16 = 255
	PROTOCOL_HALF_PACK  uint16 = 254
	PROTOCOL_COMMON_LEN uint16 = 18

	POLYNOMIAL   uint16 = 0x8772
	PRESET_VALUE uint16 = 0xFFFF
	CHECK_VALUE  uint16 = 0xF0B8
)

func CRC_ISO13239(buffer []byte, Len uint16) uint16 {
	current_crc_value := uint32(PRESET_VALUE)
	for i := 0; i < int(Len); i++ {
		current_crc_value = current_crc_value ^ (uint32(buffer[i]))
		for j := 0; j < 8; j++ {
			if current_crc_value%2 == 1 {
				current_crc_value = (current_crc_value >> 1) ^ uint32(POLYNOMIAL)
			} else {
				current_crc_value = (current_crc_value >> 1)
			}
		}
	}

	return uint16((^current_crc_value)) & 0xFFFF
}

func AddPackage(buffer []byte) []byte {
	return buffer
	//	protocol_len := len(buffer) + 6
	//	head := make([]byte, 3)
	//	head[0] = 0x55
	//	binary.LittleEndian.PutUint16(head[1:3], uint16(protocol_len))
	//	result := append(head, buffer...)
	//
	//	crc := CRC_ISO13239(result, len(result))
	//
	//	crc_byte := make([]byte, 3)
	//	binary.LittleEndian.PutUint16(crc_byte[0:2], crc)
	//	crc_byte[2] = 0xaa
	//
	//	result = append(result, crc_byte...)
	//
	//	return result
}

func DelPackage(buffer []byte) []byte {
	return buffer
	//	length := len(buffer)
	//	result := buffer[3 : length-3]
	//	return result
}

func CheckProtocol(buffer *bytes.Buffer) (uint16, uint16) {
	bufferlen := buffer.Len()
	if bufferlen == 0 {
		return PROTOCOL_ILLEGAL, 0
	}
	if buffer.Bytes()[0] != PROTOCOL_START_FLAG {
		buffer.ReadByte()
		CheckProtocol(buffer)
	} else if bufferlen > 2 {
		pkglen := base.GetWord(buffer.Bytes()[1:3])
		//log.Printf("pkglen %d\n", pkglen)
		if pkglen < 5 || pkglen > 2048 {
			buffer.ReadByte()
			CheckProtocol(buffer)
		}

		if int(pkglen) > bufferlen {
			return PROTOCOL_HALF_PACK, 0
		} else {
			crc_calc := CRC_ISO13239(buffer.Bytes()[1:], pkglen-4)
			log.Printf("crc value %x\n", crc_calc)
			log.Printf("crc in protocol %x\n", base.GetWord(buffer.Bytes()[pkglen-3:pkglen-1]))
			if crc_calc == base.GetWord(buffer.Bytes()[pkglen-3:pkglen-1]) && buffer.Bytes()[pkglen-1] == PROTOCOL_END_FLAG {
				protocol_id := base.GetWord(buffer.Bytes()[3:5])
				return protocol_id, pkglen
			} else {
				buffer.ReadByte()
				CheckProtocol(buffer)
			}
		}
	} else {
		return PROTOCOL_HALF_PACK, 0
	}

	return PROTOCOL_HALF_PACK, 0
}

func WriteHeader(writer *bytes.Buffer, length uint16, cmdid uint16, register_id uint64, serial_id uint16) {
	writer.WriteByte(PROTOCOL_START_FLAG)
	base.WriteWord(writer, length)
	base.WriteWord(writer, cmdid)
	base.WriteQuaWord(writer, register_id)
	base.WriteWord(writer, serial_id)
}

func ParseHeader(buffer []byte) (*bytes.Reader, uint16, uint16, uint64, uint16) {
	reader := bytes.NewReader(buffer)
	reader.Seek(1, 0)
	length := base.ReadWord(reader)
	protocol_id := base.ReadWord(reader)
	register_id := base.ReadQuaWord(reader)
	serial_id := base.ReadWord(reader)

	return reader, length, protocol_id, register_id, serial_id
}
