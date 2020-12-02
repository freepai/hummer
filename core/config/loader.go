package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func LoadFromString(data string) (*HummerConfig, error) {
	return LoadFromBytes([]byte(data))
}

func LoadFromBytes(data []byte) (*HummerConfig, error) {
	cfg := &HummerConfig{}

	err := yaml.Unmarshal(data, cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return cfg, nil
}

func LoadFromFile(path string) (*HummerConfig, error) {
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

	return LoadFromBytes(data)
}