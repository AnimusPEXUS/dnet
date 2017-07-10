package main

import (
	"errors"

	"github.com/AnimusPEXUS/dnet/common_types"
)

// Wrap is for DNet's safety. So, for instance, App could not change .Name()
// output at runtime

//
type ControllerApplicationWrap struct {
	controller *Controller
	Name       *common_types.ModuleName
	Module     common_types.ApplicationModule
	DBStatus   *ApplicationStatus
	Instance   common_types.ApplicationModuleInstance
}

func ControllerApplicationWrapNew(
	controller *Controller,
	builtin bool,
	name *common_types.ModuleName,
	checksum *common_types.ModuleChecksum,
) (
	*ControllerApplicationWrap,
	error,
) {
	// NOTE: this structure should be considered for controller internal use
	//       only, so it was ok to move some of controller's controlling
	//       functionality here

	search_res, err :=
		controller.ModSearcher.SearchMod(builtin, name, checksum)
	if err != nil {
		return nil, errors.New("can't find module: " + err.Error())
	}

	ret := new(ControllerApplicationWrap)
	ret.controller = controller

	mod, err := search_res.Mod()
	if err != nil {
		return nil, errors.New("Error getting module structure: " + err.Error())
	}

	ret.Module = mod
	ret.Name = ret.Module.Name()

	dbstat, err := controller.DB.GetApplicationStatus(name.Value())
	if err != nil {
		return nil, errors.New("can't get application status from storage")
	}

	dbstat.Builtin = builtin

	ret.DBStatus = dbstat

	if builtin {
		dbstat.Checksum = ""
	} else {
		dbstat.Checksum = search_res.Checksum()
	}

	dbstat.Name = name.Value()

	// TODO: here should be added security checks, to ensure we are passing
	//       DB to right module
	db, err := ret.controller.DB.GetAppDB(ret.Name.Value())
	if err != nil {
		return nil, errors.New("Error getting DB connection: " + err.Error())
	}

	cc := &ControllerCommunicatorForApp{
		controller: controller,
		wrap:       ret,
		db:         db.DB,
	}

	inst, err := mod.Instance(cc)
	if err != nil {
		return nil, errors.New("Error instantinating module: " + err.Error())
	}
	ret.Instance = inst

	return ret, nil

}
