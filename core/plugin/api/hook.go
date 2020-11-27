package api

type Hook interface {
	BeforeIDGen() error
	AfterIDGen() error

	BeforeIDEncode() error
	AfterIDEncode() error

	BeforeIDStore() error
	AfterIDStore() error
}
