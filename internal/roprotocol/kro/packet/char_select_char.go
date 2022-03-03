package packet

import "github.com/drgomesp/rhizom/pkg/romulus/net"

type CharSelectChar struct {
	Slot byte
}

func (c *CharSelectChar) Decode(p *net.PacketData) error {
	p.Read(&c.Slot)

	return nil
}
