package packet

import "github.com/drgomesp/rhizom/pkg/romulus/net"

type CharListNotify struct {
	Count uint16
}

func (r *CharListNotify) Encode(p *net.PacketData) error {
	_ = p.Write(r.Count)

	return nil
}
