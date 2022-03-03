package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/drgomesp/rhizom/internal/roserver"
	"github.com/drgomesp/rhizom/pkg/romulus/net"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var app *cli.App

func init() {
	app = &cli.App{
		Name: "romulus-server",
		Action: func(c *cli.Context) (err error) {
			ctx, cancelFunc := context.WithCancel(c.Context)
			defer cancelFunc()

			srv, err := makeServer(ctx)
			if err != nil {
				return errors.Wrap(err, "failed to initialize server")
			}

			return startServer(ctx, srv.Server)
		},
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func buildLogger() (*zap.Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.Kitchen)

	logger, err := config.Build()
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize logger")
	}

	return logger, nil
}

func makeServer(_ context.Context) (*roserver.Server, error) {
	logger, err := buildLogger()
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize logger")
	}

	var srv *roserver.Server
	if srv, err = roserver.NewServer(logger.Sugar()); err != nil {
		return nil, errors.Wrap(err, "failed to initialize server")
	}

	return srv, nil
}

func startServer(ctx context.Context, srv *net.Server) error {
	if err := srv.Start(ctx); err != nil {
		return err
	}

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		defer signal.Stop(sig)

		<-sig
		log.Println("interrupt signal, shutting down...")
	}()

	return nil
}
