package config

import (
	"github.com/freepai/hummer/core/plugin"
	"github.com/freepai/hummer/core/server"
	"github.com/freepai/hummer/core/shorturl"
)

type HummerConfig struct {
	Server   *server.Config   `yaml:"server,omitempty"`
	ShortUrl *shorturl.Config `yaml:"shorturl,omitempty"`
	Plugins  []*plugin.Config `yaml:"plugins,omitempty"`
}
