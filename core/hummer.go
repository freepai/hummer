package core

import (
	"github.com/freepai/hummer/core/config"
	"github.com/freepai/hummer/core/plugin"
	"github.com/freepai/hummer/core/server"
	"github.com/freepai/hummer/core/shorturl"
	"log"
)

type Hummer struct {
	config *config.HummerConfig
	beans  map[string]interface{}
}

func NewHummer(path string) *Hummer {
	cfg, _ := config.LoadFromFile(path)
	beans := make(map[string]interface{}, 0)

	return &Hummer{
		config: cfg,
		beans:  beans,
	}
}

func (r *Hummer) RegisterBean(name string, bean interface{}) error {
	r.beans[name] = bean
	return nil
}

func (r *Hummer) GetBean(name string) interface{} {
	return r.beans[name]
}

func (h *Hummer) ApplyPlugins() {
	// shorturl and server plugin
	h.ApplyPlugin(server.PluginName, nil)
	h.ApplyPlugin(shorturl.PluginName, nil)

	// server protocol plugin
	s := h.config.Server
	h.ApplyPlugin(s.Protocol, nil)

	// shorturl extpoint plugin
	su := h.config.ShortUrl
	h.ApplyPluginByConfig(su.IDGen)
	h.ApplyPluginByConfig(su.IDEncode)
	h.ApplyPluginByConfig(su.IDStore)

	// others plugins
	plugins := h.config.Plugins
	for _, cfg := range plugins {
		h.ApplyPluginByConfig(cfg)
	}
}

func (h *Hummer) ApplyPlugin(name string, params map[string]string) error {
	setup := plugin.Get(name)

	if setup != nil {
		ctx := plugin.NewContext(h, params)
		setup(ctx)
	} else {
		log.Fatal("not found plugin with name: " + name)
	}

	return nil
}

func (h *Hummer) ApplyPluginByConfig(cfg *plugin.Config) error {
	return h.ApplyPlugin(cfg.Name, cfg.Params)
}

func (h *Hummer) Start() {
	h.ApplyPlugins()

	mgr := server.GetManagerFromRegistry(h)
	mgr.ListenAndServe(h.config.Server.Addr)
}
