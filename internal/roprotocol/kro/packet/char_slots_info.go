package packet

import (
	"log"

	"github.com/drgomesp/rhizom/pkg/romulus/net"
)

type CharSlotsInfo struct {
	NormalSlots     byte
	PremiumSlots    byte
	BillingSlots    byte
	ProducibleSlots byte
	ValidSlots      byte
	Chars           []*CharacterInfo
}

func (r *CharSlotsInfo) Encode(p *net.PacketData) error {
	p.Grow(25)

	p.Write(r.NormalSlots)
	p.Write(r.PremiumSlots)
	p.Write(r.BillingSlots)
	p.Write(r.ProducibleSlots)
	p.Write(r.ValidSlots)
	p.Skip(20)

	for _, ch := range r.Chars {
		if err := ch.Encode(p); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
