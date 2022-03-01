package packet

import "github.com/drgomesp/rhizom/pkg/romulus/net"

type ClientRequestCharList struct{}

func (*ClientRequestCharList) Decode(data *net.PacketData) error {
	return nil
}
