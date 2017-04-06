package main

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
	//"bitbucket.org/AnimusPEXUS/dnet"
	//"bitbucket.org/AnimusPEXUS/dnet/common_types"
)

type UIWindowMain struct {
	controller *Controller

	*UIWindowMainTabKeys
	*UIWindowMainTabNetworks
	*UIWindowMainTabTLSCertificate

	win *gtk.Window

	/*
		button_dnet     *gtk.Button
		button_storage  *gtk.Button
		button_keys     *gtk.Button
		button_certs    *gtk.Button
		button_networks *gtk.Button
		button_services *gtk.Button
	*/
	button_online   *gtk.ToggleButton
	button_home_sep *gtk.Separator
	button_home     *gtk.Button
	mi_storage      *gtk.MenuItem
	mi_keys         *gtk.MenuItem
	mi_networks     *gtk.MenuItem
	mi_tls_cert     *gtk.MenuItem
	mi_about        *gtk.MenuItem

	notebook_main *gtk.Notebook

	services_presets *gtk.ListStore
	services_modules *gtk.ListStore
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
		if res, err := UIWindowMainTabNetworksNew(
			builder,
			ret,
		); err == nil {
			ret.UIWindowMainTabNetworks = res
		} else {
			panic(err.Error())
		}
	}

	{
		if res, err := UIWindowMainTabKeysNew(
			builder,
			ret,
		); err == nil {
			ret.UIWindowMainTabKeys = res
		} else {
			panic(err.Error())
		}
	}

	{
		if res, err := UIWindowMainTabTLSCertificateNew(
			builder,
			ret,
		); err == nil {
			ret.UIWindowMainTabTLSCertificate = res
		} else {
			panic(err.Error())
		}
	}

	{
		t0, _ := builder.GetObject("root")
		t1, _ := t0.(*gtk.Window)
		ret.win = t1
	}

	{
		t0, _ := builder.GetObject("button_online")
		t1, _ := t0.(*gtk.ToggleButton)
		ret.button_online = t1
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
		t0, _ := builder.GetObject("mi_tls_cert")
		t1, _ := t0.(*gtk.MenuItem)
		ret.mi_tls_cert = t1
	}

	{
		t0, _ := builder.GetObject("mi_networks")
		t1, _ := t0.(*gtk.MenuItem)
		ret.mi_networks = t1
	}

	ret.notebook_main.Connect(
		"switch-page",
		func(notebook *gtk.Notebook,
			page *gtk.Widget,
			page_num uint,
		) {
			show := page_num != 8

			if show {
				ret.notebook_main.SetShowTabs(true)
				ret.button_home.Show()
				ret.button_home_sep.Show()
			} else {
				ret.notebook_main.SetShowTabs(false)
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

	ret.mi_tls_cert.Connect(
		"activate",
		func() {
			ret.notebook_main.SetCurrentPage(3)
		},
	)

	ret.mi_networks.Connect(
		"activate",
		func() {
			ret.notebook_main.SetCurrentPage(6)
		},
	)

	ret.button_home.Connect(
		"clicked",
		func() {
			ret.notebook_main.SetCurrentPage(8)
		},
	)

	ret.button_online.Connect(
		"toggled",
		func() {
			fmt.Println("toggled")
		},
	)

	{
		ret.cert_editor_own = UIKeyCertEditorNew(ret.win, "certificate")
		r := ret.cert_editor_own.GetRoot()
		ret.box_certificate.Add(r)
		r.SetHExpand(true)
	}

	return ret
}

func (self *UIWindowMain) Show() {
	self.win.ShowAll()
	return
}
