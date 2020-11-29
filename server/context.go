package server

import (
	"github.com/coredns/caddy"
	"github.com/coredns/caddy/caddyfile"
	"github.com/freepai/hummer/core"
	"github.com/freepai/hummer/util"
	"strings"
)

type HummerContext struct {
	serverConfigs []*Config
	hummer *core.Hummer
}

// Compile-time check to ensure dnsContext implements the caddy.Context interface
var _ caddy.Context = &HummerContext{}

// InspectServerBlocks make sure that everything checks out before
// executing directives and otherwise prepares the directives to
// be parsed and executed.
func (h *HummerContext) InspectServerBlocks(sourceFile string, serverBlocks []caddyfile.ServerBlock) ([]caddyfile.ServerBlock, error) {
	for _, sb := range serverBlocks {
		keys := sb.Keys

		for _, k := range keys {
			tokens := strings.Split(k, ":")
			cfg := &Config{
				Protocol: tokens[0],
				Host: tokens[1],
				Port: util.Str2Int(tokens[2]),
			}

			h.serverConfigs = append(h.serverConfigs, cfg)
		}
	}

	return serverBlocks, nil
}

// MakeServers uses the newly-created siteConfigs to create and return a list of server instances.
func (h *HummerContext) MakeServers() ([]caddy.Server, error) {
	servers := make([]caddy.Server, 0)

	if h.serverConfigs!=nil {
		for _, servCfg := range h.serverConfigs {
			switch servCfg.Protocol {
			case "http":
				server := NewHTTPServer(servCfg.Host, servCfg.Port)
				servers = append(servers, server)
			case "grpc":
				server := NewGRPCServer(servCfg.Host, servCfg.Port)
				servers = append(servers, server)
			}
		}
	}

	return servers, nil
}

// Get core hummer
func (h *HummerContext) Hummer() *core.Hummer {
	return h.hummer
}
