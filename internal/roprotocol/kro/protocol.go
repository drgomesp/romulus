package kro

import (
	"github.com/drgomesp/rhizom/internal/roprotocol/kro/handler"
	"github.com/drgomesp/rhizom/pkg/romulus/net"
)

var (
	clientHandlers = map[net.PacketID]net.PacketHandlerFunc{
		PacketClientAccountLogin: handler.HandleAccountLogin,
		PacketClientCharEnter:    handler.HandleCharEnter,
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
