package dnet

import (
	"github.com/AnimusPEXUS/dnet/common_types"
)

type ModuleController struct {
	builtin_modules  map[string]common_types.ApplicationModule
	external_modules map[string]common_types.ApplicationModule
	module_instances map[string]common_types.ApplicationModuleInstance

	external_module_search_paths func() []string
}

func ModuleControllerNew(
	builtin_modules func() map[string]common_types.ApplicationModule,
	external_module_search_paths func() []string,

) *ModuleController {

	ret.builtin_modules = builtin_modules
	ret.external_module_search_paths = external_module_search_paths

	return ret
}

func (self *ModuleController) isModuleNameBuiltIn(name string) bool {
	for _, i := range self.builtin_app_modules {
		if i.Name().Value() == name {
			return true
		}
	}
	return false
}

/*
	Controller should be considered selfsaficient functionality and if
	passed `builtin' == false (with necessary checksum ofcourse), then
	Controller shoult perform search manually and retrieve it's name right from
	the module.
*/
func (self *ModuleController) EnableModule(name string, value bool) error {
	for _, i := range self.application_presets {
		if i.Name.Value() == name {
			stat, err := self.DB.GetApplicationStatus(name)
			if err != nil {
				return errors.New("can't get ApplicationStatus for named module")
			}
			i.DBStatus.Enabled = value
			stat.Enabled = value
			self.DB.SetApplicationStatus(stat)
			return nil
		}
	}
	return errors.New("so preset with this name")
}

func (self *ModuleController) AcceptModule(
	builtin bool,
	name *common_types.ModuleName,
	checksum *common_types.ModuleChecksum,
) error {

	appstat, err := self.acceptModuleNoSaveToDB(builtin, name, checksum)
	if err != nil {
		return err
	}

	err = self.DB.SetApplicationStatus(appstat)
	if err != nil {
		return err
	}

	return nil
}

// this func is for internal uses
func (self *ModuleController) acceptModuleNoSaveToDB(
	builtin bool,
	name *common_types.ModuleName,
	checksum *common_types.ModuleChecksum,
) (
	*ApplicationStatus,
	error,
) {

	if mbdn :=
		self.isModuleNameBuiltIn(name.Value()); (builtin && !mbdn) ||
		(!builtin && mbdn) {
		return nil, errors.New("trying to accept external module as builtin")
	}

	for _, i := range self.application_presets {
		if i.Name.Value() == name.Value() {
			return nil, errors.New("already have preset for module with this name")
		}
	}

	wrap, err := ControllerApplicationWrapNew(self, builtin, name, checksum)
	if err != nil {
		panic("module wrapping error: " + err.Error())
	}

	self.application_presets = append(
		self.application_presets,
		wrap,
	)

	return wrap.DBStatus, nil

}

func (self *ModuleController) RejectModule(name string) {
	self.DB.DelApplicationStatus(name)
}

func (self *ModuleController) RestorePresetsFromStorage() {

	self.application_presets = append(self.application_presets[0:0])

	names := self.DB.ListApplicationStatusNames()

	for _, i := range names {
		if dbstat, err := self.DB.GetApplicationStatus(i); err == nil {

			name_obj, err := common_types.ModuleNameNew(dbstat.Name)
			if err != nil {
				fmt.Println(
					"rejecting module " + dbstat.Name + " because name invalid",
				)
				self.RejectModule(i)
				continue
			}

			checksum_obj := (*common_types.ModuleChecksum)(nil)
			if !dbstat.Builtin {
				checksum_obj_, err :=
					common_types.ModuleChecksumNewFromString(dbstat.Checksum)
				if err != nil {
					fmt.Println(
						"rejecting module " + dbstat.Name + " because checksum invalid",
					)
					self.RejectModule(i)
					continue
				}
				checksum_obj = checksum_obj_
			}

			// TODO: error should be tracked
			self.acceptModuleNoSaveToDB(
				dbstat.Builtin,
				name_obj,
				checksum_obj,
			)

		} else {
			self.RejectModule(i)
		}
	}
}
