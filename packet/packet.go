package packet

type Packet interface {
	Size() int
	FromBytes([]byte) error
	ToBytes() ([]byte, error)
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
