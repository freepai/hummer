package shorturl

import (
	"github.com/freepai/hummer/core/plugin"
	"github.com/freepai/hummer/core/shorturl/service"
)

const (
	PluginName = "shorturl"
	ManagerName = "shorturl.manager"
)

type Config struct {
	IDGen    *plugin.Config `yaml:"idGen,omitempty"`
	IDEncode *plugin.Config `yaml:"idEncode,omitempty"`
	IDStore  *plugin.Config `yaml:"idStore,omitempty"`
}

func GetManager(ctx *plugin.Context) *service.Manager {
	return ctx.GetBean(ManagerName).(*service.Manager)
}

