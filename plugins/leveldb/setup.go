package leveldb

import (
	"github.com/freepai/hummer/core/plugin"
)

var (
	pluginInfo = &plugin.Info{}
)

type LevelDBPlugin struct {

}

func (p *LevelDBPlugin) Info() *plugin.Info {
	return pluginInfo
}

func (p *LevelDBPlugin) Setup(ctx plugin.Context, params map[string]string) error {
	idStore := &LevelDB{}
	ctx.RegisterIdStore(idStore)
	return nil
}

func (p *LevelDBPlugin) TearDown(ctx plugin.Context) error {
	return nil
}

func init() {
	pluginInfo.Name = "leveldb"
	pluginInfo.New = func() plugin.Plugin {
		return &LevelDBPlugin{}
	}

	plugin.Register(pluginInfo)
}
