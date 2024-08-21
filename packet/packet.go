package packet

import (
	"fmt"
	"net"

	"github.com/Piggey/bsr/packet/binary"
)

type Packet interface {
	ToBytes() ([]byte, error)
	FromBytes(b []byte) error
	Validate() error
}

func ReadPacket[T Packet](sender *net.UDPConn) (T, net.Addr, error) {
	var p T
	addr, err := binary.ReadFrom(sender, byteOrder, &p)
	if err != nil {
		return p, nil, fmt.Errorf("binary.ReadFrom: %w", err)
	}

	if err := p.Validate(); err != nil {
		return p, addr, fmt.Errorf("p.Validate: %w", err)
	}

	return p, addr, nil
}

func WritePacket(sender *net.UDPConn, addr net.Addr, packet Packet) error {
	err := binary.Write(sender, byteOrder, packet)
	if err != nil {
		return fmt.Errorf("binary.Write: %w", err)
	}

	return nil
}
