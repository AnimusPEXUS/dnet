package main

import (
	"errors"
	"fmt"

	"bitbucket.org/AnimusPEXUS/dnet"
	"bitbucket.org/AnimusPEXUS/dnet/common_types"
)

type ControllerApplicationPresetAndInstance struct {
	ControllerApplicationPreset
	common_types.ApplicationInstance
}

func ControllerApplicationPresetAndInstanceNew(
	preset *ControllerApplicationPreset,
	dnet_controller *dnet.Controller,
) (
	*ControllerApplicationPresetAndInstance,
	error,
) {
	ret := new(ControllerApplicationPresetAndInstance)
	ret.ControllerApplicationPreset = *preset

	for _, i := range dnet_controller.ApplicationModules {
		if i.Name() == preset.Module() {
			if t, err := i.Instance(); err != nil {
				panic("Application module instantination error")
			} else {
				t.SetConfig(preset.Config())
				ret.ApplicationInstance = t
			}
			break
		}
	}

	if ret.ApplicationInstance == nil {
		return nil, errors.New(
			fmt.Sprintf(
				"module '%s' instantination requested but it"+
					" not found in BUILTIN_Application_MODULES",
				preset.Module(),
			),
		)
	}

	return ret, nil
}

type ControllerApplicationPresetExt struct {
	a *ControllerApplicationPreset
}

func ControllerApplicationPresetExtNew(
	orig *ControllerApplicationPreset,
) *ControllerApplicationPresetExt {
	ret := new(ControllerApplicationPresetExt)
	ret.a = orig
	return ret
}

func (self *ControllerApplicationPresetExt) Name() string {
	return self.a.Name()
}

func (self *ControllerApplicationPresetExt) Module() string {
	return self.a.Module()
}

func (self *ControllerApplicationPresetExt) Enabled() bool {
	return self.a.Enabled()
}

func (self *ControllerApplicationPresetExt) Config() string {
	return self.a.Config()
}

type ControllerApplicationPreset struct {
	name    string
	module  string
	enabled bool
	config  string
}

/*
	NOTE: We placing into this structure (and we know what module instance may
				be duplicating) name, module, enabled and config values, as module
				(potentially imported as plugin) should not have right to change those
				settings.
*/
func ControllerApplicationPresetNew(
	name string,
	module string,
	enabled bool,
	config string,
) *ControllerApplicationPreset {

	if !dnet.IsValidPresetName(name) {
		panic("invalid name for ControllerApplicationPreset")
	}

	if !dnet.IsValidModuleName(module) {
		panic("invalid module name for ControllerApplicationPreset")
	}

	ret := new(ControllerApplicationPreset)
	ret.name = name
	ret.module = module
	ret.enabled = enabled
	ret.config = config
	return ret
}

func (self *ControllerApplicationPreset) Name() string {
	return self.name
}

func (self *ControllerApplicationPreset) Module() string {
	return self.module
}

func (self *ControllerApplicationPreset) SetModule(value string) {
	self.module = value
	return
}

func (self *ControllerApplicationPreset) Enabled() bool {
	return self.enabled
}

func (self *ControllerApplicationPreset) SetEnabled(value bool) {
	self.enabled = value
	return
}

func (self *ControllerApplicationPreset) Config() string {
	return self.config
}

func (self *ControllerApplicationPreset) SetConfig(value string) {
	self.config = value
	return
}
