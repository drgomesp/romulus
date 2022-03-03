package kro

import (
	"log"

	"go.uber.org/zap"

	"github.com/drgomesp/rhizom/internal/roprotocol/kro/handler"
	"github.com/drgomesp/rhizom/pkg/romulus/net"
)

var (
	clientHandlers = map[net.PacketID]net.PacketHandlerFunc{
		PacketClientPing:                handler.HandlePing,
		PacketClientAccountLogin:        handler.HandleAccountLogin,
		PacketClientCharEnter:           handler.HandleCharEnter,
		PacketClientCharSelectChar:      handler.HandleCharSelectChar,
		PacketClientCharRequestCharList: handler.HandleCharRequestCharList,
	}
)

func ClientPacketHandler(logger *zap.SugaredLogger) net.ProtocolHandlerFunc {
	return func(rw net.PacketExchange) error {
		go readPackets(logger, rw)
		return nil
	}
}

func readPackets(logger *zap.SugaredLogger, rw net.PacketExchange) {
	for {
		pid, clientPacket, err := rw.ReadPacket()
		if err != nil {
			log.Fatal(rw.Error(err))
			return
		}

		if pid == net.NilPacketID {
			return
		}

		h := packetHandler{
			logger:      logger,
			handlerFunc: clientHandlers[pid],
		}

		if err := h.Handle(pid, clientPacket, rw); err != nil {
			log.Fatal(rw.Error(err))
			return
		}
	}
}
