package main

import (
	"errors"
	"fmt"

	"bitbucket.org/AnimusPEXUS/dnet"
	"bitbucket.org/AnimusPEXUS/dnet/application_modules/pm"
	"bitbucket.org/AnimusPEXUS/dnet/network_modules/tcpip"
)

const CONFIG_DIR = "~/.config/DNet"

type Controller struct {
	//db_file  string
	//password string
	//opened   bool

	DB *DB

	network_presets []*ControllerNetworkPresetAndInstance

	dnet_controller *dnet.Controller

	window_main *UIWindowMain
}

func NewController(db_file string, password string) (*Controller, error) {

	ret := new(Controller)

	{
		t, err := NewDB(db_file, password)
		if err != nil {
			return nil, err
		}
		ret.DB = t
	}

	ret.network_presets = make([]*ControllerNetworkPresetAndInstance, 0)

	if cntrlr, err := dnet.NewController(); err != nil {
		return nil, err
	} else {
		ret.dnet_controller = cntrlr
	}

	// network modules
	ret.dnet_controller.AppendNetworkModule(tcpip.NewModule())

	// application modules
	ret.dnet_controller.AppendApplicationModule(pm.NewModule())

	// Next line requires modules to be present already
	ret.RestorePresetsFromStorage()

	ret.window_main = UIWindowMainNew(ret)

	return ret, nil
}

/*
func (self *Controller) IsOpened() bool {
	return self.opened
}
*/

func (self *Controller) DNet() *dnet.Controller {
	return self.dnet_controller
}

func (self *Controller) ShowMainWindow() {
	self.window_main.Show()
	return
}

/*
func (self *Controller) NetworkPresetsCopy() []*ControllerNetworkPresetExt {
	ret := make([]*ControllerNetworkPresetExt, len(self.network_presets))
	for _, i := range self.network_presets {
		ret = append(ret, ControllerNetworkPresetExtNew(&i.ControllerNetworkPreset))
	}
	return ret
}
*/

func (self *Controller) NetworkPresetList() []string {
	ret := make([]string, 0)
	for _, i := range self.network_presets {
		ret = append(ret, i.Name())
	}
	return ret
}

func (self *Controller) HasNetworkPresetName(name string) bool {
	for _, i := range self.network_presets {
		if i.Name() == name {
			return true
		}
	}
	return false
}

func (self *Controller) GasNetworkPresetByName(
	name string,
) *ControllerNetworkPresetAndInstance {
	for _, i := range self.network_presets {
		if i.Name() == name {
			return i
		}
	}
	return nil
}

func (self *Controller) DeleteNetworkPreset(name string) error {
	for ii, i := range self.network_presets {
		if i.Name() == name {
			if !i.NetworkInstance.Status().Stopped() {
				return errors.New(
					"Preset instance have to be stopped, before deleting it's preset",
				)
			}
			self.network_presets = append(
				// self.network_presets[0:0],
				self.network_presets[0:ii],
				self.network_presets[ii+1:len(self.network_presets)]...,
			)
			break
		}
	}
	self.DB.DelNetPreset(name)
	return nil
}

func (self *Controller) addNetworkPreset(
	name string,
	module string,
	enabled bool,
	config string,
	avoid_db_save bool,
) error {

	if self.HasNetworkPresetName(name) {
		return errors.New(
			"already has preset with same name." +
				" choose new name or edit existing.",
		)
	}

	cnp := ControllerNetworkPresetNew(
		name,
		module,
		enabled,
		config,
	)

	if a, err := ControllerNetworkPresetAndInstanceNew(
		cnp,
		self.dnet_controller,
	); err != nil {
		panic("can't ControllerNetworkPresetAndInstanceNew(): " + err.Error())
	} else {
		self.network_presets = append(self.network_presets, a)
	}

	if avoid_db_save == false {
		self.DB.SetNetPreset(
			name,
			module,
			enabled,
			config,
		)
	}

	return nil
}

func (self *Controller) AddNetworkPreset(
	name string,
	module string,
	enabled bool,
	config string,
) error {
	return self.addNetworkPreset(name, module, enabled, config, false)
}

func (self *Controller) ChangeNetworkPresetEnabled(
	name string,
	value bool,
) error {
	ret := error(nil)
	for _, i := range self.network_presets {
		if i.Name() == name {

			if value {
				i.NetworkInstance.Start()
			} else {
				i.NetworkInstance.Stop()
			}

			_, module, _, config := self.DB.GetNetPreset(name)

			self.DB.SetNetPreset(
				name,
				module,
				value,
				config,
			)

			break

		}
	}

	return ret
}

func (self *Controller) ChangeNetworkPresetConfig(
	name string,
	config string,
) error {
	ret := error(nil)
	for _, i := range self.network_presets {
		if i.Name() == name {
			if !i.NetworkInstance.Status().Stopped() {
				ret = errors.New("Can't change while not stopped")
			} else {
				i.NetworkInstance.SetConfig(config)

				_, module, enabled, _ := self.DB.GetNetPreset(name)

				self.DB.SetNetPreset(
					name,
					module,
					enabled,
					config,
				)

				ret = nil
			}
			break
		}
	}

	return ret
}

func (self *Controller) RestorePresetsFromStorage() {
	lst := self.DB.LstNetPresets()
	for _, name := range lst {
		found, module, enabled, config := self.DB.GetNetPreset(name)
		if !found {
			fmt.Println("error: storage does not have such network preset")
		} else {
			self.addNetworkPreset(
				name,
				module,
				enabled,
				config,
				true,
			)
		}
	}

}
