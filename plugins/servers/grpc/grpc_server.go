package grpc

import "github.com/freepai/hummer/core/server"

type GRPCServer struct {
	host string
	port int
}

func NewGRPCServer(host string, port int) server.Server {
	return &GRPCServer{
		host: host,
		port: port,
	}
}

func (server *GRPCServer) Start() error {
	return nil
}
