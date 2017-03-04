package main

const CONFIG_DIR = "~/.config/DNet"

type Controller struct {
	db_file  string
	password string
	opened   bool

	window_main *UIWindowMain
}

func NewController(db_file string, password string) *Controller {
	ret := new(Controller)

	ret.window_main = UIWindowMainNew(ret)

	return ret
}

func (self *Controller) IsOpened() bool {
	return self.opened
}

func (self *Controller) ShowMainWindow() {
	self.window_main.Show()
	return
}
