package snowflake

import (
	"github.com/freepai/hummer/core/plugin"
)

var (
	pluginInfo = &plugin.Info{}
)

type SnowflakePlugin struct {

}

func (p *SnowflakePlugin) Info() *plugin.Info {
	return pluginInfo
}

func (p *SnowflakePlugin) Setup(ctx plugin.Context, parmas map[string]string) error {
	idGen := &Snowflake{}
	ctx.RegisterIdGen(idGen)
	return nil
}

func (p *SnowflakePlugin) TearDown(ctx plugin.Context) error {
	return nil
}

func init() {
	pluginInfo.Name = "snowflake"
	pluginInfo.New = func() plugin.Plugin {
		return &SnowflakePlugin{}
	}

	plugin.Register(pluginInfo)
}