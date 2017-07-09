package common_types

type ApplicationControllerI interface {
	GetAcceptedModuleNameList() []string

	IsModuleBuiltin(name string) bool
	IsModuleAccepted(name string) bool

	SearchModules()
	AcceptModule(
		builtin bool,
		name *ModuleName,
		checksum *ModuleChecksum,
	) error
	RejectModule(name string) error

	EnableModule(name string) error
	DisableModule(name string) error

	StartModuleInstance(name string) error
	StopModuleInstance(name string) error

	GetModule(name string) ApplicationModule
	GetModuleInstance(name string) ApplicationModuleInstance

	SaveInstances() error
	RestoreInstances() error

	ShowUI(module_name ModuleName) error
}
