package config

type PluginConfig struct {
	Name   string            `yaml:"name,omitempty"`
	Params map[string]string `yaml:"params,omitempty"`
}

type PluginsConfig struct {
	IDGen    *PluginConfig `yaml:"idGen,omitempty"`
	IDEncode *PluginConfig `yaml:"idEncode,omitempty"`
	IDStore  *PluginConfig `yaml:"idStore,omitempty"`
}