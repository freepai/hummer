package server

import (
	"github.com/coredns/caddy"
	"github.com/freepai/hummer/core"
)

const serverType = "hummer"

var Directives = []string{
	"idGen",
	"idEncode",
	"idStore",
}

// Any flags defined here, need to be namespaced to the serverType other
// wise they potentially clash with other server types.
func init() {
	caddy.RegisterServerType(serverType, caddy.ServerType{
		Directives: func() []string { return Directives },
		NewContext: newContext,
	})
}

func newContext(i *caddy.Instance) caddy.Context {
	hummer := core.NewHummer()
	return &HummerContext{hummer: hummer}
}

// Register registers your plugin with CoreDNS and allows it to be called when the server is running.
func RegisterPlugin(name string, action caddy.SetupFunc) {
	caddy.RegisterPlugin(name, caddy.Plugin{
		ServerType: "hummer",
		Action:     action,
	})
}

