package packet

type Packet interface {
	ToBytes() []byte
}
