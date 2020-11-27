package core

import (
	"errors"
	"github.com/freepai/hummer/config"
	"github.com/freepai/hummer/core/domain"
	"github.com/freepai/hummer/core/plugin"
	"github.com/freepai/hummer/core/plugin/api"
)

type Hummer struct {
	idGen api.IdGen
	idEncode api.IDEncode
	idStore api.IDStore
	hooks map[*api.Hook]api.Hook
}

func NewHummer() *Hummer {
	return &Hummer{}
}

func (h *Hummer) RegisterIdGen(idGen api.IdGen) func() {
	h.idGen = idGen

	return func(){
		h.idGen = nil
	}
}

func (h *Hummer) RegisterIdEncode(idEncode api.IDEncode) func() {
	h.idEncode = idEncode

	return func() {
		h.idEncode = nil
	}
}

func (h *Hummer) RegisterIdStore(IdStore api.IDStore) func() {
	h.idStore = IdStore

	return func(){
		h.idStore = nil
	}
}

func (h *Hummer) RegisterHook(hook api.Hook) func() {
	h.hooks[&hook] = hook

	return func() {
		delete(h.hooks, &hook)
	}
}

func (h *Hummer) InitPlugins(cfg *config.PluginsConfig) {
	h.ApplyPlugin(cfg.IDGen)
	h.ApplyPlugin(cfg.IDEncode)
	h.ApplyPlugin(cfg.IDStore)
}

func (h *Hummer) ApplyPlugin(cfg *config.PluginConfig) error {
	info := plugin.Get(cfg.Name)

	if info != nil {
		plugin := info.New()
		plugin.Setup(h, cfg.Params)
	} else {
		return errors.New("not found plugin with name:" + cfg.Name)
	}

	return nil
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