package core

import (
	"errors"
	"github.com/freepai/hummer/config"
	"github.com/freepai/hummer/core/domain"
	"github.com/freepai/hummer/core/plugin"
	"github.com/freepai/hummer/core/plugin/api"
)

var (
	pluginRegistry map[string]*plugin.Info
)

type Hummer struct {
	idGen api.IdGen
	idEncode api.IDEncode
	idStore api.IDStore
}

func RegisterPlugin(plug *plugin.Info) error {
	pluginRegistry[plug.Name] = plug
	return nil
}

func NewHummer() *Hummer {
	return &Hummer{}
}

func (h *Hummer) RegisterIdGen(idGen api.IdGen) error {
	h.idGen = idGen
	return nil
}

func (h *Hummer) RegisterIdEncode(idEncode api.IDEncode) error {
	h.idEncode = idEncode
	return nil
}

func (h *Hummer) RegisterIdStore(IdStore api.IDStore) error {
	h.idStore = IdStore
	return nil
}

func (h *Hummer) setupPlugins(cfgs []*config.PluginConfig) error {
	for i:=0; i<len(cfgs); i++ {
		pluginCfg := cfgs[i]
		info := pluginRegistry[pluginCfg.Name]

		if info != nil {
			plugin := info.New()
			plugin.Setup(h, pluginCfg.Params)
		} else {
			return errors.New("not found plugin with name:" + pluginCfg.Name)
		}
	}

	return nil
}

func (h *Hummer) InitPlugins(cfg *config.PluginsConfig) {
	h.setupPlugins(cfg.IDGens)
	h.setupPlugins(cfg.IDEncodes)
	h.setupPlugins(cfg.IDStores)
}

func (h *Hummer) Post(ns string, longUrl string) (*domain.ShortUrl, error) {
	id, err := h.idGen.NextUniqueId(ns)
	if err!=nil {
		return nil, err
	}

	code, err := h.idEncode.EncodeId(ns, id)
	if err!=nil {
		return nil, err
	}

	su, err := h.idStore.Save(ns, code, longUrl)
	if err!=nil {
		return nil, err
	}

	return su, nil
}

func (h *Hummer) Get(ns string, code string) (*domain.ShortUrl, error) {
	su, err := h.idStore.Get(ns, code)
	if err!=nil {
		return nil, err
	}

	return su, nil
}