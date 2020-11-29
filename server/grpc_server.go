package server

import (
	"fmt"
	"net"
)

type GRPCServer struct {
	*Server
}

func NewGRPCServer(host string, port int) *GRPCServer {
	s := NewServer(fmt.Sprintf("%s:%d", host, port))

	return &GRPCServer{
		Server: s,
	}
}

// Listen implements caddy.TCPServer interface.
func (s *GRPCServer) Listen() (net.Listener, error) {
	return nil, nil
}

// Serve implements caddy.TCPServer interface.
func (s *GRPCServer) Serve(l net.Listener) error {
	return nil
}

// ListenPacket implements caddy.UDPServer interface.
func (s *GRPCServer) ListenPacket() (net.PacketConn, error) { return nil, nil }
// ServePacket implements caddy.UDPServer interface.
func (s *GRPCServer) ServePacket(p net.PacketConn) error { return nil }

