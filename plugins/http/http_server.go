package http

import (
	"github.com/freepai/hummer/core/server"
	"log"
	"net/http"
)

type HTTPServer struct {
	manager *server.Manager
}

func NewHTTPServer(mgr *server.Manager) *HTTPServer {
	return &HTTPServer{
		manager: mgr,
	}
}

func (server *HTTPServer) Config() error {
	routes := server.manager.AllRoutes()

	for _, route := range routes {
		http.HandleFunc(route.Path, route.Handler)
	}

	return nil
}

func (server *HTTPServer) ListenAndServe(addr string) error {
	log.Printf("ListenAndServe: %s", addr)
	http.ListenAndServe(addr, nil)
	return nil
}
