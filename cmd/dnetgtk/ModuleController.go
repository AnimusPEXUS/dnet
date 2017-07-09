package dnetgtk

type ModuleController struct {
}

func (self *ModuleController) GetAcceptedModuleNameList() []string {
}

func (self *ModuleController) IsModuleBuiltin(name string) bool {
}

func (self *ModuleController) IsModuleAccepted(name string) bool {
}

func (self *ModuleController) SearchModules() {
}

func (self *ModuleController) AcceptModule(
	builtin bool,
	name *ModuleName,
	checksum *ModuleChecksum,
) error {
}

func (self *ModuleController) RejectModule(name) error {
}

func (self *ModuleController) EnableModule(name) error {
}

func (self *ModuleController) DisableModule(name) error {
}

func (self *ModuleController) StartModuleInstance(name) error {
}

func (self *ModuleController) StopModuleInstance(name) error {
}

func (self *ModuleController) GetModule(name string) ApplicationModule {
}

func (self *ModuleController) GetModuleInstance(
	name string,
) ApplicationModuleInstance {
}

func (self *ModuleController) SaveInstances() error {
}

func (self *ModuleController) RestoreInstances() error {
}
