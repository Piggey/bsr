package packet

type Packet interface {
	Type() Type
}

func Unmarshal(p []byte, t Type) (Packet, error) {
	panic("implement")
}
