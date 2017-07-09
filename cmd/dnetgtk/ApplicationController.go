package main

import (
	"errors"
	"fmt"

	"github.com/AnimusPEXUS/dnet/common_types"
)

type ApplicationController struct {
	//builtin_modules map[string]common_types.ApplicationModule

	//external_modules map[string]common_types.ApplicationModule

	application_instances map[string]ControllerApplicationWrap

	module_searcher *ModuleSercher

	db *DB

	//application_presets []*ControllerApplicationWrap
}

func NewApplicationController(
	module_searcher *ModuleSercher,
	db *DB,
) (
	*ApplicationController,
	error,
) {
	ret := new(ModuleController)
	//ret.builtin_modules = builtin_modules
	ret.module_searcher = module_searcher
	ret.db = db
	return ret, nil
}

func (self *ApplicationController) GetInstances()
map[string]common_types.ApplicationModuleInstance {
	return self.module_instances
}

func (self *ApplicationController) IsModuleBuiltin(name string) bool {
	for i, _ := range self.module_searcher.builtin {
		if i == name {
			return true
		}
	}
	return false
}

func (self *ApplicationController) IsModuleAccepted(name string) bool {
}

func (self *ApplicationController) SearchModules(
	builtin bool,
	name *common_types.ModuleName,
	checksum *common_types.ModuleChecksum,
) (
	*ModuleSercherSearchResult,
	error,
) {
	return self.module_searcher.SearchMod(builtin, name, checksum)
}

func (self *ApplicationController) AcceptModule(
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

// this func is for internal uses (when restoring from DB and don't need to
// resave)
func (self *ApplicationController) AcceptModuleNoSaveToDB(
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

func (self *ApplicationController) RejectModule(name) error {
}

func (self *ApplicationController) EnableModule(name) error {
}

func (self *ApplicationController) DisableModule(name) error {
}

func (self *ApplicationController) StartModuleInstance(name) error {
}

func (self *ApplicationController) isModuleNameBuiltIn(name string) bool {
	for _, i := range self.builtin_app_modules {
		if i.Name().Value() == name {
			return true
		}
	}
	return false
}



func (self *ApplicationController) StopModuleInstance(name) error {
}

func (self *ApplicationController) GetModule(name string) (
	common_types.ApplicationModule,
	error,
) {
}

func (self *ApplicationController) GetModuleInstance(
	name string,
) ApplicationModuleInstance {
}

func (self *ApplicationController) SaveInstances() error {
}

func (self *ApplicationController) RestoreInstances() error {

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

/*
	Controller should be considered selfsaficient functionality and if
	passed `builtin' == false (with necessary checksum ofcourse), then
	Controller shoult perform search manually and retrieve it's name right from
	the module.
*/
func (self *ApplicationController) EnableModule(name string, value bool) error {
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

func (self *Controller) RejectModule(name string) {
	self.DB.DelApplicationStatus(name)
}
