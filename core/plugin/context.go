package plugin

import "github.com/freepai/hummer/core/plugin/api"

type Context interface {
	RegisterIdGen(idGen api.IdGen) func()
	RegisterIdEncode(idEncode api.IDEncode) func()
	RegisterIdStore(IdStore api.IDStore) func()
	RegisterHook(hooks api.Hook) func()
}
