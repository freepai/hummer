package config

type ServerConfig struct {
	Protocol string `yaml:"protocol,omitempty"`
	Host     string `yaml:"host,omitempty"`
	Port     int    `yaml:"port,omitempty"`
}

type PluginConfig struct {
	Name   string            `yaml:"name,omitempty"`
	Params map[string]string `yaml:"params,omitempty"`
}

type PluginsConfig struct {
	IDGen    *PluginConfig `yaml:"idGen,omitempty"`
	IDEncode *PluginConfig `yaml:"idEncode,omitempty"`
	IDStore  *PluginConfig `yaml:"idStore,omitempty"`
}

type HummerConfig struct {
	Server  *ServerConfig  `yaml:"server,omitempty"`
	Plugins *PluginsConfig `yaml:"plugins,omitempty"`
}
