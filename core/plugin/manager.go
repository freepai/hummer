package plugin

var (
	registry map[string]Plugin
)

func init() {
	registry = make(map[string]Plugin, 0)
}

func Register(name string, setup Plugin) error {
	registry[name] = setup
	return nil
}

func Get(name string) Plugin {
	return registry[name]
}

