package server

import (
	"errors"
	"github.com/freepai/hummer/core/plugin"
)

func setup(ctx *plugin.Context) error {
	cfg, ok := ctx.GetConfig().(*Config)
	if !ok {
		return errors.New("not support config")
	}

	// register ServerManager
	mgr := NewManager(cfg.Addr)
	ctx.Register(ManagerName, mgr)

	// apply protocol plugin
	ctx.ApplyPlugin(cfg.Protocol, nil)

	return nil
}

func init() {
	plugin.Register(PluginName, setup)
}
