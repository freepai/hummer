package shorturl

import (
	"errors"
	"github.com/freepai/hummer/core/plugin"
	"github.com/freepai/hummer/core/server"
	"github.com/freepai/hummer/core/shorturl/controller"
	"github.com/freepai/hummer/core/shorturl/service"
)

func setup(ctx *plugin.Context) error {
	// shorturl extpoint plugin
	cfg, ok := ctx.GetConfig().(*Config)
	if !ok {
		return errors.New("not support shorturl config")
	}

	// register shorturl manager
	shorturl := &service.Manager{}
	ctx.Register(ManagerName, shorturl)

	// apply extpoint plugins
	ctx.ApplyPlugin(cfg.IDGen.Name, cfg.IDGen.Params)
	ctx.ApplyPlugin(cfg.IDEncode.Name, cfg.IDEncode.Params)
	ctx.ApplyPlugin(cfg.IDStore.Name, cfg.IDStore.Params)

	// export routes
	mgr := server.GetManager(ctx)
	c := controller.NewShortUrlController(shorturl)
	mgr.HandleFunc("/api/1/shorturls", c.PostShortUrl)

	return nil
}

func init() {
	plugin.Register(PluginName, setup)
}
