package handler

import (
	"github.com/drgomesp/rhizom/internal/roprotocol/kro/packet"
	"github.com/drgomesp/rhizom/pkg/romulus/net"
)

func HandlePing(pid net.PacketID, p net.ClientPacket, sender net.PacketWriter) error {
	p, ok := p.(*packet.Ping)
	if !ok {
		return net.ErrPacketType(pid)
	}

	return nil
}
