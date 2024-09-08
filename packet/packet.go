package packet

import "github.com/vmihailenco/msgpack"

type Packet interface {
	Type() Type
}

func Marshal(p Packet) ([]byte, error) {
	return msgpack.Marshal(p)
}

func Unmarshal(p []byte, t Type) (Packet, error) {
	panic("implement")
}
