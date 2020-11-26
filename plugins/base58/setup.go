package snowflake

import (
	"github.com/freepai/hummer/core"
	"github.com/freepai/hummer/core/plugin"
)

var (
	pluginInfo = &plugin.Info{}
)

type Base58Plugin struct {

}

func (p *Base58Plugin) Info() *plugin.Info {
	return pluginInfo
}

func (p *Base58Plugin) Setup(ctx plugin.Context, params map[string]string) error {
	idGen := &Base58{}
	ctx.RegisterIdEncode("base58", idGen)
	return nil
}

func (p *Base58Plugin) TearDown(ctx plugin.Context) error {
	return nil
}

func init() {
	pluginInfo.Name = "base58"
	pluginInfo.New = func() plugin.Plugin {
		return &Base58Plugin{}
	}

	core.RegisterPlugin(pluginInfo)
}