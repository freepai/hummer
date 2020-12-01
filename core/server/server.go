package server

type Server interface {
	ListenAndServe(addr string) error
}
