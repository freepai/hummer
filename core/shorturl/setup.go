package shorturl

import (
	"github.com/freepai/hummer/core/plugin"
	"github.com/freepai/hummer/core/server"
	"github.com/freepai/hummer/core/shorturl/controller"
	"github.com/freepai/hummer/core/shorturl/service"
)

func setup(ctx *plugin.Context) error {
	shorturl := &service.Manager{}
	ctx.RegisterBean(ShortUrlManager, shorturl)

	mgr := server.GetManager(ctx)

	c := controller.NewShortUrlController(shorturl)
	mgr.HandleFunc("/api/1/shorturls", c.PostShortUrl)

	return nil
}

func init() {
	plugin.Register(PluginName, setup)
}
