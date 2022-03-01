package packet

import (
	"github.com/drgomesp/rhizom/pkg/romulus/net"
)

type AccountLoginRefused struct {
	Code      byte
	BlockDate string
}

func (p *AccountLoginRefused) Encode(data *net.PacketData) error {
	_ = data.Write(&p.Code)
	_ = data.WriteString(20, p.BlockDate)

	return nil
}
