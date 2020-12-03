package plugin

import "errors"

type Config struct {
	Name   string            `yaml:"name,omitempty"`
	Params map[string]string `yaml:"params,omitempty"`
}

type Container interface {
	AddBean(name string, bean interface{}) error
	GetBean(name string) interface{}
	ApplyPlugin(name string, config interface{}) error
}

type Context struct {
	container Container
	config    interface{}
}

type Plugin func(ctx *Context) error

func NewContext(container Container, config interface{}) *Context {
	return &Context{
		container: container,
		config:    config,
	}
}

func (c *Context) Register(name string, bean interface{}) error {
	return c.AddBean(name, bean)
}

func (c *Context) AddBean(name string, bean interface{}) error {
	c.container.AddBean(name, bean)
	return nil
}

func (c *Context) GetBean(name string) interface{} {
	return c.container.GetBean(name)
}

func (c *Context) GetConfig() interface{} {
	return c.config
}

func (c *Context) GetParam(key string) (string, error) {
	params, ok := c.config.(map[string]string)
	if !ok {
		return "", errors.New("not support")
	}

	return params[key], nil
}

func (c *Context) ApplyPlugin(name string, params map[string]string) error {
	return c.container.ApplyPlugin(name, params)
}
