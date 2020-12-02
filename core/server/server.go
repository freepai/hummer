package server

import (
	"github.com/freepai/hummer/core/plugin"
	"log"
	"net/http"
)

const (
	PluginName  = "server"
	ManagerName = "server.manager"
)

type Config struct {
	Protocol string `yaml:"protocol,omitempty"`
	Addr     string `yaml:"addr,omitempty"`
}

type Route struct {
	Path string
	Handler func(http.ResponseWriter, *http.Request)
}

type Server interface {
	Config() error
	ListenAndServe(addr string) error
}

type Manager struct {
	routes []*Route
	server Server
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) SetServer(server Server) {
	m.server = server
}

func (m *Manager) GetServer() Server {
	return m.server
}

func (m *Manager) AddRoute(route *Route) {
	m.routes = append(m.routes, route)
}

func (m *Manager) HandleFunc(path string, handler func(http.ResponseWriter, *http.Request)) {
	route := &Route{
		Path:    path,
		Handler: handler,
	}

	m.AddRoute(route)
}

func (m *Manager) AllRoutes() []*Route {
	return m.routes
}

func (m *Manager) ListenAndServe(addr string) {
	svr := m.server

	if svr != nil {
		svr.Config()
		svr.ListenAndServe(addr)
	} else {
		log.Fatal("not config server protocol plugin")
	}
}

func GetManager(ctx *plugin.Context) *Manager {
	return ctx.GetBean(ManagerName).(*Manager)
}

func GetManagerFromRegistry(registry plugin.Registry) *Manager {
	return registry.GetBean(ManagerName).(*Manager)
}
