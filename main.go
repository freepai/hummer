package main

import (
	"github.com/freepai/hummer/config"
	"github.com/freepai/hummer/core"
	"github.com/freepai/hummer/core/server"
)

func main() {
	cfg := config.Load("HummerFile")
	hummer := core.NewHummer()
	hummer.InitPlugins(cfg.PluginsCfg)

	server := server.NewServer(cfg.ServerCfg, hummer)
	server.Start()
}
