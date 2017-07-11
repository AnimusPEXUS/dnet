package common_types

type ApplicationControllerI interface {
	GetBuiltinModules() []ApplicationModule
	GetImportedModules() []ApplicationModule
	GetModuleInstances() []ApplicationModuleInstance

	IsModuleExists(name *ModuleName) bool
	IsModuleBuiltin(name *ModuleName) bool
	GetModule(name *ModuleName) ApplicationModule

	IsInstanceExists(name *ModuleName) bool
	IsInstanceBuiltin(name *ModuleName) bool
	GetInstance(name *ModuleName) ApplicationModuleInstance
}
