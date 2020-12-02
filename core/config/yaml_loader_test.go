package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadFromString_ServerConfig(t *testing.T) {
	data := `
server:
    protocol: http
    addr: 0.0.0.0:8081
`

	cfg,err := LoadFromYamlString(data)

	assert.Nil(t, err, "load cfg error should  be nil")
	assert.NotNil(t, cfg, "load cfg should not be nil")

	assert.Equal(t, cfg.Server.Protocol, "http", "server.protocol should be http")
	assert.Equal(t, cfg.Server.Addr, "0.0.0.0:8081", "server.host should be 0.0.0.0")
}

func TestLoadFromString_PluginsConfig(t *testing.T) {
	data := `
shorturl:
    idGen:
        name: snowflake
        params:
            abc: abc
`

	cfg,err := LoadFromYamlString(data)

	assert.Nil(t, err, "load cfg error should  be nil")
	assert.NotNil(t, cfg, "load cfg should not be nil")

	assert.NotNil(t, cfg.ShortUrl, "Plugins should not be nil")
	assert.NotNil(t, cfg.ShortUrl.IDGen, "IDStore should not be nil")
	assert.Equal(t, cfg.ShortUrl.IDGen.Name, "snowflake", "IDGen.Name should be snowflake")
	assert.Equal(t, cfg.ShortUrl.IDGen.Params["abc"], "abc", "IDGen.Params[\"abc\"] should be snowflake")
}

func TestLoadFromFile(t *testing.T) {
	cfg,err := LoadFromYamlFile("./hummer_test.yaml")

	assert.Nil(t, err, "load cfg error should  be nil")
	assert.NotNil(t, cfg, "load cfg should not be nil")

	assert.Equal(t, cfg.Server.Protocol, "http", "server.protocol should be http")
	assert.Equal(t, cfg.Server.Addr, "0.0.0.0:8081", "server.host should be 0.0.0.0")

	assert.NotNil(t, cfg.ShortUrl, "Plugins should not be nil")
	assert.NotNil(t, cfg.ShortUrl.IDGen, "IDStore should not be nil")
	assert.Equal(t, cfg.ShortUrl.IDGen.Name, "snowflake", "IDGen.Name should be snowflake")
	assert.Equal(t, cfg.ShortUrl.IDGen.Params["abc"], "abc", "IDGen.Params[\"abc\"] should be snowflake")
}