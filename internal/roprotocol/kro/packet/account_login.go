package packet

import "github.com/drgomesp/rhizom/pkg/romulus/net"

type AccountLogin struct {
	Version    uint32
	Username   string
	Password   string
	ClientType byte
}

func (p *AccountLogin) Decode(data *net.PacketData) error {
	data.Read(&p.Version)
	data.ReadString(24, &p.Username)
	data.ReadString(24, &p.Password)
	data.Read(&p.ClientType)

	return nil
}
