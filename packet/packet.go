package packet

var (
	magicBytes = []byte{'b', 's', 'r'}
)

const (
	protocolVersion1 uint32 = 1
)

type Packet interface {
	ToBytes() ([]byte, error)
	FromBytes(b []byte) error
	Validate() error
}
