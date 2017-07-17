package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/AnimusPEXUS/dnet/common_types"
	"github.com/gotk3/gotk3/gtk"
)

type SafeApplicationModuleInstanceWrap struct {
	//Name     *common_types.ModuleName
	//Builtin  bool
	Module   common_types.ApplicationModule
	Instance common_types.ApplicationModuleInstance
}

type ApplicationController struct {
	controller *Controller
	db         *DB

	// Builtin modules map shoud be got via module_searcher
	module_searcher *ModuleSearcher

	// External modules should be re-searched each time then they needed.

	// One wrapper contains both Preset and Instance. Instance is made with Preset
	// if Preset.Enabled value is true.
	application_wrappers map[string]*SafeApplicationModuleInstanceWrap
}

func NewApplicationController(
	controller *Controller,
	module_searcher *ModuleSearcher,
	db *DB,
) (
	*ApplicationController,
	error,
) {
	ret := new(ApplicationController)

	ret.controller = controller
	ret.db = db
	ret.module_searcher = module_searcher
	ret.application_wrappers = make(map[string]*SafeApplicationModuleInstanceWrap)

	return ret, nil
}

// ----------- Interface Part -----------

func (self *ApplicationController) GetBuiltinModules() common_types.ApplicationModuleMap {
	return self.module_searcher.builtin
}

func (self *ApplicationController) GetImportedModules() common_types.ApplicationModuleMap {
	builtins := self.GetBuiltinModules()
	ret := make(common_types.ApplicationModuleMap)
search:
	for key, val := range self.application_wrappers {
		for key2, _ := range builtins {
			if key == key2 {
				continue search
			}
		}
		ret[key] = val.Module
	}
	return ret
}

func (
	self *ApplicationController,
) GetModuleInstances() common_types.ApplicationModuleInstanceMap {
	ret := make(common_types.ApplicationModuleInstanceMap)
	for key, val := range self.application_wrappers {
		if val.Instance != nil {
			ret[key] = val.Instance
		}
	}
	return ret
}

func (self *ApplicationController) IsModuleExists(
	name *common_types.ModuleName,
) bool {
	return false
}

func (self *ApplicationController) IsModuleBuiltin(
	name *common_types.ModuleName,
) bool {
	for i, _ := range self.module_searcher.builtin {
		if i == name.Value() {
			return true
		}
	}
	return false

}

func (self *ApplicationController) GetModule(
	name *common_types.ModuleName,
) common_types.ApplicationModule {
	return nil
}

func (self *ApplicationController) IsInstanceExists(
	name *common_types.ModuleName,
) bool {
	return false
}

func (self *ApplicationController) IsInstanceBuiltin(
	name *common_types.ModuleName,
) bool {
	return false
}

func (self *ApplicationController) GetInstance(
	name *common_types.ModuleName,
) common_types.ApplicationModuleInstance {
	return nil
}

// ----------- Implimentation Part -----------

func (self *ApplicationController) AcceptModule(
	builtin bool,
	name *common_types.ModuleName,
	checksum *common_types.ModuleChecksum,
	save_to_db bool,
) error {

	// TODO: security and sanity checks

	for key, _ := range self.application_wrappers {
		if key == name.Value() {
			return errors.New("already have preset for module with same name")
		}
	}

	if t :=
		self.IsModuleBuiltin(name); (builtin && !t) || (!builtin && t) {
		return errors.New("trying to accept external module as builtin")
	}

	{
		search_res, err :=
			self.module_searcher.SearchMod(builtin, name, checksum)
		if err != nil {
			return errors.New("can't find module: " + err.Error())
		}

		if builtin != search_res.Builtin() {
			return errors.New("addional builtin != builtin safety check failure")
		}
	}

	{
		module, err :=
			self.module_searcher.GetMod(builtin, name, checksum)
		if err != nil {
			return errors.New("can't get module: " + err.Error())
		}

		wrap := new(SafeApplicationModuleInstanceWrap)
		wrap.Module = module

		self.application_wrappers[name.Value()] = wrap
	}

	{
		appstat, err := self.db.GetApplicationStatus(name.Value())
		if err == nil {
			app_db, err := self.db.GetAppDB(name.Value())
			if err != nil {
				return err
			}

			needs, err := self.IsModuleNeedsReKey(name)
			if err != nil {
				return err
			}

			if !needs {
				app_db.Key(appstat.DBKey)
			}
		}
	}

	if save_to_db {
		appstat, err := self.db.GetApplicationStatus(name.Value())
		if err != nil {
			appstat = &ApplicationStatus{}
			appstat.Name = name.Value()
			appstat.Enabled = false
		}
		appstat.Builtin = builtin
		if checksum != nil {
			appstat.Checksum = checksum.String()
		}

		err = self.db.SetApplicationStatus(appstat)
		if err != nil {
			return err
		}

		{
			needs, err := self.IsModuleNeedsReKey(name)
			if err != nil {
				return err
			}

			if needs {
				self.ModuleReKey(name)
			}
		}

	}

	return nil
}

func (self *ApplicationController) RejectModule(
	name *common_types.ModuleName,
) error {
	for key, _ := range self.application_wrappers {
		if key == name.Value() {
			delete(self.application_wrappers, key)
			break
		}
	}
	self.db.DelApplicationStatus(name.Value())
	return nil
}

func (self *ApplicationController) Load() error {

	for key, _ := range self.application_wrappers {
		delete(self.application_wrappers, key)
	}

	for _, i := range self.db.ListApplicationStatusNames() {

		i_obj, err := common_types.ModuleNameNew(i)
		if err != nil {
			fmt.Println(
				"rejecting module invalid with invalid name",
			)
			self.db.DelApplicationStatus(i)
			//self.RejectModule(i)
			continue
		}

		if dbstat, err := self.db.GetApplicationStatus(i_obj.Value()); err != nil {

			self.RejectModule(i_obj)

		} else {

			var (
				name_obj     *common_types.ModuleName
				checksum_obj *common_types.ModuleChecksum
			)
			{
				var err error

				name_obj, err = common_types.ModuleNameNew(dbstat.Name)
				if err != nil {
					fmt.Println(
						"rejecting module ", dbstat.Name, " because name invalid",
					)
					self.RejectModule(i_obj)
					continue
				}

				checksum_obj = (*common_types.ModuleChecksum)(nil)
				if !dbstat.Builtin {
					checksum_obj_, err :=
						common_types.ModuleChecksumNewFromString(dbstat.Checksum)
					if err != nil {
						fmt.Println(
							"rejecting module", dbstat.Name, "because checksum invalid",
						)
						self.RejectModule(i_obj)
						continue
					}
					checksum_obj = checksum_obj_
				}
			}

			// TODO: error should be tracked
			self.AcceptModule(
				dbstat.Builtin,
				name_obj,
				checksum_obj,
				false,
			)

			self.SetModuleEnabled(
				name_obj,
				dbstat.Enabled,
				false,
			)

		}
	}
	return nil
}

func (self *ApplicationController) GetModuleEnabled(
	name *common_types.ModuleName,
) (bool, error) {
	for key, _ := range self.application_wrappers {
		if key == name.Value() {

			stat, err := self.db.GetApplicationStatus(key)
			if err != nil {
				return false, errors.New("can't get ApplicationStatus for named module")
			}
			return stat.Enabled, nil

			break
		}
	}
	return false, errors.New("module not found")
}

/*
	Controller should be considered selfsaficient functionality and if
	passed `builtin' == false (with necessary checksum ofcourse), then
	Controller shoult perform search manually and retrieve it's name right from
	the module.
*/
func (self *ApplicationController) SetModuleEnabled(
	name *common_types.ModuleName,
	value bool,
	save_to_db bool,
) error {
	for key, val := range self.application_wrappers {
		if key == name.Value() {
			stat, err := self.db.GetApplicationStatus(key)
			if err != nil {
				return errors.New("can't get ApplicationStatus for named module")
			}
			stat.Enabled = value

			if stat.Enabled {

				db, err := self.db.GetAppDB(name.Value())
				if err != nil {
					return errors.New("Error getting DB connection: " + err.Error())
				}

				cc := &ControllerCommunicatorForApp{
					name:       name,
					controller: self.controller,
					wrap:       val,
					db:         db.db,
				}

				if ins, err := val.Module.Instance(cc); err != nil {
					return errors.New("Error instantiating module " + name.Value())
				} else {
					val.Instance = ins
				}
			} else {
				// TODO: possibly Instance have to have Destroy method. clear this out
				// val.Instance.Destroy()
				val.Instance = nil
			}

			if save_to_db {
				stat, err := self.db.GetApplicationStatus(name.Value())
				if err != nil {
					// TODO: possibly in this case, some additional action should be done
					return errors.New(
						"Can't change enabled status for module, which isn't registered",
					)
				}
				stat.Enabled = value
				self.db.SetApplicationStatus(stat)
			}

			return nil
		}
	}
	return errors.New("named module not found")
}

func (self *ApplicationController) GetModuleStatus(
	name *common_types.ModuleName,
) (
	*ApplicationStatus,
	error,
) {
	ret, err := self.db.GetApplicationStatus(name.Value())
	if err != nil {
		return nil, errors.New("can't get application status from storage")
	}

	return ret, err
}

func (self *ApplicationController) SetModuleStatus(
	status *ApplicationStatus,
) error {
	return self.db.SetApplicationStatus(status)
}

func (self *ApplicationController) ModuleHaveUI(
	name *common_types.ModuleName,
) (bool, error) {
	for key, val := range self.application_wrappers {
		if key == name.Value() {
			return val.Module.HaveUI(), nil
		}
	}
	return false, errors.New("module not found")
}

func (self *ApplicationController) ModuleShowUI(
	name *common_types.ModuleName,
) error {
	ok, err := self.ModuleHaveUI(name)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("module have no UI")
	}

	for key, val := range self.application_wrappers {
		if key == name.Value() {
			if val.Instance == nil {
				return errors.New("Module not instantiated, so can't get it's UI")
			}
			ui, err := val.Instance.GetUI(nil)
			if err != nil {
				return err
			}
			switch ui.(type) {

			case interface {
				Show() error
			}:
				return ui.(interface {
					Show() error
				}).Show()

			case interface {
				Get() (*gtk.Window, error)
			}:
				wind, err := ui.(interface {
					Get() (*gtk.Window, error)
				}).Get()

				if err != nil {
					return errors.New(
						"Trying to get gtk.Window from module '" + key +
							"' resulted in error:\n" +
							err.Error(),
					)
				}

				wind.ShowAll()

			default:
				return errors.New(
					"ApplicationController doesn't know how to handle '" + key +
						"' module window, or said module doesn't have window at all\n" +
						"This should be considered programming error ether of module " +
						"ether of DNetGtk",
				)
			}
			return nil
		}
	}

	return errors.New(
		"some unknown error. this shouldn't been happen. contact developer",
	)
}

func (self *ApplicationController) IsModuleNeedsReKey(
	name *common_types.ModuleName,
) (bool, error) {
	stat, err := self.GetModuleStatus(name)
	if err != nil {
		return false, err
	}
	return stat.DBKey == "" || stat.LastDBReKey == nil, nil
}

func (self *ApplicationController) ModuleReKey(
	name *common_types.ModuleName,
) error {
	app_db, err := self.db.GetAppDB(name.Value())
	if err != nil {
		return err
	}

	stat, err := self.GetModuleStatus(name)
	if err != nil {
		return err
	}
	{
		buff := make([]byte, 50)
		rand.Read(buff)
		buff_str := base64.RawStdEncoding.EncodeToString(buff)
		stat.DBKey = buff_str

		app_db.ReKey(buff_str)
		app_db.Key(buff_str)
	}
	t := time.Now().UTC()
	stat.LastDBReKey = &t
	err = self.SetModuleStatus(stat)
	if err != nil {
		return err
	}
	return nil
}
