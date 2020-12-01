package server

import "net/http"

type Route struct {
	path string
	handler func(http.ResponseWriter, *http.Request)
}

func NewRoute(path string, handler func(http.ResponseWriter, *http.Request)) *Route {
	return &Route{
		path: path,
		handler: handler,
	}
}