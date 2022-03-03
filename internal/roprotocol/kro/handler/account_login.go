package handler

import (
	gonet "net"

	"github.com/pkg/errors"

	"github.com/drgomesp/rhizom/internal/roprotocol/kro/packet"
	"github.com/drgomesp/rhizom/pkg/romulus/net"
)

func HandleAccountLogin(pid net.PacketID, p net.ClientPacket, sender net.PacketWriter) error {
	p, ok := p.(*packet.AccountLogin)
	if !ok {
		return errors.New("fuk!")
	}

	err := sender.Send(&packet.AccountAcceptLogin{
		AuthCode:     0xdeadbeef,
		AccountID:    2000000,
		AccountLevel: 99,
		Sex:          0,
		Servers: []*packet.CharServer{
			{
				IP:       gonet.ParseIP("127.0.0.1"),
				Port:     6900,
				Name:     "ROmulus (dev)",
				Users:    0,
				State:    5,
				Property: 5,
			},
		},
	})

	if err != nil {
		return err
	}

	return nil
}
