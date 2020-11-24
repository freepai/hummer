package plugin

type Plugin interface {
	Name() string
	Setup(ctx Context) error
	TearDown(ctx Context) error
}