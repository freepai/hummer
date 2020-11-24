package core

import (
	"github.com/freepai/hummer/config"
	"github.com/freepai/hummer/core/domain"
	"github.com/freepai/hummer/core/plugin"
	"github.com/freepai/hummer/core/plugin/api"
)

var (
	plugins map[string]plugin.Plugin
)

type Hummer struct {
	idGens map[string]api.IdGen
	idEncodes map[string]api.IDEncode
	idStores map[string]api.IDStore
}

func Register(plug plugin.Plugin) error {
	plugins[plug.Name()] = plug
	return nil
}

func NewHummer() *Hummer {
	return &Hummer{}
}

func (hummer *Hummer) RegisterIdGen(name string, idGen api.IdGen) error {
	hummer.idGens[name] = idGen
	return nil
}

func (hummer *Hummer) RegisterIdEncode(name string, idEncode api.IDEncode) error {
	hummer.idEncodes[name] = idEncode
	return nil
}

func (hummer *Hummer) RegisterIdStore(name string, IdStore api.IDStore) error {
	hummer.idStores[name] = IdStore
	return nil
}

func (hummer *Hummer) setupPlugins(cfgs []*config.PluginConfig) {
	for i:=0; i<len(cfgs); i++ {
		pluginCfg := cfgs[i]
		plugin := plugins[pluginCfg.Name]
		plugin.Setup(hummer)
	}
}

func (hummer *Hummer) InitPlugins(cfg *config.PluginsConfig) {
	hummer.setupPlugins(cfg.IDGens)
	hummer.setupPlugins(cfg.IDEncodes)
	hummer.setupPlugins(cfg.IDStores)
}

func (hummer *Hummer) Post(ns string, longUrl string) (*domain.ShortUrl, error) {
	return nil, nil
}

func (hummer *Hummer) Get(ns string, code string) (*domain.ShortUrl, error) {
	return nil, nil
}