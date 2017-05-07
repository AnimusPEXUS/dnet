package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	// "bitbucket.org/AnimusPEXUS/dnet"
	// "bitbucket.org/AnimusPEXUS/dnet/common_types"
)

type UIWindowMainTabKeys struct {
	main_window *UIWindowMain

	button_generate_own_key_pair    *gtk.Button
	button_save_own_key_pair        *gtk.Button
	button_load_own_key_pair        *gtk.Button
	box_keys                        *gtk.Box
	key_editor_own                  *UIKeyCertEditor
}

func UIWindowMainTabKeysNew(
	builder *gtk.Builder,
	main_window *UIWindowMain,
) (*UIWindowMainTabKeys, error) {

	ret := new(UIWindowMainTabKeys)

	ret.main_window = main_window

	{
		t0, _ := builder.GetObject("button_generate_own_key_pair")
		t1, _ := t0.(*gtk.Button)
		ret.button_generate_own_key_pair = t1
	}

	{
		t0, _ := builder.GetObject("button_save_own_key_pair")
		t1, _ := t0.(*gtk.Button)
		ret.button_save_own_key_pair = t1
	}

	{
		t0, _ := builder.GetObject("button_load_own_key_pair")
		t1, _ := t0.(*gtk.Button)
		ret.button_load_own_key_pair = t1
	}

	{
		t0, _ := builder.GetObject("box_keys")
		t1, _ := t0.(*gtk.Box)
		ret.box_keys = t1
	}

	{
		ret.key_editor_own = UIKeyCertEditorNew(ret.main_window.win, "private")
		r := ret.key_editor_own.GetRoot()
		ret.box_keys.Add(r)
		r.SetHExpand(true)
	}

	ret.button_generate_own_key_pair.Connect(
		"clicked",
		func() {
			go func() {
				key, err := rsa.GenerateKey(
					rand.Reader,
					1024,
				)
				if err != nil {
					glib.IdleAdd(
						func() {
							d := gtk.MessageDialogNew(
								ret.main_window.win,
								0,
								gtk.MESSAGE_ERROR,
								gtk.BUTTONS_OK,
								err.Error(),
							)
							d.Run()
							d.Destroy()
						},
					)
					return
				}

				marshaled := x509.MarshalPKCS1PrivateKey(key)

				der := pem.EncodeToMemory(
					&pem.Block{
						Bytes: marshaled,
						Type:  "RSA PRIVATE KEY",
					},
				)

				glib.IdleAdd(ret.key_editor_own.SetText, string(der))
			}()
		},
	)

	ret.button_save_own_key_pair.Connect(
		"clicked",
		func() {
			txt, err := ret.key_editor_own.GetText()
			if err != nil {
				glib.IdleAdd(
					func() {
						d := gtk.MessageDialogNew(
							ret.main_window.win,
							0,
							gtk.MESSAGE_ERROR,
							gtk.BUTTONS_OK,
							"Error getting text from key editor ui: "+err.Error(),
						)
						d.Run()
						d.Destroy()
					},
				)
				return
			}
			ret.main_window.controller.DB.SetOwnPrivKey(txt)
		},
	)

	ret.button_load_own_key_pair.Connect(
		"clicked",
		func() {
			txt, err := ret.main_window.controller.DB.GetOwnPrivKey()
			if err != nil {
				glib.IdleAdd(
					func() {
						d := gtk.MessageDialogNew(
							ret.main_window.win,
							0,
							gtk.MESSAGE_ERROR,
							gtk.BUTTONS_OK,
							"Error getting key from storage: "+err.Error(),
						)
						d.Run()
						d.Destroy()
					},
				)
				return
			}
			ret.key_editor_own.SetText(txt)
		},
	)

	return ret, nil
}
