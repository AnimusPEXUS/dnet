package common_types

import "net/rpc"

type ApplicationControllerI interface {
	GetBuiltinModules() ApplicationModuleMap
	GetImportedModules() ApplicationModuleMap
	GetModuleInstances() ApplicationModuleInstanceMap

	IsModuleExists(name *ModuleName) bool
	IsModuleBuiltin(name *ModuleName) bool
	GetModule(name *ModuleName) ApplicationModule

	IsInstanceExists(name *ModuleName) bool
	IsInstanceBuiltin(name *ModuleName) bool
	GetInstance(name *ModuleName) ApplicationModuleInstance

	GetInnodeRPC(
		who_asks *ModuleName,
		target_name *ModuleName,
	) (*rpc.Client, error)
}
