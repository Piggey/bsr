package packet

import "fmt"

const HeaderSize = 3

type Header struct {
	Type Type
	Size uint16
}

func UnmarshalHeader(p []byte) (Header, error) {
	if len(p) != HeaderSize {
		return Header{}, fmt.Errorf("len(p) != HeaderSize")
	}

	return Header{
		Type: Type(p[0]),
		Size: uint16(p[1])<<8 | uint16(p[2]),
	}, nil
}
