package server

type Server interface {
	Config() error
	ListenAndServe(addr string) error
}
