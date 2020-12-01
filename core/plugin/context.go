package plugin

type Context struct {
	registry Registry
	params map[string]string
}

func NewContext(registry Registry, params map[string]string) *Context {
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

func (this *Context) GetParam(key string) string {
	return this.params[key]
}