package cmd

import (
	"context"
	"errors"
	"grpc-boiler-plate-go/env"
	"grpc-boiler-plate-go/infra/network"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type CLI interface {
	Start()
	Error() error
}

type Command struct {
	di   *env.Dependency
	args []string
	err  error
}

func NewCLI(params *env.Dependency, args []string) CLI {
	return &Command{params, args, nil}
}

func (cmd *Command) Start() {
	cmd.gRPCService()
}
func (cmd *Command) Error() error {
	return cmd.err
}

func (cmd *Command) gRPCService() error {
	routeHandler := network.InitGRPCServer(cmd.di)
	exit := make(chan os.Signal, 1)
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
		signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
		<-exit
		routeHandler.GracefulStop()
	}()
	listen, err := net.Listen("tcp", cmd.di.Params.Ports.GRPC)
	if err != nil {
		return err
	}
	go func() {
		if err := routeHandler.Serve(listen); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
			return
		}
	}()

	return nil
}
