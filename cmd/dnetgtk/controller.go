package main

import (
	"errors"

	"bitbucket.org/AnimusPEXUS/dnet"
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

	ret.window_main = UIWindowMainNew(ret)

	if cntrlr, err := dnet.NewController(); err != nil {
		return nil, err
	} else {
		ret.dnet_controller = cntrlr
	}

	ret.network_presets = make([]*ControllerNetworkPresetAndInstance, 0)

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

func (self *Controller) NetworkPresetsCopy() []*ControllerNetworkPresetExt {
	ret := make([]*ControllerNetworkPresetExt, len(self.network_presets))
	for _, i := range self.network_presets {
		ret = append(ret, ControllerNetworkPresetExtNew(&i.ControllerNetworkPreset))
	}
	return ret
}

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
	return nil
}

func (self *Controller) AddNetworkPreset(
	name string,
	module string,
	config string,
) error {

	if self.HasNetworkPresetName(name) {
		return errors.New("Already has this preset")
	}

	cnp := ControllerNetworkPresetNew(
		name,
		module,
		false,
		config,
	)

	a := &ControllerNetworkPresetAndInstance{
		ControllerNetworkPreset: *cnp,
		NetworkInstance:         nil,
	}

	self.network_presets = append(self.network_presets, a)

	return nil
}

func (self *Controller) ChangeNetworkPresetConfig(
	name string,
	value string,
) error {
	ret := error(nil)
	for _, i := range self.network_presets {
		if i.Name() == name {
			if !i.NetworkInstance.Status().Stopped() {
				ret = errors.New("Can't change while not stopped")
			} else {
				i.SetConfig(value)
				ret = nil
			}
			break
		}
	}

	return ret
}
