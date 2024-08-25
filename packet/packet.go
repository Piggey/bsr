package packet

type Packet interface {
	ToBytes() ([]byte, error)
	FromBytes(b []byte) error
	Validate() error
}
