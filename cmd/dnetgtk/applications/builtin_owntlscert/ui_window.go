package builtin_owntlscert

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"math/big"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	// "github.com/AnimusPEXUS/dnet"
	// "github.com/AnimusPEXUS/dnet/common_types"

	"github.com/AnimusPEXUS/dnet/cmd/dnetgtk/common_widgets/key_cert_editor"
)

type UIWindow struct {
	inst *Instance

	window                          *gtk.Window
	button_generate_own_certificate *gtk.Button
	button_save_own_certificate     *gtk.Button
	button_load_own_certificate     *gtk.Button
	box_certificate                 *gtk.Box
	cert_editor_own                 *key_cert_editor.UIKeyCertEditor
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
		t0, _ := builder.GetObject("box_certificate")
		t1, _ := t0.(*gtk.Box)
		ret.box_certificate = t1
	}

	{
		ret.cert_editor_own =
			key_cert_editor.UIKeyCertEditorNew(ret.window, "certificate")
		r := ret.cert_editor_own.GetRoot()
		ret.box_certificate.Add(r)
		r.SetHExpand(true)
	}

	ret.button_generate_own_certificate.Connect(
		"clicked",
		func() {
			go func() {

				var priv_key *rsa.PrivateKey
				var pub_key *rsa.PublicKey

				rpc_client, err := ret.inst.com.GetInnodeRPC("builtin_ownkeypair")

				if err != nil {
					glib.IdleAdd(
						func() {
							d := gtk.MessageDialogNew(
								ret.window,
								0,
								gtk.MESSAGE_ERROR,
								gtk.BUTTONS_OK,
								"Coldn't access module `builtin_ownkeypair':"+err.Error(),
							)
							d.Run()
							d.Destroy()
						},
					)
					return
				}

				{

					var priv_pem string
					err := rpc_client.Call("RPC.GetOwnPrivKey", false, &priv_pem)

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

					block, _ := pem.Decode([]byte(priv_pem))
					if block == nil {
						glib.IdleAdd(
							func() {
								d := gtk.MessageDialogNew(
									ret.window,
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
									ret.window,
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
					ExtKeyUsage: []x509.ExtKeyUsage{
						x509.ExtKeyUsageServerAuth,
					},
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

	ret.button_save_own_certificate.Connect(
		"clicked",
		func() {
			txt, err := ret.cert_editor_own.GetText()
			if err != nil {
				glib.IdleAdd(
					func() {
						d := gtk.MessageDialogNew(
							ret.window,
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
			ret.inst.db.SetOwnTLSCertificate(txt)
		},
	)

	ret.button_load_own_certificate.Connect(
		"clicked",
		func() {
			txt, err := ret.inst.db.GetOwnTLSCertificate()
			if err != nil {
				glib.IdleAdd(
					func() {
						d := gtk.MessageDialogNew(
							ret.window,
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

	return ret, nil
}

func (self *UIWindow) Show() error {
	self.window.ShowAll()
	return nil
}
