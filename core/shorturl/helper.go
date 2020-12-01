package shorturl

import "github.com/freepai/hummer/core/plugin"

func GetManager(ctx *plugin.Context) *Manager {
	return ctx.GetBean(ShortUrlManager).(*Manager)
}
