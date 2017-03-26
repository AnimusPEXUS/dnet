package main

import (
	"github.com/gotk3/gotk3/gtk"
)

type UINetworkModuleConfigEditor struct {
	root              *gtk.Window
	transient         *gtk.Window
	tw_config_text    *gtk.TextView
	entry_preset_name *gtk.Entry
	label_editor_mode *gtk.Label
	label_module_name *gtk.Label
	button_ok         *gtk.Button
	button_cancel     *gtk.Button
}

func UINetworkModuleConfigEditorNew(
	transient *gtk.Window,
	mode_edit bool,
	preset_name string,
	module_name string,
) *UINetworkModuleConfigEditor {

	ret := new(UINetworkModuleConfigEditor)

	ret.transient = transient

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
		ret.root = t1
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

	return ret
}
