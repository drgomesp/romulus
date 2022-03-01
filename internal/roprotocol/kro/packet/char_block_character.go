package packet

import "github.com/drgomesp/rhizom/pkg/romulus/net"

type CharBlockCharacter struct {
}

func (r *CharBlockCharacter) Encode(p *net.PacketData) error {
	return nil
}
