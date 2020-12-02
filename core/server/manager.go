package server

import (
	"log"
	"net/http"
)

type Manager struct {
	routes []*Route
	server  Server
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) SetServer(server Server) {
	m.server = server
}

func (m *Manager) GetServer() Server{
	return m.server
}

func (m *Manager) AddRoute(route *Route) {
	m.routes = append(m.routes, route)
}

func (m *Manager) HandleFunc(path string, handler func(http.ResponseWriter, *http.Request)) {
	route := NewRoute(path, handler)
	m.AddRoute(route)
}

func (m *Manager) AllRoutes() []*Route {
	return m.routes
}

func (m *Manager) ListenAndServe(addr string) {
	svr := m.server

	if svr!= nil {
		svr.Config()
		svr.ListenAndServe(addr)
	} else {
		log.Fatal("not config server protocol plugin")
	}
}