package main

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type UIWindowLogin struct {
	root *gtk.Window

	button_open           *gtk.Button
	button_browse_storage *gtk.Button

	entry_name     *gtk.Entry
	entry_password *gtk.Entry
}

func UIWindowLoginNew(preset_entry_name string) *UIWindowLogin {

	ret := new(UIWindowLogin)

	builder, err := gtk.BuilderNew()
	if err != nil {
		panic(err.Error())
	}

	data, err := uiLoginGladeBytes()
	if err != nil {
		panic(err.Error())
	}

	err = builder.AddFromString(string(data))
	if err != nil {
		panic(err.Error())
	}

	{
		t0, _ := builder.GetObject("root")
		t1, _ := t0.(*gtk.Window)
		ret.root = t1
	}

	{
		t0, _ := builder.GetObject("button_open")
		t1, _ := t0.(*gtk.Button)
		ret.button_open = t1
	}

	{
		t0, _ := builder.GetObject("button_browse_storage")
		t1, _ := t0.(*gtk.Button)
		ret.button_browse_storage = t1
	}

	{
		t0, _ := builder.GetObject("entry_name")
		t1, _ := t0.(*gtk.Entry)
		ret.entry_name = t1
	}

	{
		t0, _ := builder.GetObject("entry_password")
		t1, _ := t0.(*gtk.Entry)
		ret.entry_password = t1
	}

	ret.entry_name.SetText(preset_entry_name)

	ret.button_open.Connect(
		"clicked",
		func(
			button *gtk.Button,
			win *UIWindowLogin,
		) {
			name, err := win.entry_name.GetText()
			if err != nil {
				panic(err.Error)
			}
			password, err := win.entry_password.GetText()
			if err != nil {
				panic(err.Error)
			}
			controller, err := NewController(
				name,
				password,
			)
			if err != nil {
				glib.IdleAdd(
					func() {
						d := gtk.MessageDialogNew(
							ret.root,
							0,
							gtk.MESSAGE_ERROR,
							gtk.BUTTONS_OK,
							"Error creating controller: "+err.Error(),
						)
						d.Run()
						d.Destroy()
					},
				)
				return
			}
			controller.ShowMainWindow()
			win.root.Destroy()
		},
		ret,
	)

	return ret
}

func (self *UIWindowLogin) Show() {
	self.root.ShowAll()
}
