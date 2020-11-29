package leveldb

import (
	"github.com/coredns/caddy"
	"github.com/freepai/hummer/server"
)

func init() {
	server.RegisterPlugin("leveldb", setup)
}

func setup(c *caddy.Controller) error {
	ctx := c.Context().(*server.HummerContext)
	idStore := &LevelDB{}
	ctx.Hummer().RegisterIdStore(idStore)
	return nil
}