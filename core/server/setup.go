package server

import "github.com/freepai/hummer/core/plugin"

func setup(ctx *plugin.Context) error {
	mgr := &Manager{}
	ctx.GetRegistry().RegisterBean(ManagerName, mgr)

	// server protocol plugin
	cfg := ctx.GetParams().(*Config)
	ctx.ApplyPlugin(cfg.Protocol, nil)

	return nil
}

func init() {
	plugin.Register(PluginName, setup)
}
