package packet

import (
	"encoding/binary"
	gonet "net"

	"github.com/drgomesp/rhizom/pkg/romulus/net"
)

type AccountAcceptLogin struct {
	AuthCode     uint32
	AccountID    uint32
	AccountLevel uint32
	Sex          byte
	Servers      []*CharServer
}

type CharServer struct {
	IP       gonet.IP
	Port     uint16
	Name     string
	Users    uint16
	State    uint16
	Property uint16
}

func (p *AccountAcceptLogin) Encode(data *net.PacketData) error {
	data.Grow(60 + (len(p.Servers)+128)*32)

	_ = data.Write(p.AuthCode)
	_ = data.Write(p.AccountID)
	_ = data.Write(p.AccountLevel)

	data.Skip(30)
	_ = data.Write(p.Sex)
	data.Skip(17)

	for _, srv := range p.Servers {
		_ = data.Write(binary.LittleEndian.Uint32(srv.IP.To4()))
		_ = data.Write(srv.Port)
		_ = data.WriteString(20, srv.Name)
		_ = data.Write(srv.Users)
		_ = data.Write(srv.State)
		_ = data.Write(srv.Property)
		_ = data.WriteString(128, "")
	}

	return nil
}
