package packet

import (
	"github.com/drgomesp/rhizom/pkg/romulus/net"
)

type Ping struct {
	AccountID uint32
}

func (r *Ping) Decode(p *net.PacketData) error {
	p.Read(&r.AccountID)

	return nil
}
