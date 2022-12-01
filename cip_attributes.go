package gologix

import "encoding/binary"

// Here are the objects

type CIPAttribute uint16

// Used to indicate how many bytes are used for the data. If they are more than 8 bits,
// they actually actually take n+1 bytes.  First byte after specifier is a 0
type CIPAttributeType byte

const (
	CIPAttribute_8bit  CIPAttributeType = 0x30
	CIPAttribute_16bit CIPAttributeType = 0x31
)

func (p CIPAttribute) Bytes() []byte {
	if p < 256 {
		b := make([]byte, 2)
		b[0] = byte(CIPAttribute_8bit)
		b[1] = byte(p)
		return b
	} else {

		b := make([]byte, 4)
		b[0] = byte(CIPAttribute_16bit)
		binary.LittleEndian.PutUint16(b[2:], uint16(p))
		return b
	}
}

const (
	CIPAttribute_Data CIPAttribute = 0x03
)