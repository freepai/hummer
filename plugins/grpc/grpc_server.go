package grpc

type GRPCServer struct{}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

func (server *GRPCServer) Config() error {
	return nil
}

func (server *GRPCServer) ListenAndServe(addr string) error {
	return nil
}
