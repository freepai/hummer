package plugin

var (
	registry map[string]*Info
)

func Register(plug *Info) error {
	registry[plug.Name] = plug
	return nil
}

func Get(name string) *Info {
	return registry[name]
}