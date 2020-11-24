package snowflake

import (
	"github.com/freepai/hummer/core"
	"github.com/freepai/hummer/core/plugin"
)

type SnowflakePlugin struct {

}

func (this *SnowflakePlugin) Name() string {
	return "snowflake"
}

func (this *SnowflakePlugin) Setup(ctx plugin.Context) error {
	idGen := &Snowflake{}
	ctx.RegisterIdGen("snowflake", idGen)
	return nil
}

func (this *SnowflakePlugin) TearDown(ctx plugin.Context) error {
	return nil
}

func init() {
	plugin := &SnowflakePlugin{}
	core.Register(plugin)
}