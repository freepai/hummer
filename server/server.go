package server

type Server struct {
	Addr string // Address we listen on
}

func NewServer(addr string) *Server{
	return &Server{
		Addr: addr,
	}
}