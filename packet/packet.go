package packet

type Packet interface {
	Validate() error
}
