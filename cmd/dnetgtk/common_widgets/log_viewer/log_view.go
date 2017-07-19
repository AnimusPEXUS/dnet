package log_viewer

import (
	"github.com/AnimusPEXUS/gologger"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type UILogViewer struct {
	root      *gtk.Paned
	store     *gtk.ListStore
	treev_log *gtk.TreeView
	textv_log *gtk.TextView
}

func UILogViewerNew(
	logger *gologger.Logger,
) *UILogViewer {

	ret := new(UILogViewer)

	if s, err := gtk.ListStoreNew(
		glib.TYPE_STRING,
		glib.TYPE_STRING,
		glib.TYPE_STRING,
	); err != nil {
		panic("ой всё.. " + err.Error())
	} else {
		ret.store = s
	}

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

		{
			rend, _ := gtk.CellRendererTextNew()
			column, _ := gtk.TreeViewColumnNewWithAttribute(
				"Event",
				rend,
				"text",
				0,
			)
			ret.treev_log.AppendColumn(column)
		}

		{
			rend, _ := gtk.CellRendererTextNew()
			column, _ := gtk.TreeViewColumnNewWithAttribute(
				"Time",
				rend,
				"text",
				1,
			)
			ret.treev_log.AppendColumn(column)
		}

		{
			rend, _ := gtk.CellRendererTextNew()
			column, _ := gtk.TreeViewColumnNewWithAttribute(
				"Message",
				rend,
				"text",
				2,
			)
			ret.treev_log.AppendColumn(column)
		}

	}

	{
		t0, _ := builder.GetObject("textv_log")
		t1, _ := t0.(*gtk.TextView)
		ret.textv_log = t1
	}

	logger.ConnectCallback(ret.LoggerCallback)

	ret.treev_log.SetModel(ret.store)

	return ret
}

func (self *UILogViewer) GetRoot() *gtk.Paned {
	return self.root
}

func (self *UILogViewer) LoggerCallback(
	entry *gologger.LogEntry,
	logger *gologger.Logger,
) {

	iter := self.store.Append()
	self.store.Set(
		iter,
		[]int{0, 1, 2},
		[]interface{}{
			entry.TypeStringT(),
			entry.Time.String(),
			entry.Text,
		},
	)

}
