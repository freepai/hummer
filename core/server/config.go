package server

type Config struct {
	Protocol string `yaml:"protocol,omitempty"`
	Addr     string `yaml:"addr,omitempty"`
}
