package leveldb

import (
	"github.com/freepai/hummer/core/plugin"
	"github.com/freepai/hummer/core/shorturl"
)

func setup(ctx *plugin.Context) error {
	manager := shorturl.GetManager(ctx)

	idStore := &LevelDB{}
	manager.RegisterIdStore(idStore)
	return nil
}

func init() {
	plugin.Register("leveldb", setup)
}
