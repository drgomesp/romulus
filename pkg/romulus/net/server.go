package net

import (
	"bufio"
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"

	"github.com/drgomesp/rhizom/pkg/romulus"
)

type ServerOption func(*Server)

type ConnHandlerFunc func(conn net.Conn) error

type Server struct {
	cfg       Config
	logger    *zap.SugaredLogger
	listener  net.Listener
	protocols []*Protocol
	handler   func(pid ProtocolID, handler ProtocolHandlerFunc) ConnHandlerFunc
	packetReg *PacketRegistry

	running bool
	quit    chan bool
}

func WithLogger(l *zap.SugaredLogger) ServerOption {
	return func(srv *Server) {
		srv.logger = l
	}
}

func WithPacketRegistry(reg *PacketRegistry) ServerOption {
	return func(srv *Server) {
		srv.packetReg = reg
	}
}

func NewServer(cfg Config, opts ...ServerOption) (*Server, error) {
	srv := &Server{
		cfg: cfg,
	}

	for _, option := range opts {
		option(srv)
	}

	return srv, nil
}

func (s *Server) Name() string {
	return "net.Server"
}

func (s *Server) Start(ctx context.Context) error {
	addr := fmt.Sprintf(":%s", s.cfg.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	s.listener = listener
	s.running = true

	s.logger.Infof("listening on %s", addr)

	return s.run()
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.listener.Close(); err != nil {
		return err
	}

	return nil
}

func (s *Server) run() error {
running:
	for s.running {
		select {
		case <-s.quit:
			break running
		default:
			conn, err := s.listener.Accept()
			if err != nil {
				s.logger.Error(err)
				continue
			}

			for _, proto := range s.protocols {
				go func(proto *Protocol, conn net.Conn) {
					if err := s.handler(proto.ID, proto.Run)(conn); err != nil {
						s.logger.Error(err)
					}
				}(proto, conn)
			}
		}
	}

	return nil
}

func (s *Server) RegisterProtocols(protocols ...*Protocol) {
	s.protocols = protocols

	s.handler = func(pid ProtocolID, handler ProtocolHandlerFunc) ConnHandlerFunc {
		return func(conn net.Conn) error {
			r, w := bufio.NewReader(conn), bufio.NewWriter(conn)

			rw := &protoRW{
				logger: s.logger,
				conn:   conn,
				stream: bufio.NewReadWriter(r, w),
				reg:    s.packetReg,
			}

			if err := handler(rw); err != nil {
				s.logger.Error(err)

				if err := conn.Close(); err != nil {
					s.logger.Error(err)
				}

				return err
			}

			return nil
		}
	}

}

func (s *Server) RegisterAPIs(apis ...interface{}) {

}

func (s *Server) RegisterServices(services ...romulus.Lifecycle) {

}
