package snowflake

import (
	"github.com/coredns/caddy"
	"github.com/freepai/hummer/server"
)

func init() {
	server.RegisterPlugin("snowflake", setup)
}

func setup(c *caddy.Controller) error {
	ctx := c.Context().(*server.HummerContext)
	idGen := &Snowflake{}
	ctx.Hummer().RegisterIdGen(idGen)
	return nil
}