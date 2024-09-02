package packet

import (
	"fmt"
	"net"
)

type Packet interface {
	Size() int
	FromBytes([]byte) error
	ToBytes() ([]byte, error)
}

func ReadPacket[P Packet](conn net.PacketConn) (p P, err error) {
	buf := make([]byte, p.Size())
	n, _, err := conn.ReadFrom(buf)
	if err != nil {
		return p, fmt.Errorf("conn.ReadFrom: %w", err)
	}
	if n != p.Size() {
		return p, fmt.Errorf("n != p.Size()")
	}

	err = p.FromBytes(buf)
	if err != nil {
		return p, fmt.Errorf("p.FromBytes: %w", err)
	}

	return p, nil
}

func SendPacket(conn net.PacketConn, addr net.Addr, p Packet) error {
	data, err := p.ToBytes()
	if err != nil {
		return fmt.Errorf("p.ToBytes: %w", err)
	}

	n, err := conn.WriteTo(data, addr)
	if err != nil {
		return fmt.Errorf("conn.WriteTo: %w", err)
	}
	if len(data) != n {
		return fmt.Errorf("len(data) != n")
	}

	return nil
}

type PacketType uint8

const (
	PacketTypeProtocolHandshakeRequest PacketType = 1
	PacketTypeProtocolHandshakeResponse PacketType = 2
)

func (pt PacketType) String() string {
	out := ""
	switch pt {
	case PacketTypeProtocolHandshakeRequest:
		out = "PacketTypeProtocolHandshakeRequest"
	case PacketTypeProtocolHandshakeResponse:
		out = "PacketTypeProtocolHandshakeResponse"
	default:
		out = "PacketTypeUnknown"
	}

	return out
}
