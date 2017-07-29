package main

import (
	"errors"

	"github.com/AnimusPEXUS/dnet"
	"github.com/AnimusPEXUS/dnet/common_types"

	"github.com/AnimusPEXUS/gologger"

	"github.com/AnimusPEXUS/dnet/cmd/dnetgtk/applications/builtin_address_tracker"
	"github.com/AnimusPEXUS/dnet/cmd/dnetgtk/applications/builtin_net_ip"
	"github.com/AnimusPEXUS/dnet/cmd/dnetgtk/applications/builtin_ownkeypair"
	"github.com/AnimusPEXUS/dnet/cmd/dnetgtk/applications/builtin_owntlscert"
)

//const CONFIG_DIR = "~/.config/DNet"

type Controller struct {
	//db_file  string
	//password string
	//opened   bool
	// *worker.Worker

	dnet_controller *dnet.Controller

	db *DB

	module_searcher *ModuleSearcher

	window_main *UIWindowMain

	//builtin_modules

	application_controller *ApplicationController

	logger *gologger.Logger
}

func NewController(username string, key string) (*Controller, error) {

	ret := new(Controller)
	ret.logger = gologger.New()

	{
		t, err := NewDB(username, key)
		if err != nil {
			return nil, err
		}
		ret.db = t
	}

	builtin_modules := make(common_types.ApplicationModuleMap)

	builtin_modules["builtin_ownkeypair"] = new(builtin_ownkeypair.Module)
	builtin_modules["builtin_owntlscert"] = new(builtin_owntlscert.Module)
	builtin_modules["builtin_address_tracker"] =
		new(builtin_address_tracker.Module)
	// builtin_modules["builtin_ownsshcert"] = new(builtin_ownsshcert.Module)
	//builtin_modules["builtin_net"] = new(builtin_net.Module)
	builtin_modules["builtin_net_ip"] = new(builtin_net_ip.Module)

	ret.module_searcher = ModuleSearcherNew(builtin_modules)

	if ac, err := NewApplicationController(
		ret,
		ret.module_searcher,
		ret.db,
	); err != nil {
		return nil,
			errors.New("could not create new ApplicationController " + err.Error())
	} else {
		ret.application_controller = ac
	}

	// go func() {
	// }()
	// Next line requires modules to be present already
	ret.application_controller.Load()

	// NOTE: commented because of GTK+3 memory work problems and crashes
	ret.application_controller.Start()

	if d, err := dnet.NewController(
		ret.application_controller,
		ret.logger,
	); err != nil {
		return nil,
			errors.New("could not create new DNet Controller " + err.Error())
	} else {
		ret.dnet_controller = d
	}

	ret.window_main = UIWindowMainNew(ret)

	return ret, nil
}

func (self *Controller) ShowMainWindow() {
	self.window_main.Show()
	return
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
