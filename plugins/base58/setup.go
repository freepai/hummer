package snowflake

import (
	"github.com/freepai/hummer/core/plugin"
	"github.com/freepai/hummer/core/shorturl"
)

func setup(ctx *plugin.Context) error {
	manager := shorturl.GetManager(ctx)

	idEncode := &Base58{}
	manager.RegisterIdEncode(idEncode)

	return nil
}

func init() {
	plugin.Register("base58", setup)
}
