package roserver

import (
	"context"
	"strconv"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/drgomesp/rhizom/internal"
	"github.com/drgomesp/rhizom/internal/roprotocol/kro"
	"github.com/drgomesp/rhizom/pkg/romulus/net"
)

type Account struct {
	*net.Server

	logger *zap.SugaredLogger
}

func NewAccountServer(logger *zap.SugaredLogger) (*Account, error) {
	packetRegistry := net.NewPacketRegistry(20211117, kro.ClientPackets, kro.ServerPackets)

	srv, err := net.NewServer(
		net.Config{Port: strconv.Itoa(internal.AccountServerDefaultPort)},
		net.WithLogger(logger),
		net.WithPacketRegistry(packetRegistry),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize account server")
	}

	accountServer := &Account{
		Server: srv,
		logger: logger,
	}

	srv.RegisterAPIs(accountServer.APIs())
	srv.RegisterProtocols(accountServer.Protocols()...)
	srv.RegisterServices(accountServer)

	return accountServer, nil
}

func (a *Account) Start(ctx context.Context) error {
	if err := a.Server.Start(ctx); err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return a.Stop(ctx)
		default:
			{
			}
		}
	}
}

func (a *Account) APIs() []interface{} {
	return []interface{}{}
}

func (a *Account) Protocols() []*net.Protocol {
	return []*net.Protocol{
		{
			Run: kro.ClientPacketHandler(),
		},
	}
}
