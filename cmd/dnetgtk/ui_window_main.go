package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type UIWindowMain struct {
	controller *Controller

	root *gtk.Window

	/*
		button_dnet     *gtk.Button
		button_storage  *gtk.Button
		button_keys     *gtk.Button
		button_certs    *gtk.Button
		button_networks *gtk.Button
		button_services *gtk.Button
	*/
	button_online                *gtk.Button
	button_home_sep              *gtk.Separator
	button_home                  *gtk.Button
	button_generate_own_key_pair *gtk.Button
	mi_storage                   *gtk.MenuItem
	mi_keys                      *gtk.MenuItem
	mi_about                     *gtk.MenuItem
	box_keys                     *gtk.Box
	notebook_main                *gtk.Notebook

	key_editor_own_private *UIKeyEditor
	//key_editor_own_public  *UIKeyEditor
}

func UIWindowMainNew(controller *Controller) *UIWindowMain {

	ret := new(UIWindowMain)

	ret.controller = controller

	builder, err := gtk.BuilderNew()
	if err != nil {
		panic(err.Error())
	}

	data, err := uiMainGladeBytes()
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
		t0, _ := builder.GetObject("button_home")
		t1, _ := t0.(*gtk.Button)
		ret.button_home = t1
	}

	{
		t0, _ := builder.GetObject("button_home_sep")
		t1, _ := t0.(*gtk.Separator)
		ret.button_home_sep = t1
	}

	{
		t0, _ := builder.GetObject("button_generate_own_key_pair")
		t1, _ := t0.(*gtk.Button)
		ret.button_generate_own_key_pair = t1
	}

	{
		t0, _ := builder.GetObject("notebook_main")
		t1, _ := t0.(*gtk.Notebook)
		ret.notebook_main = t1
	}

	{
		t0, _ := builder.GetObject("mi_about")
		t1, _ := t0.(*gtk.MenuItem)
		ret.mi_about = t1
	}

	{
		t0, _ := builder.GetObject("mi_storage")
		t1, _ := t0.(*gtk.MenuItem)
		ret.mi_storage = t1
	}

	{
		t0, _ := builder.GetObject("mi_keys")
		t1, _ := t0.(*gtk.MenuItem)
		ret.mi_keys = t1
	}

	{
		t0, _ := builder.GetObject("box_keys")
		t1, _ := t0.(*gtk.Box)
		ret.box_keys = t1
	}

	ret.button_generate_own_key_pair.Connect(
		"clicked",
		func() {
			go ret.onButtonGenerateOwnKeyPair()
		},
	)

	ret.notebook_main.Connect(
		"switch-page",
		func(notebook *gtk.Notebook,
			page *gtk.Widget,
			page_num uint,
		) {
			show := page_num != 7

			if show {
				ret.button_home.Show()
				ret.button_home_sep.Show()
			} else {
				ret.button_home.Hide()
				ret.button_home_sep.Hide()
			}
		},
	)

	ret.mi_about.Connect(
		"activate",
		func() {
			ret.notebook_main.SetCurrentPage(0)
		},
	)

	ret.mi_storage.Connect(
		"activate",
		func() {
			ret.notebook_main.SetCurrentPage(1)
		},
	)

	ret.mi_keys.Connect(
		"activate",
		func() {
			ret.notebook_main.SetCurrentPage(2)
		},
	)

	ret.button_home.Connect(
		"clicked",
		func() {
			ret.notebook_main.SetCurrentPage(7)
		},
	)

	{
		ret.key_editor_own_private = UIKeyEditorNew(ret.root, false)
		r := ret.key_editor_own_private.GetRoot()
		ret.box_keys.Add(r)
		r.SetHExpand(true)
	}
	/*
		{
				ret.key_editor_own_public = UIKeyEditorNew(ret.root, true)
				r := ret.key_editor_own_public.GetRoot()
				ret.box_keys.Add(r)
				r.SetHExpand(true)
		}
	*/

	return ret
}

func (self *UIWindowMain) Show() {
	self.root.ShowAll()
	return
}

func (self *UIWindowMain) onButtonGenerateOwnKeyPair() {
	key, err := rsa.GenerateKey(
		rand.Reader,
		1024,
	)
	if err != nil {
		panic("error")
	}
	// TODO: do better error handler
	marshaled := x509.MarshalPKCS1PrivateKey(key)
	der := pem.EncodeToMemory(
		&pem.Block{
			Bytes: marshaled,
			Type:  "RSA PRIVATE KEY",
		},
	)

	glib.IdleAdd(self.key_editor_own_private.SetText, string(der))

	/*
		marshaled, err = x509.MarshalPKIXPublicKey(key.Public())
		if err != nil {
			// TODO: do better error handler
			panic("error")
		}
		der = pem.EncodeToMemory(
			&pem.Block{
				Bytes: marshaled,
				Type:  "PUBLIC KEY",
			},
		)

		glib.IdleAdd(self.key_editor_own_public.SetText, string(der))
	*/
}
