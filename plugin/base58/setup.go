package snowflake

import (
	"github.com/coredns/caddy"
	"github.com/freepai/hummer/server"
)

func init() {
	caddy.RegisterPlugin("base58", caddy.Plugin{
		ServerType: "hummer",
		Action:     setup,
	})

	caddy.RegisterParsingCallback("hummer", "idGen", )
}

func setup(c *caddy.Controller) error {
	ctx := c.Context().(*server.HummerContext)
	idGen := &Base58{}
	ctx.Hummer().RegisterIdEncode(idGen)
	return nil
}
