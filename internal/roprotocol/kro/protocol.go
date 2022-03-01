package kro

import (
	"github.com/drgomesp/rhizom/internal/roprotocol/kro/handler"
	"github.com/drgomesp/rhizom/pkg/romulus/net"
)

const (
	PacketTypeAccountLogin       = net.PacketID(0x0064)
	PacketTypeAccountLoginAccept = net.PacketID(0x0ac4)
)

var (
	clientHandlers = map[net.PacketID]net.PacketHandlerFunc{
		PacketTypeAccountLogin: handler.HandleAccountLogin,
	}
)

func ClientPacketHandler() net.ProtocolHandlerFunc {
	return func(rw net.PacketExchange) error {
		packetID, clientPacket, err := rw.ReadPacket()
		if err != nil {
			return err
		}

		err = clientHandlers[packetID](clientPacket, rw)
		if err != nil {
			return err
		}

		return nil
	}
}
