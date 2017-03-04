package main

import (
	//"github.com/gotk3/gotk3/glib"
	//	"bytes"
	"fmt"
	//"io"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type UIKeyEditor struct {
	mode_public     bool
	mode_public_str string

	mode_edit bool

	win *gtk.Window

	root *gtk.Box

	label_type  *gtk.Label
	label_error *gtk.Label

	nb *gtk.Notebook

	tw_full *gtk.TextView
	tw_part  *gtk.TextView

	mi_load                 *gtk.MenuItem
	mi_save                 *gtk.MenuItem
	mi_copy                 *gtk.MenuItem
	smi_toggles_separator   *gtk.SeparatorMenuItem
	cmi_show                *gtk.CheckMenuItem
	cmi_copy                *gtk.CheckMenuItem
	mi_certtool_key_info    *gtk.MenuItem
	mi_certtool_pubkey_info *gtk.MenuItem
}

func UIKeyEditorNew(transient_window *gtk.Window, public bool) *UIKeyEditor {
	ret := new(UIKeyEditor)

	ret.mode_edit = false

	ret.win = transient_window
	ret.mode_public = public
	if public {
		ret.mode_public_str = "public"
	} else {
		ret.mode_public_str = "private"
	}

	builder, err := gtk.BuilderNew()
	if err != nil {
		panic(err.Error())
	}

	data, err := uiKeyEditorGladeBytes()
	if err != nil {
		panic(err.Error())
	}

	err = builder.AddFromString(string(data))
	if err != nil {
		panic(err.Error())
	}

	{
		t0, _ := builder.GetObject("root")
		t1, _ := t0.(*gtk.Box)
		ret.root = t1
	}

	{
		t0, _ := builder.GetObject("nb")
		t1, _ := t0.(*gtk.Notebook)
		ret.nb = t1
	}

	{
		t0, _ := builder.GetObject("tw_full")
		t1, _ := t0.(*gtk.TextView)
		ret.tw_full = t1
	}

	{
		t0, _ := builder.GetObject("tw_part")
		t1, _ := t0.(*gtk.TextView)
		ret.tw_part = t1
	}

	{
		t0, _ := builder.GetObject("label_type")
		t1, _ := t0.(*gtk.Label)
		ret.label_type = t1
	}

	{
		t0, _ := builder.GetObject("mi_load")
		t1, _ := t0.(*gtk.MenuItem)
		ret.mi_load = t1
	}

	{
		t0, _ := builder.GetObject("mi_save")
		t1, _ := t0.(*gtk.MenuItem)
		ret.mi_save = t1
	}

	{
		t0, _ := builder.GetObject("mi_copy")
		t1, _ := t0.(*gtk.MenuItem)
		ret.mi_copy = t1
	}

	{
		t0, _ := builder.GetObject("smi_toggles_separator")
		t1, _ := t0.(*gtk.SeparatorMenuItem)
		ret.smi_toggles_separator = t1
	}

	{
		t0, _ := builder.GetObject("cmi_show")
		t1, _ := t0.(*gtk.CheckMenuItem)
		ret.cmi_show = t1
	}

	{
		t0, _ := builder.GetObject("cmi_copy")
		t1, _ := t0.(*gtk.CheckMenuItem)
		ret.cmi_copy = t1
	}

	{
		t0, _ := builder.GetObject("mi_certtool_key_info")
		t1, _ := t0.(*gtk.MenuItem)
		ret.mi_certtool_key_info = t1
	}

	{
		t0, _ := builder.GetObject("mi_certtool_pubkey_info")
		t1, _ := t0.(*gtk.MenuItem)
		ret.mi_certtool_pubkey_info = t1
	}

	ret.mi_save.Connect(
		"activate",
		func() {
			dialog, err := gtk.DialogNew()
			if err != nil {
				panic(err.Error())
			}

			dialog.SetTransientFor(ret.win)

			chooser, err := gtk.FileChooserWidgetNew(gtk.FILE_CHOOSER_ACTION_SAVE)
			if err != nil {
				panic(err.Error())
			}
			dialog.AddButton("Save", gtk.RESPONSE_ACCEPT)
			dialog.AddButton("Cancel", gtk.RESPONSE_CANCEL)

			dialog.SetTitle(
				fmt.Sprintf(
					"Select File Name To Write %s Key",
					strings.Title(ret.mode_public_str),
				),
			)

			box, err := dialog.GetContentArea()
			if err != nil {
				panic(err.Error())
			}
			box.Add(chooser)
			box.ShowAll()

			res := dialog.Run()

			if gtk.ResponseType(res) == gtk.RESPONSE_ACCEPT {
				b, err := ret.tw_private.GetBuffer()
				if err != nil {
					panic("error")
				}
				t, err := b.GetText(
					b.GetStartIter(),
					b.GetEndIter(),
					false,
				)
				if err != nil {
					panic("error")
				}
				ioutil.WriteFile(
					chooser.GetFilename(),
					([]byte)(t),
					0700,
				)
			}
			dialog.Destroy()
		},
	)

	ret.mi_load.Connect(
		"activate",
		func() {
			dialog, err := gtk.DialogNew()
			if err != nil {
				panic(err.Error())
			}

			dialog.SetTransientFor(ret.win)

			chooser, err := gtk.FileChooserWidgetNew(gtk.FILE_CHOOSER_ACTION_OPEN)
			if err != nil {
				panic(err.Error())
			}
			dialog.AddButton("Load", gtk.RESPONSE_ACCEPT)
			dialog.AddButton("Cancel", gtk.RESPONSE_CANCEL)

			dialog.SetTitle(
				fmt.Sprintf(
					"Select File Name To Read %s Key",
					strings.Title(ret.mode_public_str),
				),
			)

			box, err := dialog.GetContentArea()
			if err != nil {
				panic(err.Error())
			}
			box.Add(chooser)
			box.ShowAll()

			res := dialog.Run()

			if gtk.ResponseType(res) == gtk.RESPONSE_ACCEPT {
				b, err := ret.tw.GetBuffer()
				if err != nil {
					panic("error")
				}
				text, err := ioutil.ReadFile(chooser.GetFilename())
				if err != nil {
					panic("error: TODO error message")
				} else {
					b.SetText(string(text))
				}
			}
			dialog.Destroy()
		},
	)

	ret.mi_certtool_key_info.Connect(
		"activate",
		func() {
			go ret.onCertToolKeyInfo()
		},
	)

	ret.mi_certtool_pubkey_info.Connect(
		"activate",
		func() {
			go ret.onCertToolPubKeyInfo()
		},
	)

	ret.cmi_show.Connect(
		"toggled",
		func() {
			ret.CheckStates()
		},
	)

	ret.cmi_copy.Connect(
		"toggled",
		func() {
			ret.CheckStates()
		},
	)

	if public {
		ret.label_type.SetText("Public Key")
		ret.cmi_show.Hide()
		ret.cmi_copy.Hide()
		ret.smi_toggles_separator.Hide()
		ret.mi_certtool_key_info.Hide()
		ret.mi_certtool_pubkey_info.Show()

	} else {
		ret.label_type.SetText("Private Key")
		ret.cmi_show.Show()
		ret.cmi_copy.Show()
		ret.smi_toggles_separator.Show()
		ret.mi_certtool_key_info.Show()
		ret.mi_certtool_pubkey_info.Show()
	}

	ret.CheckStates()
	ret.CheckKeyAndIndicate("")

	return ret
}

func (self *UIKeyEditor) CheckKey() error {

	var txt string

	{
		b, _ := self.tw_private.GetBuffer()
		txt, _ := b.GetText(
			b.GetStartIter(),
			b.GetEndIter(),
			false,
		)
	}

	block, _ := pem.Decode([]byte(txt))
	if block == nil {
		return errors.New("Can't decode provided PEM to DER")
	}

	if !self.mode_public {
		// private
		if block.Type != "RSA PRIVATE KEY" {
			return errors.New("Not a private key supplied")
		}
		_, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return errors.New(
				"Can't parse provided text as private key: " + err.Error(),
			)
		}
	} else {
		// public
		if block.Type != "PUBLIC KEY" {
			return errors.New("Not a public key supplied")
		}
		_, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return errors.New(
				"Can't parse provided text as public key: " + err.Error(),
			)
		}
	}
	return nil
}

func (self *UIKeyEditor) CheckKeyAndIndicate() error {
	err := self.CheckKey()
	if err != nil {
		self.label_error.SetText(fmt.Sprintf("(Error: %s)", err.Error()))
	} else {
		self.label_error.SetText("")
	}
	return err
}

func (self *UIKeyEditor) GetRoot() *gtk.Box {
	return self.root
}

func (self *UIKeyEditor) CheckStates() {

	self.mi_load.SetSensitive(self.mode_edit)

	if !self.mode_public {
		if self.cmi_copy.GetActive() == false {
			self.cmi_show.SetActive(false)
		}
		self.cmi_show.SetSensitive(self.cmi_copy.GetActive())
	}

	if self.mode_public || (!self.mode_public && self.cmi_show.GetActive()) {
		self.nb.SetCurrentPage(0)
	} else {
		self.nb.SetCurrentPage(1)
	}

	{
		t := self.mode_public || (!self.mode_public && self.cmi_copy.GetActive())
		self.mi_save.SetSensitive(t)
		self.mi_copy.SetSensitive(t)
		self.mi_certtool_key_info.SetSensitive(t)
	}

}

func (self *UIKeyEditor) StartEdit() {
	self.mode_edit = true
	self.CheckStates()
}

func (self *UIKeyEditor) StopEdit() {
	self.mode_edit = false
	self.CheckStates()
}

func (self *UIKeyEditor) SetText(txt string) {
	b, err := self.tw.GetBuffer()
	if err != nil {
		panic("error")
	}
	b.SetText(txt)
	self.CheckKeyAndIndicate(txt)
}

func (self *UIKeyEditor) GetText() (string, error) {
	b, err := self.tw.GetBuffer()
	if err != nil {
		return "", err
	}
	t, err := b.GetText(
		b.GetStartIter(),
		b.GetEndIter(),
		false,
	)
	if err != nil {
		return "", err
	}
	return t, nil
}

func (self *UIKeyEditor) onCertToolKeyInfo() {
	txt, err := self.GetText()
	if err != nil {
		panic(err.Error())
	}
	proc := exec.Command("certtool", "--key-info")

	stdin_pipe, err := proc.StdinPipe()
	if err != nil {
		panic(err.Error())
	}

	stdin_pipe.Write([]byte(txt))
	stdin_pipe.Close()

	txt2, err := proc.Output()
	if err != nil {
		glib.IdleAdd(
			func() {
				d := gtk.MessageDialogNew(
					self.win,
					0,
					gtk.MESSAGE_ERROR,
					gtk.BUTTONS_OK,
					"Error executing certtool: "+err.Error(),
				)
				d.Run()
				d.Destroy()
			},
		)
		return
	}

	output := string(txt2)

	glib.IdleAdd(
		func() {
			vww := UITextViewerNew(
				"certtool --key-info Result",
				output,
				nil,
			)
			vww.Show()
		},
	)
}

func (self *UIKeyEditor) onCertToolPubKeyInfo() {
	txt, err := self.GetText()
	if err != nil {
		panic(err.Error())
	}
	if !self.mode_public {
		{
			block, _ := pem.Decode([]byte(txt))
			if block == nil {
				glib.IdleAdd(
					func() {
						d := gtk.MessageDialogNew(
							self.win,
							0,
							gtk.MESSAGE_ERROR,
							gtk.BUTTONS_OK,
							"Can't decode private key PEM",
						)
						d.Run()
						d.Destroy()
					},
				)
				return
			}
			if block.Type != "RSA PRIVATE KEY" {
				glib.IdleAdd(
					func() {
						d := gtk.MessageDialogNew(
							self.win,
							0,
							gtk.MESSAGE_ERROR,
							gtk.BUTTONS_OK,
							"Not a private key supplied",
						)
						d.Run()
						d.Destroy()
					},
				)
				return
			}
			priv_key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
			if err != nil {
				glib.IdleAdd(
					func() {
						d := gtk.MessageDialogNew(
							self.win,
							0,
							gtk.MESSAGE_ERROR,
							gtk.BUTTONS_OK,
							"Can't parse provided text as private key",
						)
						d.Run()
						d.Destroy()
					},
				)
				return
			}
			pub_key, err := x509.MarshalPKIXPublicKey(priv_key.Public())
			if err != nil {
				glib.IdleAdd(
					func() {
						d := gtk.MessageDialogNew(
							self.win,
							0,
							gtk.MESSAGE_ERROR,
							gtk.BUTTONS_OK,
							"Can't render public key PKI from provided private key",
						)
						d.Run()
						d.Destroy()
					},
				)
				return
			}
			block = &pem.Block{Type: "PUBLIC KEY", Bytes: pub_key}
			txt = string(pem.EncodeToMemory(block))
		}

	}
	proc := exec.Command("certtool", "--pubkey-info")

	stdin_pipe, err := proc.StdinPipe()
	if err != nil {
		panic(err.Error())
	}

	stdin_pipe.Write([]byte(txt))
	stdin_pipe.Close()

	txt2, err := proc.Output()
	if err != nil {
		glib.IdleAdd(
			func() {
				d := gtk.MessageDialogNew(
					self.win,
					0,
					gtk.MESSAGE_ERROR,
					gtk.BUTTONS_OK,
					"Error executing certtool: "+err.Error(),
				)
				d.Run()
				d.Destroy()
			},
		)
		return
	}

	output := string(txt2)

	glib.IdleAdd(
		func() {
			title := "certtool --pubkey-info Result"
			if !self.mode_public {
				title += " (extracted from private key)"
			}
			vww := UITextViewerNew(
				title,
				output,
				nil,
			)
			vww.Show()
		},
	)
}
