package http

import (
	"log"
	"net/http"
)

type HTTPServer struct {}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{}
}

func (server *HTTPServer) ListenAndServe(addr string) error {
	log.Printf("ListenAndServe: %s", addr)
	http.ListenAndServe(addr, nil)
	return nil
}