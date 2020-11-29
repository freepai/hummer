package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

type HTTPServer struct {
	*Server
	httpServer *http.Server
}

func NewHTTPServer(host string, port int) *HTTPServer {
	s := NewServer(fmt.Sprintf("%s:%d", host, port))
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("new HTTPServer")

	return &HTTPServer{
		Server: s,
		httpServer: srv,
	}
}

// Listen implements caddy.TCPServer interface.
func (s *HTTPServer) Listen() (net.Listener, error) {
	l, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return nil, err
	}
	return l, nil
}

// Serve implements caddy.TCPServer interface.
func (s *HTTPServer) Serve(l net.Listener) error {
	return s.httpServer.Serve(l)
}

// ListenPacket implements caddy.UDPServer interface.
func (s *HTTPServer) ListenPacket() (net.PacketConn, error) { return nil, nil }
// ServePacket implements caddy.UDPServer interface.
func (s *HTTPServer) ServePacket(p net.PacketConn) error { return nil }
