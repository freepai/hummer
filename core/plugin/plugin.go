package plugin

type Config struct {
	Name   string            `yaml:"name,omitempty"`
	Params map[string]string `yaml:"params,omitempty"`
}

type Registry interface {
	RegisterBean(name string, bean interface{}) error
	GetBean(name string) interface{}
	ApplyPlugin(name string, params interface{}) error
}

type Context struct {
	registry Registry
	params interface{}
}

type Plugin func(ctx *Context)error

func NewContext(registry Registry, params interface{}) *Context {
	return &Context{
		registry: registry,
		params: params,
	}
}

func (this *Context) GetRegistry() Registry {
	return this.registry
}

func (this *Context) RegisterBean(name string, bean interface{}) error {
	this.registry.RegisterBean(name, bean)
	return nil
}

func (this *Context) GetBean(name string) interface{} {
	return this.registry.GetBean(name)
}

func (this *Context) GetParams() interface{} {
	return this.params
}

func (this *Context) GetParam(key string) string {
	params := this.params.(map[string]string)

	if params!=nil {
		return params[key]
	}

	return ""
}

func (this *Context) ApplyPlugin(name string, params map[string]string) error {
	return this.registry.ApplyPlugin(name, params)
}