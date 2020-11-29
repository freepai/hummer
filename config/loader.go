package config

import (
	"github.com/coredns/caddy"
	"io/ioutil"
	"os"
)

const (
	DefaultConfigFile ="./Hummerfile"
)

// confLoader loads the Caddyfile using the -conf flag.
func Loader(serverType string) (caddy.Input, error) {
	conf := DefaultConfigFile

	contents, err := ioutil.ReadFile(conf)
	if err != nil {
		return nil, err
	}
	return caddy.CaddyfileInput{
		Contents:       contents,
		Filepath:       conf,
		ServerTypeName: serverType,
	}, nil
}

// defaultLoader loads the Corefile from the current working directory.
func DefaultLoader(serverType string) (caddy.Input, error) {
	contents, err := ioutil.ReadFile(caddy.DefaultConfigFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	return caddy.CaddyfileInput{
		Contents:       contents,
		Filepath:       caddy.DefaultConfigFile,
		ServerTypeName: serverType,
	}, nil
}