package transfer

import (
	"bytes"
	"encoding/binary"
	"github.com/giskook/transfer/base"
	"log"
)

const POLYNOMIAL uint16 = 0x8772
const PRESET_VALUE uint16 = 0xFFFF
const CHECK_VALUE uint16 = 0xF0B8

func CRC_ISO13239(buffer []byte, Len int) uint16 {
	current_crc_value := uint32(PRESET_VALUE)
	for i := 0; i < Len; i++ {
		current_crc_value = current_crc_value ^ (uint32(buffer[i]))
		for j := 0; j < 8; j++ {
			if current_crc_value&0x0001 != 0 {
				current_crc_value = (current_crc_value >> 1) ^ uint32(POLYNOMIAL)
			} else {
				current_crc_value = (current_crc_value >> 1)
			}
		}
	}

	return uint16((^current_crc_value)) & 0xFFFF
}

func AddPackage(buffer []byte) []byte {
	protocol_len := len(buffer) + 6
	head := make([]byte, 3)
	head[0] = 0x55
	binary.LittleEndian.PutUint16(head[1:3], uint16(protocol_len))
	result := append(head, buffer...)

	crc := CRC_ISO13239(result, len(result))

	crc_byte := make([]byte, 3)
	binary.LittleEndian.PutUint16(crc_byte[0:2], crc)
	crc_byte[2] = 0xaa

	result = append(result, crc_byte...)

	return result
}

func DelPackage(buffer []byte) []byte {
	length := len(buffer)
	result := buffer[3 : length-3]
	return result
}

const PROTOCOL_ILLEGAL uint16 = 255
const PROTOCOL_HALF_PACK uint16 = 254
const PROTOCOL_OK uint16 = 1

func CheckProtocol(buffer *bytes.Buffer) (uint16, uint16) {
	bufferlen := buffer.Len()
	if bufferlen == 0 {
		return PROTOCOL_ILLEGAL, 0
	}
	if buffer.Bytes()[0] != 0x55 {
		buffer.ReadByte()
		CheckProtocol(buffer)
	} else if bufferlen > 2 {
		pkglen := base.GetWord(buffer.Bytes()[1:3])
		log.Printf("pkglen %d\n", pkglen)
		if pkglen < 5 || pkglen > 2048 {
			buffer.ReadByte()
			CheckProtocol(buffer)
		}

		if int(pkglen) > bufferlen {
			return PROTOCOL_HALF_PACK, 0
		} else {
			crc_calc := CRC_ISO13239(buffer.Bytes(), int(pkglen-3))
			log.Printf("crc value %x\n", crc_calc)
			log.Printf("crc in protocol %x\n", base.GetWord(buffer.Bytes()[pkglen-3:pkglen-1]))
			if crc_calc == base.GetWord(buffer.Bytes()[pkglen-3:pkglen-1]) && buffer.Bytes()[pkglen-1] == 0xaa {
				return PROTOCOL_OK, pkglen
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
