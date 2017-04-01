package main

import (
	"bitbucket.org/AnimusPEXUS/dnet/common_types"
)

type ControllerNetworkPresetAndInstance struct {
	ControllerNetworkPreset
	common_types.NetworkInstance
}

type ControllerNetworkPresetExt struct {
	a *ControllerNetworkPreset
}

func ControllerNetworkPresetExtNew(
	orig *ControllerNetworkPreset,
) *ControllerNetworkPresetExt {
	ret := new(ControllerNetworkPresetExt)
	ret.a = orig
	return ret
}

func (self *ControllerNetworkPresetExt) Name() string {
	return self.a.Name()
}

func (self *ControllerNetworkPresetExt) Module() string {
	return self.a.Module()
}

func (self *ControllerNetworkPresetExt) AutoStart() bool {
	return self.a.AutoStart()
}

func (self *ControllerNetworkPresetExt) Config() string {
	return self.a.Config()
}

type ControllerNetworkPreset struct {
	name      string
	module    string
	autostart bool
	config    string
}

func ControllerNetworkPresetNew(
	name string,
	module string,
	autostart bool,
	config string,
) *ControllerNetworkPreset {
	ret := new(ControllerNetworkPreset)
	ret.name = name
	ret.autostart = autostart
	ret.config = config
	return ret
}

func (self *ControllerNetworkPreset) Name() string {
	return self.name
}

func (self *ControllerNetworkPreset) Module() string {
	return self.module
}

func (self *ControllerNetworkPreset) SetModule(value string) {
	self.module = value
	return
}

func (self *ControllerNetworkPreset) AutoStart() bool {
	return self.autostart
}

func (self *ControllerNetworkPreset) SetAutoStart(value bool) {
	self.autostart = value
	return
}

func (self *ControllerNetworkPreset) Config() string {
	return self.config
}

func (self *ControllerNetworkPreset) SetConfig(value string) {
	self.config = value
	return
}
