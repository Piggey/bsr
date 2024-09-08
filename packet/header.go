package packet

import "fmt"

const HeaderSize = 3

type Header struct {
	Type Type
	Size int // uint16 is being sent
}

func MarshalHeader(h Header) ([]byte, error) {
	if h.Size > 0xffff {
		return nil, fmt.Errorf("packet size too large")
	}

	return []byte{
		byte(h.Type),
		byte(h.Size >> 8), // hi
		byte(h.Size),      // lo
	}, nil
}

func UnmarshalHeader(p []byte) (Header, error) {
	if len(p) != HeaderSize {
		return Header{}, fmt.Errorf("len(p) != HeaderSize")
	}

	return Header{
		Type: Type(p[0]),
		Size: int(p[1])<<8 | int(p[2]),
	}, nil
}
