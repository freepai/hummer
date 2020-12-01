package shorturl

import (
	"github.com/freepai/hummer/core/plugin"
	"github.com/freepai/hummer/core/server"
	"net/http"
)

func setup(ctx *plugin.Context) error {
	shorturl := &Manager{}
	ctx.RegisterBean(ShortUrlManager, shorturl)

	mgr := server.GetManager(ctx)
	mgr.AddRouteFunc("/api/1/shorturls", func(http.ResponseWriter, *http.Request) {

	})

	return nil
}

func init() {
	plugin.Register(PluginName, setup)
}
