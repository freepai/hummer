package shorturl

import "github.com/freepai/hummer/core/plugin"

type Config struct {
	IDGen    *plugin.Config `yaml:"idGen,omitempty"`
	IDEncode *plugin.Config `yaml:"idEncode,omitempty"`
	IDStore  *plugin.Config `yaml:"idStore,omitempty"`
}
