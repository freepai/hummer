package plugin

import "github.com/freepai/hummer/core/plugin/api"

type Context interface {
	RegisterIdGen(name string, idGen api.IdGen) error
	RegisterIdEncode(name string, idEncode api.IDEncode) error
	RegisterIdStore(name string, IdStore api.IDStore) error
}
