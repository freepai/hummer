package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLoadFromString_ServerConfig(t *testing.T) {
	data := `
server:
    protocol: http
    host: 0.0.0.0
    port: 8081
`

	cfg,err := LoadFromString(data)

	assert.Nil(t, err, "load cfg error should  be nil")
	assert.NotNil(t, cfg, "load cfg should not be nil")

	assert.Equal(t, cfg.Server.Protocol, "http", "server.protocol should be http")
	assert.Equal(t, cfg.Server.Host, "0.0.0.0", "server.host should be 0.0.0.0")
	assert.Equal(t, cfg.Server.Port, 8081, "server.port should be 8081")
}

func TestLoadFromString_PluginsConfig(t *testing.T) {
	data := `
plugins:
    idGen:
        name: snowflake
        params:
            abc: abc
`

	cfg,err := LoadFromString(data)

	assert.Nil(t, err, "load cfg error should  be nil")
	assert.NotNil(t, cfg, "load cfg should not be nil")

	assert.NotNil(t, cfg.Plugins, "Plugins should not be nil")
	assert.NotNil(t, cfg.Plugins.IDGen, "IDStore should not be nil")
	assert.Equal(t, cfg.Plugins.IDGen.Name, "snowflake", "IDGen.Name should be snowflake")
	assert.Equal(t, cfg.Plugins.IDGen.Params["abc"], "abc", "IDGen.Params[\"abc\"] should be snowflake")
}

func TestLoadFromFile(t *testing.T) {
	cfg,err := LoadFromFile("./hummer_test.yaml")

	assert.Nil(t, err, "load cfg error should  be nil")
	assert.NotNil(t, cfg, "load cfg should not be nil")

	assert.Equal(t, cfg.Server.Protocol, "http", "server.protocol should be http")
	assert.Equal(t, cfg.Server.Host, "0.0.0.0", "server.host should be 0.0.0.0")
	assert.Equal(t, cfg.Server.Port, 8081, "server.port should be 8081")

	assert.NotNil(t, cfg.Plugins, "Plugins should not be nil")
	assert.NotNil(t, cfg.Plugins.IDGen, "IDStore should not be nil")
	assert.Equal(t, cfg.Plugins.IDGen.Name, "snowflake", "IDGen.Name should be snowflake")
	assert.Equal(t, cfg.Plugins.IDGen.Params["abc"], "abc", "IDGen.Params[\"abc\"] should be snowflake")
}