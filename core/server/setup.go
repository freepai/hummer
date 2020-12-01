package server

import "github.com/freepai/hummer/core/plugin"

func setup(ctx *plugin.Context) error {
	mgr := &Manager{}
	ctx.GetRegistry().RegisterBean(ManagerName, mgr)
	return nil
}

func init() {
	plugin.Register(PluginName, setup)
}

