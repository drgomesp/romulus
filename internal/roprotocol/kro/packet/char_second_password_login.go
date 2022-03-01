package packet

import (
	"github.com/drgomesp/rhizom/pkg/romulus/net"
)

type CharSecondPasswordLogin struct {
	Seed      uint32
	AccountID uint32
	Result    uint16
}

func (r *CharSecondPasswordLogin) Encode(p *net.PacketData) error {
	p.Write(r.Seed)
	p.Write(r.AccountID)
	p.Write(r.Result)

	return nil
}
