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

func (r *Route) GetPath() string {
	return r.path
}

func (r *Route) GetHandler() func(http.ResponseWriter, *http.Request) {
	return r.handler
}