package core

import (
	"github.com/freepai/hummer/core/domain"
	"github.com/freepai/hummer/core/extension"
)

type Hummer struct {
	idGen    extension.IdGen
	idEncode extension.IDEncode
	idStore  extension.IDStore
	hooks    map[*extension.Hook]extension.Hook
}

func NewHummer() *Hummer {
	return &Hummer{}
}

func (h *Hummer) RegisterIdGen(idGen extension.IdGen) func() {
	h.idGen = idGen

	return func(){
		h.idGen = nil
	}
}

func (h *Hummer) RegisterIdEncode(idEncode extension.IDEncode) func() {
	h.idEncode = idEncode

	return func() {
		h.idEncode = nil
	}
}

func (h *Hummer) RegisterIdStore(IdStore extension.IDStore) func() {
	h.idStore = IdStore

	return func(){
		h.idStore = nil
	}
}

func (h *Hummer) RegisterHook(hook extension.Hook) func() {
	h.hooks[&hook] = hook

	return func() {
		delete(h.hooks, &hook)
	}
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