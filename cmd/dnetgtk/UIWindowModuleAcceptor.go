package main

import (
	"fmt"

	//"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"github.com/AnimusPEXUS/dnet/common_types"
)

type UIWindowModuleAcceptor struct {
	builtin  bool
	name     *common_types.ModuleName     // for builtin
	checksum *common_types.ModuleChecksum // for external

	last_search_result *ModuleSearcherSearchResult

	main_window *UIWindowMain

	window *gtk.Window

	label_builtin        *gtk.Label
	label_name           *gtk.Label
	label_title          *gtk.Label
	label_description    *gtk.Label
	label_filepath       *gtk.Label
	label_checksum       *gtk.Label
	button_research      *gtk.Button
	button_open_external *gtk.Button
	button_accept_module *gtk.Button
	button_cancel        *gtk.Button
}

func UIWindowModuleAcceptorNew(
	main_window *UIWindowMain,
	builtin bool,
	name *common_types.ModuleName, // for builtin
	checksum *common_types.ModuleChecksum, // for external
) (*UIWindowModuleAcceptor, error) {

	ret := new(UIWindowModuleAcceptor)

	ret.builtin = builtin
	ret.name = name
	ret.checksum = checksum

	ret.main_window = main_window

	builder, err := gtk.BuilderNew()
	if err != nil {
		panic(err.Error())
	}

	data, err := uiModuleAcceptorGladeBytes()
	if err != nil {
		panic(err.Error())
	}

	err = builder.AddFromString(string(data))
	if err != nil {
		panic(err.Error())
	}

	{
		t0, _ := builder.GetObject("window")
		t1, _ := t0.(*gtk.Window)
		ret.window = t1
	}

	{
		t0, _ := builder.GetObject("label_builtin")
		t1, _ := t0.(*gtk.Label)
		ret.label_builtin = t1
	}

	{
		t0, _ := builder.GetObject("label_name")
		t1, _ := t0.(*gtk.Label)
		ret.label_name = t1
	}

	{
		t0, _ := builder.GetObject("label_title")
		t1, _ := t0.(*gtk.Label)
		ret.label_title = t1
	}

	{
		t0, _ := builder.GetObject("label_description")
		t1, _ := t0.(*gtk.Label)
		ret.label_description = t1
	}

	{
		t0, _ := builder.GetObject("label_filepath")
		t1, _ := t0.(*gtk.Label)
		ret.label_filepath = t1
	}

	{
		t0, _ := builder.GetObject("label_checksum")
		t1, _ := t0.(*gtk.Label)
		ret.label_checksum = t1
	}

	{
		t0, _ := builder.GetObject("button_research")
		t1, _ := t0.(*gtk.Button)
		ret.button_research = t1
	}

	{
		t0, _ := builder.GetObject("button_open_external")
		t1, _ := t0.(*gtk.Button)
		ret.button_open_external = t1
	}

	{
		t0, _ := builder.GetObject("button_accept_module")
		t1, _ := t0.(*gtk.Button)
		ret.button_accept_module = t1
	}

	{
		t0, _ := builder.GetObject("button_cancel")
		t1, _ := t0.(*gtk.Button)
		ret.button_cancel = t1
	}

	ret.button_cancel.Connect(
		"clicked",
		func() {
			ret.window.Close()
		},
	)

	ret.button_research.Connect(
		"clicked",
		func() {
			ret.ReSearchModuleFile()
		},
	)

	ret.button_accept_module.Connect(
		"clicked",
		func() {
			err := ret.main_window.controller.application_controller.AcceptModule(
				ret.builtin,
				ret.name,
				ret.checksum,
				true,
			)
			ret.window.Close()
			if err != nil {
				d := gtk.MessageDialogNew(
					ret.window,
					0,
					gtk.MESSAGE_ERROR,
					gtk.BUTTONS_OK,
					"Error accepting module: "+err.Error(),
				)
				d.Run()
				d.Close()
			}
		},
	)

	{

		if !builtin {
			ret.ReSearchModuleFile()
		}

		ret.RefreshStates()

	}

	return ret, nil

}

func (self *UIWindowModuleAcceptor) Show() {
	self.window.ShowAll()
	return
}

func (self *UIWindowModuleAcceptor) ReSearchModuleFile() {
	search_res, err := self.main_window.controller.application_controller.
		module_searcher.SearchMod(
		self.builtin,
		self.name,
		self.checksum,
	)

	if err != nil {
		self.last_search_result = nil
		// TODO: display search error
	} else {
		self.last_search_result = search_res
	}

	self.RefreshStates()
}

func (self *UIWindowModuleAcceptor) RefreshLabels() {
	self.label_builtin.SetText(fmt.Sprintf("%v", self.builtin))

	if self.builtin {
		self.label_filepath.SetText("N/A")
	} else {
		if self.last_search_result == nil {
			self.label_filepath.SetText("not found")
		} else {
			self.label_filepath.SetText(self.last_search_result.Path())
		}
	}
	if self.builtin {
		self.label_checksum.SetText("N/A")
	} else {
		self.label_checksum.SetText(self.checksum.String())
	}
	return
}

func (self *UIWindowModuleAcceptor) RefreshEnableds() {
}

func (self *UIWindowModuleAcceptor) RefreshStates() {
	self.RefreshLabels()
	self.RefreshEnableds()
}
