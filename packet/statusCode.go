package packet

type StatusCode uint8

const (
	StatusOk                   = 0
	StatusProtocolNotSupported = 1
	StatusIncorrectGameDetails = 2
)
