package snowflake

import (
	"github.com/freepai/hummer/core/plugin"
	"github.com/freepai/hummer/core/shorturl"
)

func setup(ctx *plugin.Context) error {
	manager := shorturl.GetManager(ctx)

	idGen := &Snowflake{}
	manager.RegisterIdGen(idGen)
	return nil
}

func init() {
	plugin.Register( "snowflake", setup)
}