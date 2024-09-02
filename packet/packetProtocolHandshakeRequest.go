package packet

import (
	"fmt"
	"unsafe"
)

type PacketProtocolHandshakeRequest struct {
	ProtocolVersion uint8
}

func (pph PacketProtocolHandshakeRequest) Size() int {
	return int(unsafe.Sizeof(pph))
}

func (pph *PacketProtocolHandshakeRequest) FromBytes(data []byte) error {
	if len(data) != pph.Size() {
		return fmt.Errorf("len(data) != pph.Size()")
	}

	pph.ProtocolVersion = data[0]
	return nil
}

func (pph *PacketProtocolHandshakeRequest) ToBytes() ([]byte, error) {
	return []byte{pph.ProtocolVersion}, nil
}
