package main

import (
	"errors"
	"fmt"

	//"github.com/AnimusPEXUS/dnet/cmd/dnetgtk/applications/builtin_dummy"
	"github.com/AnimusPEXUS/dnet/common_types"
)

//const CONFIG_DIR = "~/.config/DNet"

type Controller struct {
	//db_file  string
	//password string
	//opened   bool

	DB          *DB
	ModSearcher *ModuleSercher

	application_presets []*ControllerApplicationWrap

	window_main *UIWindowMain

	builtin_app_modules []common_types.ApplicationModule
}

func NewController(username string, key string) (*Controller, error) {

	ret := new(Controller)

	{
		t, err := NewDB(username, key)
		if err != nil {
			return nil, err
		}
		ret.DB = t
	}

	ret.builtin_app_modules = []common_types.ApplicationModule{
		//&builtin_dummy.Module{},
	}

	ret.ModSearcher = ModuleSercherNew(ret.builtin_app_modules)

	ret.application_presets = make([]*ControllerApplicationWrap, 0)

	/*
		if cntrlr, err := dnet.NewController(); err != nil {
			return nil, err
		} else {
			ret.dnet_controller = cntrlr
		}
	*/

	// Next line requires modules to be present already
	ret.RestorePresetsFromStorage()

	ret.window_main = UIWindowMainNew(ret)

	return ret, nil
}

func (self *Controller) ShowMainWindow() {
	self.window_main.Show()
	return
}

func (self *Controller) isModuleNameBuiltIn(name string) bool {
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
func (self *Controller) EnableModule(name string, value bool) error {
	for _, i := range self.application_presets {
		if i.Name.Value() == name {
			stat, err := self.DB.GetApplicationStatus(name)
			if err != nil {
				return errors.New("can't get ApplicationStatus for named module")
			}
			i.DBStatus.Enabled = value
			stat.Enabled = value
			self.DB.SetApplicationStatus(stat)
			if value {
				i.Start()
			} else {
				i.Stop()
			}
			return nil
		}
	}
	return errors.New("so preset with this name")
}

func (self *Controller) AcceptModule(
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
func (self *Controller) acceptModuleNoSaveToDB(
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

	if wrap.DBStatus.Enabled {
		wrap.Start()
	} else {
		wrap.Stop()
	}

	self.application_presets = append(
		self.application_presets,
		wrap,
	)

	return wrap.DBStatus, nil

}

func (self *Controller) RejectModule(name string) {
	self.DB.DelApplicationStatus(name)
}

func (self *Controller) RestorePresetsFromStorage() {

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
Key/ReKey code for when sqlcipher will be available for go

		_, err = db.Exec("PRAGMA key = ?;", password)
		if err != nil {
			db.Close()
			return nil, err
		}

			db.Exec("PRAGMA key = " + string(stat.DBKey))

			if time.Now >= stat.LastDBReKey+time.Duration(24*7*4)*time.Hour {
				buff := make([]byte, 255)
				rand.Read(buff)
				db.Exec("PRAGMA rekey = " + string(buff))
				stat.DBKey = string(buff)
				self.SetApplicationStatus(stat)
			}

*/
