package snowflake

import (
	"github.com/freepai/hummer/core/plugin"
	"github.com/freepai/hummer/core/shorturl"
)

func setup(ctx *plugin.Context) error {
	mgr := shorturl.GetManager(ctx)

	idGen := &Snowflake{}
	mgr.RegisterIdGen(idGen)
	return nil
}

func init() {
	plugin.Register( "snowflake", setup)
}