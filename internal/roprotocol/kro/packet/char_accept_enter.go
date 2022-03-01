package packet

import "github.com/drgomesp/rhizom/pkg/romulus/net"

type CharAcceptEnter struct {
	TotalSlotCount   byte
	PremiumSlotStart byte
	PremiumSlotEnd   byte
	Chars            []*CharacterInfo
}

func (r *CharAcceptEnter) Encode(p *net.PacketData) error {
	p.Grow(23)

	p.Write(r.TotalSlotCount)
	p.Write(r.PremiumSlotStart)
	p.Write(r.PremiumSlotEnd)
	p.Skip(20)

	for _, ch := range r.Chars {
		ch.Encode(p)
	}

	return nil
}
