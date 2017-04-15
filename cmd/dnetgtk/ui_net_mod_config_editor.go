package main

import (
	"fmt"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"bitbucket.org/AnimusPEXUS/dnet"
)

type UINetworkModuleConfigEditor struct {
	win               *gtk.Window
	transient         *gtk.Window
	tw_config_text    *gtk.TextView
	entry_preset_name *gtk.Entry
	label_editor_mode *gtk.Label
	label_module_name *gtk.Label
	button_ok         *gtk.Button
	button_cancel     *gtk.Button

	mode_edit bool
}

func UINetworkModuleConfigEditorNew(
	controller *Controller,
	transient *gtk.Window,
	mode_edit bool,
	preset_name string,
	module_name string,
	text string,
) *UINetworkModuleConfigEditor {

	ret := new(UINetworkModuleConfigEditor)

	ret.transient = transient
	ret.mode_edit = mode_edit

	builder, err := gtk.BuilderNew()
	if err != nil {
		panic(err.Error())
	}

	data, err := uiNetModConfigEditorGladeBytes()
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
		ret.win = t1
	}

	{
		t0, _ := builder.GetObject("tw_config_text")
		t1, _ := t0.(*gtk.TextView)
		ret.tw_config_text = t1
	}

	{
		t0, _ := builder.GetObject("entry_preset_name")
		t1, _ := t0.(*gtk.Entry)
		ret.entry_preset_name = t1
	}

	{
		t0, _ := builder.GetObject("label_editor_mode")
		t1, _ := t0.(*gtk.Label)
		ret.label_editor_mode = t1
	}

	{
		t0, _ := builder.GetObject("label_module_name")
		t1, _ := t0.(*gtk.Label)
		ret.label_module_name = t1
	}

	{
		t0, _ := builder.GetObject("button_ok")
		t1, _ := t0.(*gtk.Button)
		ret.button_ok = t1
	}

	{
		t0, _ := builder.GetObject("button_cancel")
		t1, _ := t0.(*gtk.Button)
		ret.button_cancel = t1
	}

	{
		t := ""
		if !mode_edit {
			t = fmt.Sprintf(
				`Creating new preset for %s module`,
				module_name,
			)
		} else {
			t = fmt.Sprintf(
				`Editing preset %s of module %s`,
				preset_name,
				module_name,
			)
		}
		ret.win.SetTitle(t)
	}

	{
		b, _ := ret.tw_config_text.GetBuffer()
		b.SetText(text)
	}

	ret.entry_preset_name.SetText(preset_name)
	ret.label_module_name.SetText(module_name)
	if mode_edit {
		ret.label_editor_mode.SetText("Editing")
	} else {
		ret.label_editor_mode.SetText("Creating")
	}

	ret.button_cancel.Connect(
		"clicked",
		func() {
			ret.win.Close()
		},
	)

	ret.button_ok.Connect(
		"clicked",
		func() {

			name, _ := ret.entry_preset_name.GetText()

			module, _ := ret.label_module_name.GetText()

			config := ""

			if !dnet.IsValidPresetName(name) {
				glib.IdleAdd(
					func() {
						d := gtk.MessageDialogNew(
							ret.win,
							0,
							gtk.MESSAGE_ERROR,
							gtk.BUTTONS_OK,
							"Given Name is Invalid for Preset",
						)
						d.Run()
						d.Destroy()
					},
				)
				return
			}

			b, _ := ret.tw_config_text.GetBuffer()
			config, _ = b.GetText(
				b.GetStartIter(),
				b.GetEndIter(),
				false,
			)

			err := controller.AddNetworkPreset(
				name,
				module,
				false,
				config,
			)

			if err != nil {
				glib.IdleAdd(
					func() {
						d := gtk.MessageDialogNew(
							ret.win,
							0,
							gtk.MESSAGE_ERROR,
							gtk.BUTTONS_OK,
							"error: "+err.Error(),
						)
						d.Run()
						d.Destroy()
					},
				)
				return
			} else {
				glib.IdleAdd(
					func() {
						d := gtk.MessageDialogNew(
							ret.win,
							0,
							gtk.MESSAGE_INFO,
							gtk.BUTTONS_OK,
							"Added Ok",
						)
						d.Run()
						d.Destroy()
					},
				)
			}

			ret.win.Close()

		},
	)

	return ret
}

func (self *UINetworkModuleConfigEditor) Show() {
	self.win.ShowAll()
	return
}
