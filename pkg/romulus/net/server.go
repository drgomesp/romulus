package net

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"time"

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
			{
				for _, proto := range s.protocols {
					conn, err := s.listener.Accept()
					if err != nil {
						log.Fatal(err)
					}

					if err = s.handler(proto.ID, proto.Run)(conn); err != nil {
						log.Fatal(err)
					}
				}
			}
		case <-time.After(0):
			{

			}
		}
	}

	panic("WTF")
}

func (s *Server) RegisterProtocols(protocols ...*Protocol) {
	s.protocols = protocols

	s.handler = func(pid ProtocolID, handler ProtocolHandlerFunc) ConnHandlerFunc {
		return func(conn net.Conn) error {
			rw := &protoRW{
				logger: s.logger,
				conn:   conn,
				reader: bufio.NewReader(conn),
				writer: bufio.NewWriter(conn),
				reg:    s.packetReg,
			}

			err := handler(rw)

			if err != nil {
				s.logger.Debugw("RegisterProtocols")

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
