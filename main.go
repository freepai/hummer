package main

import (
	"github.com/coredns/caddy"
	"github.com/freepai/hummer/config"
	_ "github.com/freepai/hummer/server"
	_ "github.com/freepai/hummer/plugin"
	"log"
)

// Various CoreDNS constants.
const (
	hummerVersion = "0.1.0"
	hummerName    = "Hummer"
	serverType    = "hummer"
)

func init() {
	caddy.DefaultConfigFile = "Hummerfile"
	caddy.Quiet = true // don't show init stuff from caddy

	caddy.RegisterCaddyfileLoader("flag", caddy.LoaderFunc(config.Loader))
	caddy.SetDefaultCaddyfileLoader("default", caddy.LoaderFunc(config.DefaultLoader))

	caddy.AppName = hummerName
	caddy.AppVersion = hummerVersion
}

func main() {
	// Get Hummerfile input
	hummerfile, err := caddy.LoadCaddyfile(serverType)
	if err != nil {
		log.Fatal(err)
	}

	// Start your engines
	instance, err := caddy.Start(hummerfile)
	if err != nil {
		log.Fatal(err)
	}

	// Twiddle your thumbs
	instance.Wait()
}
