package server

import "github.com/freepai/hummer/core/api"

type GRPCServer struct {
	host string
	port int
	api  api.HummerAPI
}

func NewGRPCServer(host string, port int, api api.HummerAPI) Server {
	return &GRPCServer{
		host: host,
		port: port,
		api: api,
	}
}

func (server *GRPCServer) Start() error {
	return nil
}
