package log_viewer

import (
	"github.com/AnimusPEXUS/gologger"

	"github.com/gotk3/gotk3/gtk"
)

type UILogViewer struct {
	root      *gtk.Paned
	treev_log *gtk.TreeView
	textv_log *gtk.TextView
}

func UILogViewerNew(
	logger *gologger.Logger,
) *UILogViewer {

	ret := new(UILogViewer)

	builder, err := gtk.BuilderNew()
	if err != nil {
		panic(err.Error())
	}

	data, err := uiWidgetGladeBytes()
	if err != nil {
		panic(err.Error())
	}

	err = builder.AddFromString(string(data))
	if err != nil {
		panic(err.Error())
	}

	{
		t0, _ := builder.GetObject("root")
		t1, _ := t0.(*gtk.Paned)
		ret.root = t1
	}

	{
		t0, _ := builder.GetObject("treev_log")
		t1, _ := t0.(*gtk.TreeView)
		ret.treev_log = t1
	}

	{
		t0, _ := builder.GetObject("textv_log")
		t1, _ := t0.(*gtk.TextView)
		ret.textv_log = t1
	}

	return ret
}

func (self *UILogViewer) GetRoot() *gtk.Paned {
	return self.root
}
