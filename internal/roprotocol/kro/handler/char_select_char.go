package handler

import (
	"log"

	"github.com/drgomesp/rhizom/internal/roprotocol/kro/packet"
	"github.com/drgomesp/rhizom/pkg/romulus/net"
)

func HandleCharSelectChar(pid net.PacketID, clientPacket net.ClientPacket, sender net.PacketWriter) error {
	p, ok := clientPacket.(*packet.CharSelectChar)
	if !ok {
		return net.ErrPacketType(pid)
	}

	log.Println(p)

	return nil
}
