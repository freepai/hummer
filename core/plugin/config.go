package plugin

type Config struct {
	Name   string            `yaml:"name,omitempty"`
	Params map[string]string `yaml:"params,omitempty"`
}
