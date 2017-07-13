package builtin_ownkeypair

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	// "github.com/AnimusPEXUS/dnet"
	// "github.com/AnimusPEXUS/dnet/common_types"
	"github.com/AnimusPEXUS/dnet/cmd/dnetgtk/common_widgets/key_cert_editor"
)

type UIWindow struct {
	inst *Instance

	window                       *gtk.Window
	button_generate_own_key_pair *gtk.Button
	button_save_own_key_pair     *gtk.Button
	button_load_own_key_pair     *gtk.Button
	box_keys                     *gtk.Box
	key_editor_own               *key_cert_editor.UIKeyCertEditor
}

func UIWindowNew(inst *Instance) (*UIWindow, error) {

	ret := new(UIWindow)

	ret.inst = inst

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
		t0, _ := builder.GetObject("window")
		t1, _ := t0.(*gtk.Window)
		ret.window = t1
	}

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
		ret.key_editor_own =
			key_cert_editor.UIKeyCertEditorNew(ret.window, "private")
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
								ret.window,
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
							ret.window,
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
			err = ret.inst.db.SetOwnPrivKey(txt)
			if err != nil {
				glib.IdleAdd(
					func() {
						d := gtk.MessageDialogNew(
							ret.window,
							0,
							gtk.MESSAGE_ERROR,
							gtk.BUTTONS_OK,
							"Key saving error: "+err.Error(),
						)
						d.Run()
						d.Destroy()
					},
				)
			} else {
				glib.IdleAdd(
					func() {
						d := gtk.MessageDialogNew(
							ret.window,
							0,
							gtk.MESSAGE_INFO,
							gtk.BUTTONS_OK,
							"Saved ",
						)
						d.Run()
						d.Destroy()
					},
				)
			}
		},
	)

	ret.button_load_own_key_pair.Connect(
		"clicked",
		func() {
			txt, err := ret.inst.db.GetOwnPrivKey()
			if err != nil {
				glib.IdleAdd(
					func() {
						d := gtk.MessageDialogNew(
							ret.window,
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

func (self *UIWindow) Show() error {
	self.window.ShowAll()
	return nil
}
