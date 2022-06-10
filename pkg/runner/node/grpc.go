package node

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type GrpcConfig struct {
	Host string
	Port uint
}

type Grpc struct {
	*grpc.Server
	Config GrpcConfig
}

func NewGrpc(config GrpcConfig) *Grpc {
	return &Grpc{
		Server: grpc.NewServer(),
		Config: config,
	}
}

func (g *Grpc) Name() string {
	return "server"
}

func (g *Grpc) Run() error {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", g.Config.Host, g.Config.Port))

	if err != nil {
		return err
	}

	err = g.Server.Serve(listen)

	if err != nil {
		return err
	}

	return nil
}

func (g *Grpc) Close() error {
	g.Server.GracefulStop()

	return nil
}
