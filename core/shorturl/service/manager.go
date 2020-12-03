package service

import (
	"github.com/freepai/hummer/core/shorturl/domain"
	"github.com/freepai/hummer/core/shorturl/extpoint"
)

type Manager struct {
	idGen    extpoint.IdGen
	idEncode extpoint.IDEncode
	idStore  extpoint.IDStore
	hooks    map[*extpoint.Hook]extpoint.Hook
}

func NewManager() *Manager {
	return &Manager{}
}

func (h *Manager) RegisterIdGen(idGen extpoint.IdGen) func() {
	h.idGen = idGen

	return func(){
		h.idGen = nil
	}
}

func (h *Manager) RegisterIdEncode(idEncode extpoint.IDEncode) func() {
	h.idEncode = idEncode

	return func() {
		h.idEncode = nil
	}
}

func (h *Manager) RegisterIdStore(IdStore extpoint.IDStore) func() {
	h.idStore = IdStore

	return func(){
		h.idStore = nil
	}
}

func (h *Manager) RegisterHook(hook extpoint.Hook) func() {
	h.hooks[&hook] = hook

	return func() {
		delete(h.hooks, &hook)
	}
}

func (h *Manager) Post(ns string, longUrl string) (*domain.ShortUrl, error) {
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

func (h *Manager) Get(ns string, code string) (*domain.ShortUrl, error) {
	su, err := h.idStore.Get(ns, code)
	if err!=nil {
		return nil, err
	}

	return su, nil
}
