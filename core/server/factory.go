package server

import (
	"github.com/freepai/hummer/config"
	"github.com/freepai/hummer/core/api"
)

func NewServer(cfg *config.ServerConfig, api api.HummerAPI) Server {
	protocol := cfg.Protocol

	if protocol=="http" {
		return NewHTTPServer(cfg.Host, cfg.Port, api)
	} else if protocol=="grpc" {
		return NewGRPCServer(cfg.Host, cfg.Port, api)
	}

	return nil
}