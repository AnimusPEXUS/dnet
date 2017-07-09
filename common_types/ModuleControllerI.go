package common_types

type ModuleControllerI interface {
	GetAcceptedModuleNameList() []string

	IsModuleBuiltin(name string) bool
	IsModuleAccepted(name string) bool

	SearchModules()
	AcceptModule(
		builtin bool,
		name *ModuleName,
		checksum *ModuleChecksum,
	) error
	RejectModule(name) error

	EnableModule(name) error
	DisableModule(name) error

	StartModuleInstance(name) error
	StopModuleInstance(name) error

	GetModule(name string) ApplicationModule
	GetModuleInstance(name string) ApplicationModuleInstance

	SaveInstances() error
	RestoreInstances() error
}
