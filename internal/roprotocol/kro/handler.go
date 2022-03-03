package kro

import (
	"go.uber.org/zap"

	"github.com/drgomesp/rhizom/pkg/romulus/net"
)

type packetHandler struct {
	logger      *zap.SugaredLogger
	handlerFunc net.PacketHandlerFunc
}

func (h *packetHandler) Handle(pid net.PacketID, p net.ClientPacket, rw net.PacketWriter) error {
	if err := h.handlerFunc(pid, p, rw); err != nil {
		h.logger.Error(err)
		return err
	}

	return nil
}
