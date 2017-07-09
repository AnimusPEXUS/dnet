package main

import (
	"github.com/AnimusPEXUS/dnet"

	"github.com/AnimusPEXUS/dnet/common_types"

	"github.com/AnimusPEXUS/dnet/cmd/dnetgtk/applications/builtin_net"
	"github.com/AnimusPEXUS/dnet/cmd/dnetgtk/applications/builtin_net_ip"
	"github.com/AnimusPEXUS/dnet/cmd/dnetgtk/applications/builtin_ownkeypair"
	"github.com/AnimusPEXUS/dnet/cmd/dnetgtk/applications/builtin_owntlscert"
)

//const CONFIG_DIR = "~/.config/DNet"

type Controller struct {
	//db_file  string
	//password string
	//opened   bool

	dnet_controller *dnet.Controller

	db *DB

	module_searcher *ModuleSercher

	window_main *UIWindowMain

	builtin_app_modules []common_types.ApplicationModule

	application_controller *ApplicationController
}

func NewController(username string, key string) (*Controller, error) {

	ret := new(Controller)

	{
		t, err := NewDB(username, key)
		if err != nil {
			return nil, err
		}
		ret.db = t
	}

	ret.builtin_app_modules = make(map[string]common_types.ApplicationModule)
	ret.builtin_app_modules[builtin_ownkeypair] = new(builtin_ownkeypair.Module)
	ret.builtin_app_modules[builtin_owntlscert] = new(builtin_owntlscert.Module)
	ret.builtin_app_modules[builtin_net] = new(builtin_net.Module)
	ret.builtin_app_modules[builtin_net_ip] = new(builtin_net_ip.Module)

	ret.module_searcher = ModuleSercherNew(ret.builtin_app_modules)

	ret.application_controller = NewApplicationController(
		ret.module_searcher,
		ret.db,
	)

	// Next line requires modules to be present already
	ret.application_controller.RestoreInstances()

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
