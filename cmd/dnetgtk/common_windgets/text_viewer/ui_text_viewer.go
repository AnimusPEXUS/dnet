package text_viewer

import (
	"io/ioutil"

	"github.com/gotk3/gotk3/gtk"
)

type UITextViewer struct {
	root      *gtk.Window
	transient *gtk.Window

	tw *gtk.TextView

	button_save  *gtk.Button
	button_close *gtk.Button
}

func UITextViewerNew(
	title string,
	text string,
	transient *gtk.Window,
) *UITextViewer {
	ret := new(UITextViewer)

	ret.transient = transient

	builder, err := gtk.BuilderNew()
	if err != nil {
		panic(err.Error())
	}

	data, err := uiTextViewGladeBytes()
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
		t0, _ := builder.GetObject("tw")
		t1, _ := t0.(*gtk.TextView)
		ret.tw = t1
	}

	{
		t0, _ := builder.GetObject("button_save")
		t1, _ := t0.(*gtk.Button)
		ret.button_save = t1
	}

	{
		t0, _ := builder.GetObject("button_close")
		t1, _ := t0.(*gtk.Button)
		ret.button_close = t1
	}

	ret.root.SetTransientFor(ret.transient)

	ret.button_save.Connect(
		"clicked",
		func() {
			dialog, err := gtk.DialogNew()
			if err != nil {
				panic(err.Error())
			}

			if ret.transient != nil {
				dialog.SetTransientFor(ret.root)
			}

			chooser, err := gtk.FileChooserWidgetNew(gtk.FILE_CHOOSER_ACTION_SAVE)
			if err != nil {
				panic(err.Error())
			}
			dialog.AddButton("Save", gtk.RESPONSE_ACCEPT)
			dialog.AddButton("Cancel", gtk.RESPONSE_CANCEL)

			dialog.SetTitle("Select File Name To Write Text")

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

	ret.button_close.Connect(
		"clicked",
		func() {
			ret.root.Close()
		},
	)

	ret.root.SetTitle(title)
	{
		b, _ := ret.tw.GetBuffer()
		b.SetText(text)
	}

	return ret
}

func (self *UITextViewer) GetWindow() *gtk.Window {
	return self.root
}

func (self *UITextViewer) Show() {
	self.root.ShowAll()
}

func (self *UITextViewer) Destroy() {
	self.root.Destroy()
}
