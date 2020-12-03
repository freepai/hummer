package config

import (
	"github.com/freepai/hummer/core/plugin"
	"github.com/freepai/hummer/core/server"
	"github.com/freepai/hummer/core/shorturl"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type HummerConfig struct {
	Server   *server.Config   `yaml:"server,omitempty"`
	ShortUrl *shorturl.Config `yaml:"shorturl,omitempty"`
	Plugins  []*plugin.Config `yaml:"plugins,omitempty"`
}

func LoadFromYamlString(data string) (*HummerConfig, error) {
	cfg := &HummerConfig{}

	err := yaml.Unmarshal([]byte(data), cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return cfg, nil
}

func LoadFromYamlFile(path string) (*HummerConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return LoadFromYamlString(string(data))
}
