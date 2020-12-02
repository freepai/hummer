package shorturl

import (
	"github.com/freepai/hummer/core/plugin"
	"github.com/freepai/hummer/core/shorturl/service"
)

func GetManager(ctx *plugin.Context) *service.Manager {
	return ctx.GetBean(ShortUrlManager).(*service.Manager)
}
