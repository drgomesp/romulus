package packet

import "github.com/drgomesp/rhizom/pkg/romulus/net"

type CharRequestCharList struct{}

func (*CharRequestCharList) Decode(_ *net.PacketData) error {
	return nil
}
