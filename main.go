package main

import (
	"github.com/freepai/hummer/config"
	"github.com/freepai/hummer/core"
	"github.com/freepai/hummer/server"
)

func main() {
	cfg,_ := config.LoadFromFile("hummer.yaml")

	hummer := core.NewHummer()
	hummer.InitPlugins(cfg.Plugins)

	server := server.NewServer(cfg.Server, hummer)
	server.Start()
}
