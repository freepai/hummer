package plugin

type Registry interface {
	RegisterBean(name string, bean interface{}) error
	GetBean(name string) interface{}
}