package handler

import (
	"errors"

	"github.com/drgomesp/rhizom/internal/roprotocol/kro/packet"
	"github.com/drgomesp/rhizom/pkg/romulus/net"
)

func HandleCharRequestCharList(pid net.PacketID, clientPacket net.ClientPacket, sender net.PacketWriter) error {
	p, ok := clientPacket.(*packet.CharRequestCharList)
	if !ok {
		return errors.New("fuk!")
	}

	_ = p
	return nil
}
