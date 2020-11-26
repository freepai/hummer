package plugin

import "github.com/freepai/hummer/core/plugin/api"

type Context interface {
	RegisterIdGen(idGen api.IdGen) error
	RegisterIdEncode(idEncode api.IDEncode) error
	RegisterIdStore(IdStore api.IDStore) error

	RegisterHook(hooks api.Hook) error
	RegisterEndpoint(endpoint api.Endpoint) error
}
