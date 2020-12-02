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
	cfg, _ := config.LoadFromYamlFile(path)
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
	h.ApplyPlugin(server.PluginName, h.config.Server)
	h.ApplyPlugin(shorturl.PluginName, h.config.ShortUrl)

	// others plugins
	plugins := h.config.Plugins
	for _, cfg := range plugins {
		h.ApplyPluginByConfig(cfg)
	}
}

func (h *Hummer) ApplyPlugin(name string, params interface{}) error {
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
