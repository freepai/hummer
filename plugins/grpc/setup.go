package grpc

import (
	"github.com/freepai/hummer/core/plugin"
	"github.com/freepai/hummer/core/server"
)

func setup(ctx *plugin.Context) error {
	manager := server.GetManager(ctx)

	svr := NewGRPCServer()
	manager.SetServer(svr)

	return nil
}

func init() {
	plugin.Register("grpc", setup)
}
