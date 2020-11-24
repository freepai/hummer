package server

import "github.com/freepai/hummer/core/api"

type HTTPServer struct {
	host string
	port int
	api  api.HummerAPI
}

func NewHTTPServer(host string, port int, api api.HummerAPI) Server {
	return &HTTPServer{
		host: host,
		port: port,
		api: api,
	}
}

func (server *HTTPServer) Start() error {
	return nil
}