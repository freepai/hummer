package config

type ServerConfig struct {
	Protocol string
	Host string
	Port int
}

type PluginConfig struct {
	Name string
	Params map[string]string
}

type PluginsConfig struct {
	IDGens []*PluginConfig
	IDEncodes []*PluginConfig
	IDStores []*PluginConfig
}

type HummerConfig struct {
	ServerCfg *ServerConfig
	PluginsCfg *PluginsConfig
}

