package server

import "github.com/freepai/hummer/core/plugin"

func GetManager(ctx *plugin.Context) *Manager {
	return ctx.GetBean(ManagerName).(*Manager)
}

func GetManagerFromRegistry(registry plugin.Registry) *Manager {
	return registry.GetBean(ManagerName).(*Manager)
}
