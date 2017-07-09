package key_cert_editor

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

	"github.com/AnimusPEXUS/dnet/cmd/dnetgtk/common_widgets/text_viewer"
)

type UIKeyCertEditor struct {
	mode string
	edit bool

	win *gtk.Window

	root *gtk.Box

	label_type  *gtk.Label
	label_error *gtk.Label

	nb      *gtk.Notebook
	nb_full *gtk.Notebook

	tw_full *gtk.TextView
	tw_part *gtk.TextView

	cmi_show                     *gtk.CheckMenuItem
	cmi_copy                     *gtk.CheckMenuItem
	smi_toggles_separator        *gtk.SeparatorMenuItem
	mi_load                      *gtk.MenuItem
	mi_save                      *gtk.MenuItem
	mi_certtool_key_info         *gtk.MenuItem
	mi_certtool_pubkey_info      *gtk.MenuItem
	mi_certtool_certificate_info *gtk.MenuItem
}

func UIKeyCertEditorNew(
	transient_window *gtk.Window,
	mode string,
) *UIKeyCertEditor {
	ret := new(UIKeyCertEditor)

	switch mode {
	case "private":
	case "public":
	case "certificate":
	default:
		panic("invalid `mode' parameter value")
	}

	ret.mode = mode

	ret.win = transient_window

	builder, err := gtk.BuilderNew()
	if err != nil {
		panic(err.Error())
	}

	data, err := uiKeyCertEditorGladeBytes()
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
		t0, _ := builder.GetObject("nb_full")
		t1, _ := t0.(*gtk.Notebook)
		ret.nb_full = t1
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
		t0, _ := builder.GetObject("label_error")
		t1, _ := t0.(*gtk.Label)
		ret.label_error = t1
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

	{
		t0, _ := builder.GetObject("mi_certtool_certificate_info")
		t1, _ := t0.(*gtk.MenuItem)
		ret.mi_certtool_certificate_info = t1
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
					strings.Title(ret.mode),
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
				b, err := ret.tw_full.GetBuffer()
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
					strings.Title(ret.mode),
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
				b, err := ret.tw_full.GetBuffer()
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

	ret.mi_certtool_certificate_info.Connect(
		"activate",
		func() {
			go ret.onCertToolCertificateInfo()
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

	switch ret.mode {
	case "private":
		{
			ret.label_type.SetText("Private Key")

			ret.cmi_show.SetActive(false)
			ret.cmi_copy.SetActive(false)

			ret.cmi_show.Show()
			ret.cmi_copy.Show()
			ret.smi_toggles_separator.Show()

			ret.mi_certtool_key_info.Show()
			ret.mi_certtool_pubkey_info.Show()
			ret.mi_certtool_certificate_info.Hide()

			ret.nb.SetShowTabs(true)
			ret.nb.SetCurrentPage(1)

			ret.nb_full.SetCurrentPage(1)
		}
	case "public":
		{
			ret.label_type.SetText("Public Key")

			ret.cmi_show.SetActive(true)
			ret.cmi_copy.SetActive(true)

			ret.cmi_show.Hide()
			ret.cmi_copy.Hide()
			ret.smi_toggles_separator.Hide()

			ret.mi_certtool_key_info.Hide()
			ret.mi_certtool_pubkey_info.Show()
			ret.mi_certtool_certificate_info.Hide()

			ret.nb.SetShowTabs(false)
			ret.nb.SetCurrentPage(0)

			ret.nb_full.SetCurrentPage(0)
		}
	case "certificate":
		{
			ret.label_type.SetText("Certificate")

			ret.cmi_show.SetActive(true)
			ret.cmi_copy.SetActive(true)

			ret.cmi_show.Hide()
			ret.cmi_copy.Hide()
			ret.smi_toggles_separator.Hide()

			ret.mi_certtool_key_info.Hide()
			ret.mi_certtool_pubkey_info.Hide()
			ret.mi_certtool_certificate_info.Show()

			ret.nb.SetShowTabs(false)
			ret.nb.SetCurrentPage(0)

			ret.nb_full.SetCurrentPage(0)
		}
	}

	ret.CheckStates()
	ret.CheckDataAndIndicate()

	return ret
}

func (self *UIKeyCertEditor) CheckStates() {

	self.mi_load.SetSensitive(self.edit)

	if self.mode == "private" {
		if self.cmi_copy.GetActive() == false {
			self.cmi_show.SetActive(false)
		}
		self.cmi_show.SetSensitive(self.cmi_copy.GetActive())
	}

	if self.mode == "public" || self.mode == "certificate" ||
		(self.mode == "private" && self.cmi_show.GetActive()) {
		self.nb_full.SetCurrentPage(0)
	} else {
		self.nb_full.SetCurrentPage(1)
	}

	{
		t := (self.mode == "public" || self.mode == "certificate" ||
			(self.mode == "private" && self.cmi_copy.GetActive()))
		self.mi_save.SetSensitive(t)
		self.mi_certtool_key_info.SetSensitive(t)
	}

}

func (self *UIKeyCertEditor) CheckData() error {

	var txt string

	{
		b, _ := self.tw_full.GetBuffer()
		txt, _ = b.GetText(
			b.GetStartIter(),
			b.GetEndIter(),
			false,
		)
	}

	block, _ := pem.Decode([]byte(txt))
	if block == nil {
		return errors.New("Can't decode provided PEM to DER")
	}

	switch self.mode {
	case "private":
		{
			if block.Type != "RSA PRIVATE KEY" {
				return errors.New("Not a private key supplied")
			}
			_, err := x509.ParsePKCS1PrivateKey(block.Bytes)
			if err != nil {
				return errors.New(
					"Can't parse provided text as private key: " + err.Error(),
				)
			}
		}
	case "public":
		{
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
	case "certificate":
		{
			if block.Type != "CERTIFICATE" {
				return errors.New("Not a certificate supplied")
			}
			_, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				return errors.New(
					"Can't parse provided text as certificate: " + err.Error(),
				)
			}
		}
	}

	return nil
}

func (self *UIKeyCertEditor) CheckDataAndIndicate() error {
	err := self.CheckData()
	if err != nil {
		self.label_error.SetText(fmt.Sprintf("(Error: %s)", err.Error()))
	} else {
		self.label_error.SetText("")
	}
	return err
}

func (self *UIKeyCertEditor) GetRoot() *gtk.Box {
	return self.root
}

func (self *UIKeyCertEditor) StartEdit() {
	self.edit = true
	self.CheckStates()
}

func (self *UIKeyCertEditor) StopEdit() {
	self.edit = false
	self.CheckStates()
}

func (self *UIKeyCertEditor) SetText(txt string) {
	b, err := self.tw_full.GetBuffer()
	if err != nil {
		panic("error")
	}
	b.SetText(txt)

	check_key_res := self.CheckDataAndIndicate()

	if check_key_res != nil {
		{
			b, _ := self.tw_part.GetBuffer()
			b.SetText("")
		}
	}

	if check_key_res == nil && self.mode == "private" {
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
				goto retu
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
				goto retu
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
				goto retu
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
				goto retu
			}
			block = &pem.Block{Type: "PUBLIC KEY", Bytes: pub_key}
			txt = string(pem.EncodeToMemory(block))
		}

		glib.IdleAdd(
			func() {
				b, _ := self.tw_part.GetBuffer()
				b.SetText(txt)
			},
		)

	retu:
	}

}

func (self *UIKeyCertEditor) GetText() (string, error) {
	b, err := self.tw_full.GetBuffer()
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

func (self *UIKeyCertEditor) onCertToolKeyInfo() {
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
			vww := text_viewer.UITextViewerNew(
				"certtool --key-info Result",
				output,
				self.win,
			)
			vww.Show()
		},
	)
}

func (self *UIKeyCertEditor) onCertToolPubKeyInfo() {
	txt, err := self.GetText()
	if err != nil {
		panic(err.Error())
	}
	if self.mode == "private" {
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
			if self.mode == "private" {
				title += " (extracted from private key)"
			}
			vww := text_viewer.UITextViewerNew(
				title,
				output,
				self.win,
			)
			vww.Show()
		},
	)
}

func (self *UIKeyCertEditor) onCertToolCertificateInfo() {
	txt, err := self.GetText()
	if err != nil {
		panic(err.Error())
	}
	proc := exec.Command("certtool", "--certificate-info")

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
			vww := text_viewer.UITextViewerNew(
				"certtool --certificate-info Result",
				output,
				self.win,
			)
			vww.Show()
		},
	)
}
