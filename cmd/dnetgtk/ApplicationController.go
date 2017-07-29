package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/rpc"
	"time"

	"github.com/AnimusPEXUS/dnet"
	"github.com/AnimusPEXUS/dnet/common_types"
	"github.com/AnimusPEXUS/goset"
	"github.com/AnimusPEXUS/goworker"
	"github.com/gotk3/gotk3/gtk"
)

type SafeApplicationModuleInstanceWrap struct {
	//Name     *common_types.ModuleName
	//Builtin  bool
	Module   common_types.ApplicationModule
	Instance common_types.ApplicationModuleInstance

	// It is required to put and use this flag, because if using .Instance ==/!=
	// nil to determine if module enabled, then it is mutch harder to check
	// instance worker status + it is harder to trace what wall Instance's windows
	// are closed then trying to disable module by assigning nil to .Instace
	Enabled bool
}

type ApplicationController struct {
	*worker.Worker

	controller *Controller
	db         *DB

	// Builtin modules map shoud be got via module_searcher
	module_searcher *ModuleSearcher

	// External modules should be re-searched each time then they needed.

	// One wrapper contains both Preset and Instance. Instance is made with Preset
	// if Preset.Enabled value is true.
	application_wrappers map[string]*SafeApplicationModuleInstanceWrap

	module_instance_status_display_map map[string]string
}

func (self *ApplicationController) threadWorker(

	set_starting func(),
	set_working func(),
	set_stopping func(),
	set_stopped func(),

	is_stop_flag func() bool,

) {

	for !is_stop_flag() {

		if self.controller != nil &&
			self.controller.window_main != nil &&
			self.controller.window_main.UIWindowMainTabApplications != nil &&
			self.controller.window_main.UIWindowMainTabApplications.
				button_accept_application != nil {
			self.RefreshAllAcceptedApplicationListItems(true)
		}

		time.Sleep(time.Second)
	}
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

	ret.Worker = worker.New(ret.threadWorker)

	ret.controller = controller
	ret.db = db
	ret.module_searcher = module_searcher
	ret.application_wrappers = make(map[string]*SafeApplicationModuleInstanceWrap)

	return ret, nil
}

func (
	self *ApplicationController,
) RefreshAllAcceptedApplicationListItems(no_db bool) {
	set := goset.NewSetString()

	{
		presets_mdl := self.controller.window_main.UIWindowMainTabApplications.
			application_presets
		wrappers := self.application_wrappers

		for i, _ := range wrappers {
			set.Add(i)
		}

		iter, ok := presets_mdl.GetIterFirst()
		for ok {
			val, _ := presets_mdl.GetValue(iter, 0)
			val_str, _ := val.GetString()
			set.Add(val_str)
			ok = presets_mdl.IterNext(iter)
		}
	}

	for _, i := range set.List() {
		self.RefreshAcceptedApplicationListItem(
			common_types.ModuleNameNewF(i.(string)),
			no_db,
		)
	}

}

func (self *ApplicationController) RefreshAcceptedApplicationListItem(
	name *common_types.ModuleName,
	no_db bool,
) {
	if self.controller != nil &&
		self.controller.window_main != nil &&
		self.controller.window_main.UIWindowMainTabApplications != nil {
		self.controller.window_main.UIWindowMainTabApplications.
			RefreshAppPresetListItem(
				name.Value(),
				no_db,
			)
	}
}

// ----------- Interface Part -----------

func (
	self *ApplicationController,
) GetBuiltinModules() common_types.ApplicationModuleMap {
	return self.module_searcher.builtin
}

func (
	self *ApplicationController,
) GetImportedModules() common_types.ApplicationModuleMap {
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
	defer self.RefreshAcceptedApplicationListItem(name, false)

	{ // security and sanity checks
		if _, ok := self.application_wrappers[name.Value()]; ok {
			return errors.New("already have preset for module with same name")
		}

		if t :=
			self.IsModuleBuiltin(name); (builtin && !t) || (!builtin && t) {
			return errors.New("trying to accept external module as builtin")
		}

		search_res, err :=
			self.module_searcher.SearchMod(builtin, name, checksum)
		if err != nil {
			return errors.New("can't find module: " + err.Error())
		}

		if builtin != search_res.Builtin() {
			return errors.New("addional builtin != builtin safety check failure")
		}
	}

	{ // Save to database if needed
		if save_to_db {
			appstat, err := self.db.GetApplicationStatus(name.Value())
			if appstat == nil {
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
				if needs, err := self.IsModuleNeedsReKey(name); err != nil {
					return err
				} else {
					if needs {
						self.ModuleReKey(name)
					}
				}
			}
		}
	}

	{ // Create corresponding Wrap
		module, err :=
			self.module_searcher.GetMod(builtin, name, checksum)
		if err != nil {
			return errors.New("can't get module: " + err.Error())
		}

		wrap := new(SafeApplicationModuleInstanceWrap)
		wrap.Module = module

		self.application_wrappers[name.Value()] = wrap
	}

	{ // Instantiate module and give it new Communicator
		appstat, err := self.db.GetApplicationStatus(name.Value())
		if err != nil {
			return err
		}

		wrap := self.application_wrappers[name.Value()]
		db, err := self.db.GetAppDB(name.Value())
		if err != nil {
			return errors.New("Error getting DB connection: " + err.Error())
		}

		db.Key(appstat.DBKey)

		cc := &ControllerCommunicatorForApp{
			name:       name,
			controller: self.controller,
			wrap:       wrap,
			db:         db.db,
		}

		cc.logger = &ControllerCommunicatorForAppLogger{p: cc}

		if ins, err := wrap.Module.Instantiate(cc); err != nil {
			return errors.New("Error instantiating module " + name.Value())
		} else {
			wrap.Instance = ins
		}
	}

	{
		wrap := self.application_wrappers[name.Value()]

		if wrap.Module.IsWorker() {
			dbstat, err := self.db.GetApplicationStatus(name.Value())
			if err != nil {
				return err
			}
			self.SetModuleEnabled(
				name,
				dbstat.Enabled,
				false,
			)
		}
	}

	return nil
}

func (self *ApplicationController) RejectModule(
	name *common_types.ModuleName,
) error {
	v := name.Value()
	if _, ok := self.application_wrappers[v]; ok {
		delete(self.application_wrappers, v)
	}
	self.db.DelApplicationStatus(v)
	return nil
}

func (self *ApplicationController) Load() error {

	for key, _ := range self.application_wrappers {
		delete(self.application_wrappers, key)
	}

	for _, i := range self.db.ListApplicationStatusNames() {

		i_obj, err := common_types.ModuleNameNew(i)
		if err != nil {
			self.controller.logger.Error(
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
					self.controller.logger.Error(
						"rejecting module " + dbstat.Name + " because name invalid",
					)
					self.RejectModule(i_obj)
					continue
				}

				checksum_obj = (*common_types.ModuleChecksum)(nil)
				if !dbstat.Builtin {
					checksum_obj_, err :=
						common_types.ModuleChecksumNewFromString(dbstat.Checksum)
					if err != nil {
						self.controller.logger.Error(
							"rejecting module" + dbstat.Name + "because checksum invalid",
						)
						self.RejectModule(i_obj)
						continue
					}
					checksum_obj = checksum_obj_
				}
			}

			// TODO: error should be tracked
			err = self.AcceptModule(
				dbstat.Builtin,
				name_obj,
				checksum_obj,
				false,
			)
			if err != nil {
				self.controller.logger.Error(
					"error accepting module " + name_obj.Value(),
				)
			} else {
				self.controller.logger.Info(
					"accepted module " + name_obj.Value(),
				)
			}

		}
	}
	return nil
}

func (self *ApplicationController) GetModuleEnabled(
	name *common_types.ModuleName,
) (bool, error) {
	v := name.Value()
	if _, ok := self.application_wrappers[v]; ok {
		stat, err := self.db.GetApplicationStatus(v)
		if err != nil {
			return false, errors.New("can't get ApplicationStatus for named module")
		}
		return stat.Enabled, nil
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
	defer self.RefreshAcceptedApplicationListItem(name, false)

	key := name.Value()
	if val, ok := self.application_wrappers[key]; ok {
		stat, err := self.db.GetApplicationStatus(key)
		if err != nil {
			return errors.New("can't get ApplicationStatus for named module")
		}
		stat.Enabled = value
		val.Enabled = value

		if value {

			if val.Module.IsWorker() {
				go val.Instance.Start()
			}

		} else {
			if val.Module.IsWorker() {
				go val.Instance.Stop()
			}
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

	return errors.New("named module not found. possible programming error")
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
	key := name.Value()
	if val, ok := self.application_wrappers[key]; ok {
		return val.Module.HaveUI(), nil
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
	key := name.Value()
	if val, ok := self.application_wrappers[key]; ok {
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

func (self *ApplicationController) GetInnodeRPC(
	who_asks *common_types.ModuleName,
	target_name *common_types.ModuleName,
) (*rpc.Client, error) {
	if target_name.Value() == dnet.DNET_UNIVERSAL_APPLICATION_NAME {
		return self.controller.dnet_controller.GetInnodeRPC(who_asks)
	} else {
		if inst, ok := self.application_wrappers[target_name.Value()]; !ok {
			return nil, errors.New("module not found")
		} else {
			return inst.Instance.GetInnodeRPC(who_asks.Value())
		}
	}
}
