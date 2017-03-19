package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"bitbucket.org/AnimusPEXUS/dnet"
)

type UIWindowMain struct {
	controller *Controller

	win *gtk.Window

	/*
		button_dnet     *gtk.Button
		button_storage  *gtk.Button
		button_keys     *gtk.Button
		button_certs    *gtk.Button
		button_networks *gtk.Button
		button_services *gtk.Button
	*/
	button_online                   *gtk.ToggleButton
	button_home_sep                 *gtk.Separator
	button_home                     *gtk.Button
	button_generate_own_key_pair    *gtk.Button
	button_save_own_key_pair        *gtk.Button
	button_load_own_key_pair        *gtk.Button
	button_generate_own_certificate *gtk.Button
	button_save_own_certificate     *gtk.Button
	button_load_own_certificate     *gtk.Button
	button_refresh_network_modules  *gtk.Button
	mi_storage                      *gtk.MenuItem
	mi_keys                         *gtk.MenuItem
	mi_networks                     *gtk.MenuItem
	mi_tls_cert                     *gtk.MenuItem
	mi_about                        *gtk.MenuItem
	box_keys                        *gtk.Box
	box_certificate                 *gtk.Box
	notebook_main                   *gtk.Notebook

	key_editor_own  *UIKeyCertEditor
	cert_editor_own *UIKeyCertEditor
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
		t0, _ := builder.GetObject("button_generate_own_certificate")
		t1, _ := t0.(*gtk.Button)
		ret.button_generate_own_certificate = t1
	}

	{
		t0, _ := builder.GetObject("button_save_own_certificate")
		t1, _ := t0.(*gtk.Button)
		ret.button_save_own_certificate = t1
	}

	{
		t0, _ := builder.GetObject("button_load_own_certificate")
		t1, _ := t0.(*gtk.Button)
		ret.button_load_own_certificate = t1
	}

	{
		t0, _ := builder.GetObject("button_refresh_network_modules")
		t1, _ := t0.(*gtk.Button)
		ret.button_refresh_network_modules = t1
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

	{
		t0, _ := builder.GetObject("box_keys")
		t1, _ := t0.(*gtk.Box)
		ret.box_keys = t1
	}

	{
		t0, _ := builder.GetObject("box_certificate")
		t1, _ := t0.(*gtk.Box)
		ret.box_certificate = t1
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
								ret.win,
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

	ret.button_generate_own_certificate.Connect(
		"clicked",
		func() {
			go func() {

				var priv_key *rsa.PrivateKey
				var pub_key *rsa.PublicKey

				{
					var err error

					priv_pem, err := ret.controller.DB.GetOwnPrivKey()
					if err != nil {
						glib.IdleAdd(
							func() {
								d := gtk.MessageDialogNew(
									ret.win,
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

					block, _ := pem.Decode([]byte(priv_pem))
					if block == nil {
						glib.IdleAdd(
							func() {
								d := gtk.MessageDialogNew(
									ret.win,
									0,
									gtk.MESSAGE_ERROR,
									gtk.BUTTONS_OK,
									"error decoding priv key PEM block",
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
									ret.win,
									0,
									gtk.MESSAGE_ERROR,
									gtk.BUTTONS_OK,
									"error: not a private key",
								)
								d.Run()
								d.Destroy()
							},
						)
						return
					}
					priv_key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
					if err != nil {
						glib.IdleAdd(
							func() {
								d := gtk.MessageDialogNew(
									ret.win,
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

					if pub_key_t, ok := priv_key.Public().(*rsa.PublicKey); !ok {
						panic("can't assert")
					} else {
						pub_key = pub_key_t
					}
				}

				cert_struct := &x509.Certificate{
					// NotBefore:             1,
					// NotAfter:              1,
					SerialNumber: big.NewInt(0),
					KeyUsage: (x509.KeyUsageDigitalSignature |
						x509.KeyUsageDataEncipherment |
						x509.KeyUsageCertSign),
					ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
					BasicConstraintsValid: true,
					IsCA:               true,
					MaxPathLen:         0,
					SubjectKeyId:       []byte{1, 2, 3, 4, 5},
					SignatureAlgorithm: x509.SHA512WithRSA,
				}

				cert, err := x509.CreateCertificate(
					rand.Reader,
					cert_struct,
					cert_struct,
					pub_key,
					priv_key,
				)
				if err != nil {
					glib.IdleAdd(
						func() {
							d := gtk.MessageDialogNew(
								ret.win,
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

				cert_pem := pem.EncodeToMemory(
					&pem.Block{
						Type:  "CERTIFICATE",
						Bytes: cert,
					},
				)

				glib.IdleAdd(ret.cert_editor_own.SetText, string(cert_pem))
			}()
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
			ret.notebook_main.SetCurrentPage(7)
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
							ret.win,
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
			ret.controller.DB.SetOwnPrivKey(txt)
		},
	)

	ret.button_load_own_key_pair.Connect(
		"clicked",
		func() {
			txt, err := ret.controller.DB.GetOwnPrivKey()
			if err != nil {
				glib.IdleAdd(
					func() {
						d := gtk.MessageDialogNew(
							ret.win,
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

	ret.button_save_own_certificate.Connect(
		"clicked",
		func() {
			txt, err := ret.cert_editor_own.GetText()
			if err != nil {
				glib.IdleAdd(
					func() {
						d := gtk.MessageDialogNew(
							ret.win,
							0,
							gtk.MESSAGE_ERROR,
							gtk.BUTTONS_OK,
							"Error getting text from certificate editor ui: "+err.Error(),
						)
						d.Run()
						d.Destroy()
					},
				)
				return
			}
			ret.controller.DB.SetOwnTLSCertificate(txt)
		},
	)

	ret.button_load_own_certificate.Connect(
		"clicked",
		func() {
			txt, err := ret.controller.DB.GetOwnTLSCertificate()
			if err != nil {
				glib.IdleAdd(
					func() {
						d := gtk.MessageDialogNew(
							ret.win,
							0,
							gtk.MESSAGE_ERROR,
							gtk.BUTTONS_OK,
							"Error getting certificate from storage: "+err.Error(),
						)
						d.Run()
						d.Destroy()
					},
				)
				return
			}
			ret.cert_editor_own.SetText(txt)
		},
	)

	ret.button_online.Connect(
		"toggled",
		func() {
			fmt.Println("toggled")
		},
	)

	ret.button_refresh_network_modules.Connect(
		"clicked",
		func() { ret.RefteshNetworkModulesList() },
	)

	{
		ret.key_editor_own = UIKeyCertEditorNew(ret.win, "private")
		r := ret.key_editor_own.GetRoot()
		ret.box_keys.Add(r)
		r.SetHExpand(true)
	}

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

func (self *UIWindowMain) RefteshNetworkModulesList() {
	for _, i := range dnet.BUILTIN_NETWORK_MODULES {
		fmt.Println(i.Name())
	}
}
