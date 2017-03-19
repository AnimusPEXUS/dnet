package main

/*
import (
	"errors"
)
*/

const CONFIG_DIR = "~/.config/DNet"

type Controller struct {
	//db_file  string
	//password string
	//opened   bool

	DB *DB

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

	return ret, nil
}

/*
func (self *Controller) IsOpened() bool {
	return self.opened
}
*/

func (self *Controller) ShowMainWindow() {
	self.window_main.Show()
	return
}
