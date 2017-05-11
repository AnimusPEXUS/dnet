package progress_window

import (
	"github.com/gotk3/gotk3/gtk"
)

type UIProgressWindow struct {
	root      *gtk.Window
	transient *gtk.Window

	label    *gtk.Label
	spinner  *gtk.Spinner
	progress *gtk.LevelBar
	button   *gtk.Button

	cancel_callback      func(callback_data interface{})
	cancel_callback_data interface{}

	finished bool
}

func UIProgressWindowNew(
	transient *gtk.Window,
	cancel_callback func(cancel_callback_data interface{}),
	cancel_callback_data interface{},
) *UIProgressWindow {

	ret := new(UIProgressWindow)
	ret.finished = false
	ret.cancel_callback = cancel_callback
	ret.cancel_callback_data = cancel_callback_data

	ret.transient = transient

	builder, err := gtk.BuilderNew()
	if err != nil {
		panic(err.Error())
	}

	data, err := uiProgressWindowGladeBytes()
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
		ret.root = t1
	}

	{
		t0, _ := builder.GetObject("progress")
		t1, _ := t0.(*gtk.LevelBar)
		ret.progress = t1
	}

	{
		t0, _ := builder.GetObject("spinner")
		t1, _ := t0.(*gtk.Spinner)
		ret.spinner = t1
	}

	{
		t0, _ := builder.GetObject("label")
		t1, _ := t0.(*gtk.Label)
		ret.label = t1
	}

	{
		t0, _ := builder.GetObject("button")
		t1, _ := t0.(*gtk.Button)
		ret.button = t1
	}

	ret.root.Connect("destroy", ret.onDestroy)
	ret.button.Connect("clicked", ret.onButtonClick)

	return ret
}

func (self *UIProgressWindow) SetTitle(text string) {
	self.root.SetTitle(text)
}

func (self *UIProgressWindow) SetText(text string) {
	self.label.SetText(text)
}

func (self *UIProgressWindow) SetActive(value bool) {
	if value {
		self.spinner.Start()
	} else {
		self.spinner.Stop()
	}
}

func (self *UIProgressWindow) SetProgress(value float64) {
	self.progress.SetValue(value)
}

func (self *UIProgressWindow) Show() {
	self.root.ShowAll()
}

func (self *UIProgressWindow) Finish() {
	self.finished = true
	self.root.Close()
}

func (self *UIProgressWindow) Close() {
	self.root.Close()
}

func (self *UIProgressWindow) onDestroy() {
	if !self.finished {
		if self.cancel_callback != nil {
			go self.cancel_callback(self.cancel_callback_data)
		}
	}
}

func (self *UIProgressWindow) onButtonClick() {
	self.root.Close()
}
