package main

import (
	"errors"
	"fmt"

	"github.com/AnimusPEXUS/dnet/common_types"
)

type ApplicationController struct {
	db *DB

	// Builtin modules map shoud be got via module_searcher
	module_searcher *ModuleSearcher

	// External modules should be re-searched each time then they needed.

	// One wrapper contains both Preset and Instance. Instance is made with Preset
	// if Preset.Enabled value is true.
	application_wrappers map[string]ControllerApplicationWrap
}

func NewApplicationController(
	module_searcher *ModuleSearcher,
	db *DB,
) (
	*ApplicationController,
	error,
) {
	ret := new(ApplicationController)

	ret.db = db
	ret.module_searcher = module_searcher
	ret.application_wrappers = make(map[string]ControllerApplicationWrap)

	return ret, nil
}

func (self *ApplicationController) GetInstances() map[string]common_types.ApplicationModuleInstance {
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
	*ModuleSearcherSearchResult,
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

func (self *ApplicationController) RejectModule(name string) error {
}

func (self *ApplicationController) StartModuleInstance(name string) error {
}

func (self *ApplicationController) isModuleNameBuiltIn(name string) bool {
	for _, i := range self.builtin_app_modules {
		if i.Name().Value() == name {
			return true
		}
	}
	return false
}

func (self *ApplicationController) StopModuleInstance(name string) error {
}

func (self *ApplicationController) GetModule(name string) (
	common_types.ApplicationModule,
	error,
) {
}

func (self *ApplicationController) GetModuleInstance(
	name string,
) common_types.ApplicationModuleInstance {

}

func (self *ApplicationController) Save() error {
}

func (self *ApplicationController) Load() error {

	for key, _ := range ret.application_wrappers {
		delete(ret.application_wrappers, key)
	}

	for _, i := range self.DB.ListApplicationStatusNames() {

		if dbstat, err := self.DB.GetApplicationStatus(i); err != nil {

			self.RejectModule(i)

		} else {

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
						"rejecting module", dbstat.Name, "because checksum invalid",
					)
					self.RejectModule(i)
					continue
				}
				checksum_obj = checksum_obj_
			}

			// TODO: error should be tracked
			self.AcceptModuleNoSaveToDB(
				dbstat.Builtin,
				name_obj,
				checksum_obj,
			)

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

func (self *ApplicationController) DisableModule(name string) error {
}

func (self *Controller) RejectModule(name string) {
	self.DB.DelApplicationStatus(name)
}
