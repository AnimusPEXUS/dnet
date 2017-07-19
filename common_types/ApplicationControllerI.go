package common_types

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
}
