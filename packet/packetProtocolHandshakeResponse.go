package packet

import (
	"fmt"
	"unsafe"
)

const (
	HandshakeStatusOK = 1
	HandshakeStatusProtocolNotSupported = 200
)

type PacketProtocolHandshakeResponse struct {
	Status uint8
}

func (pph PacketProtocolHandshakeResponse) Size() int {
	return int(unsafe.Sizeof(pph))
}

func (pph *PacketProtocolHandshakeResponse) FromBytes(data []byte) error {
	if len(data) != pph.Size() {
		return fmt.Errorf("len(data) != pph.Size()")
	}

	pph.Status = data[0]
	return nil
}

func (pph *PacketProtocolHandshakeResponse) ToBytes() ([]byte, error) {
	return []byte{pph.Status}, nil
}
