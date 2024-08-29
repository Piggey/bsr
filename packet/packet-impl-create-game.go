package packet

import (
	"fmt"
	"slices"

	"google.golang.org/protobuf/proto"
)

func (cgp *CreateGamePacket) ToBytes() ([]byte, error) {
	return proto.Marshal(cgp)
}

func (cgp *CreateGamePacket) FromBytes(b []byte) error {
	return proto.Unmarshal(b, cgp)
}

func (cgp *CreateGamePacket) Validate() error {
	if !slices.Equal(cgp.GetMagic(), magicBytes) {
		return fmt.Errorf("CreateGamePacket: incorrect header")
	}

	if cgp.GetVersion() != protocolVersion1 {
		return fmt.Errorf("CreateGamePacket: incorrect version")
	}

	return nil
}
