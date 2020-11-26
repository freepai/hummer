package plugin

type Info struct {
	Name string
	New func() Plugin
}

type Plugin interface {
	Info() *Info
	Setup(ctx Context, params map[string]string) error
	TearDown(ctx Context) error
}