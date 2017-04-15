package main

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type UIWindowMainTabApplications struct {
	main_window *UIWindowMain

	button_refresh_application_modules *gtk.Button
	button_refresh_application_presets *gtk.Button
	button_enable_application_preset   *gtk.Button
	button_disable_application_preset  *gtk.Button
	button_create_application_preset       *gtk.Button
	tw_networks_presets            *gtk.TreeView
	tw_networks_modules            *gtk.TreeView
	application_presets               *gtk.ListStore
	application_modules               *gtk.ListStore
}


func UIWindowMainTabApplicationsNew(
	builder *gtk.Builder,
	main_window *UIWindowMain,
) (*UIWindowMainTabApplications, error) {

	ret := new(UIWindowMainTabApplications)

	ret.main_window = main_window

	{
		t0, _ := builder.GetObject("button_refresh_application_modules")
		t1, _ := t0.(*gtk.Button)
		ret.button_refresh_application_modules = t1
	}

	{
		t0, _ := builder.GetObject("button_refresh_application_presets")
		t1, _ := t0.(*gtk.Button)
		ret.button_refresh_application_presets = t1
	}

	{
		t0, _ := builder.GetObject("button_enable_application_preset")
		t1, _ := t0.(*gtk.Button)
		ret.button_enable_application_preset = t1
	}

	{
		t0, _ := builder.GetObject("button_disable_application_preset")
		t1, _ := t0.(*gtk.Button)
		ret.button_disable_application_preset = t1
	}

	{
		t0, _ := builder.GetObject("button_create_application_preset")
		t1, _ := t0.(*gtk.Button)
		ret.button_create_application_preset = t1
	}

	{
		t0, _ := builder.GetObject("tw_application_presets")
		t1, _ := t0.(*gtk.TreeView)
		ret.tw_application_presets = t1
	}

	{
		t0, _ := builder.GetObject("tw_application_modules")
		t1, _ := t0.(*gtk.TreeView)
		ret.tw_application_modules = t1
	}

	{
		ret.networks_presets, _ = gtk.ListStoreNew(
			glib.TYPE_STRING,  // preset name
			glib.TYPE_STRING,  // module name
			glib.TYPE_BOOLEAN, // enabled?
			glib.TYPE_STRING,  // status
			glib.TYPE_BOOLEAN, // has errors?
			glib.TYPE_STRING,  // brief info
		)

		ret.networks_modules, _ = gtk.ListStoreNew(
			glib.TYPE_STRING, // Name
			glib.TYPE_STRING, // module name
			glib.TYPE_STRING, // Descr
		)

		ret.tw_networks_presets.SetModel(ret.networks_presets)
		ret.tw_networks_modules.SetModel(ret.networks_modules)

	}

	{
		{
			// setup coumns in tw_networks_presets
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Preset Name",
					rend,
					"text",
					0,
				)
				ret.tw_networks_presets.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Module Name",
					rend,
					"text",
					1,
				)
				ret.tw_networks_presets.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererToggleNew()
				rend.SetActivatable(false)
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Enabled?",
					rend,
					"active",
					2,
				)
				ret.tw_networks_presets.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererTextNew()
				//rend.SetActivatable(false)
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Status",
					rend,
					"text",
					3,
				)
				ret.tw_networks_presets.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererToggleNew()
				rend.SetActivatable(false)
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Errors?",
					rend,
					"active",
					4,
				)
				ret.tw_networks_presets.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Info",
					rend,
					"text",
					5,
				)
				ret.tw_networks_presets.AppendColumn(column)
			}
		}

		{
			// setup coumns in tw_networks_modules
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Name",
					rend,
					"text",
					0,
				)
				ret.tw_networks_modules.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Module Name",
					rend,
					"text",
					1,
				)
				ret.tw_networks_modules.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Description",
					rend,
					"text",
					2,
				)
				ret.tw_networks_modules.AppendColumn(column)
			}
		}
	}

	return ret, nil
}

func (self *UIWindowMainTabNetworks) GetSelectedPresetName() (string, bool) {
	sel, _ := self.tw_application_presets.GetSelection()
	model, iter, ok := sel.GetSelected()
	if ok {
		val, _ := model.(*gtk.TreeModel).GetValue(iter, 0)
		val_str, _ := val.GetString()
		return val_str, true
	}
	return "", false
}

func (self *UIWindowMainTabNetworks) GetSelectedModuleName() (string, bool) {
	sel, _ := self.tw_application_modules.GetSelection()
	model, iter, ok := sel.GetSelected()
	if ok {
		val, _ := model.(*gtk.TreeModel).GetValue(iter, 1)
		val_str, _ := val.GetString()
		return val_str, true
	}
	return "", false
}
