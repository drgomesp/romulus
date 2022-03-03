package roserver

import (
	"context"
	"strconv"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/drgomesp/rhizom/internal/roprotocol/kro"
	"github.com/drgomesp/rhizom/pkg/romulus/net"
)

const (
	DefaultServerPort    = 6900
	DefaultClientVersion = 20211117
)

type Server struct {
	*net.Server

	logger *zap.SugaredLogger
}

func NewServer(logger *zap.SugaredLogger) (*Server, error) {
	packetRegistry := net.NewPacketRegistry(DefaultClientVersion, kro.ClientPackets, kro.ServerPackets)

	netServer, err := net.NewServer(
		net.Config{Port: strconv.Itoa(DefaultServerPort)},
		net.WithLogger(logger),
		net.WithPacketRegistry(packetRegistry),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize account server")
	}

	srv := &Server{
		Server: netServer,
		logger: logger,
	}

	netServer.RegisterAPIs(srv.APIs())
	netServer.RegisterProtocols(srv.Protocols()...)
	netServer.RegisterServices(srv)

	return srv, nil
}

func (a *Server) Start(ctx context.Context) error {
	return a.Server.Start(ctx)
}

func (a *Server) APIs() []interface{} {
	return []interface{}{}
}

func (a *Server) Protocols() []*net.Protocol {
	return []*net.Protocol{
		{
			Run: kro.ClientPacketHandler(a.logger),
		},
	}
}
