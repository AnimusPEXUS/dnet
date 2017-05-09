package main

import (
	"fmt"
	"strings"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"github.com/AnimusPEXUS/dnet/common_types"
)

type UIWindowMainTabApplications struct {
	main_window *UIWindowMain

	button_refresh_application_modules *gtk.Button
	button_refresh_application_presets *gtk.Button
	button_enable_application_preset   *gtk.Button
	button_disable_application_preset  *gtk.Button
	button_accept_application          *gtk.Button
	tw_application_presets             *gtk.TreeView
	tw_application_modules             *gtk.TreeView
	application_presets                *gtk.ListStore
	application_modules                *gtk.ListStore
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
		t0, _ := builder.GetObject("button_accept_application")
		t1, _ := t0.(*gtk.Button)
		ret.button_accept_application = t1
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
		ret.application_presets, _ = gtk.ListStoreNew(
			glib.TYPE_STRING,  // Name
			glib.TYPE_BOOLEAN, // builtin?
			glib.TYPE_BOOLEAN, // enabled?
			glib.TYPE_STRING,  // Status
			glib.TYPE_STRING,  // Checksum
			glib.TYPE_STRING,  // Last ReKey time
		)

		ret.application_modules, _ = gtk.ListStoreNew(
			glib.TYPE_STRING,  // Name
			glib.TYPE_BOOLEAN, // builtin?
			glib.TYPE_STRING,  // Checksum
			glib.TYPE_STRING,  // Path
			glib.TYPE_STRING,  // Descr
		)

		ret.tw_application_presets.SetModel(ret.application_presets)
		ret.tw_application_modules.SetModel(ret.application_modules)

	}

	{
		{
			// setup columns in tw_application_presets
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Name",
					rend,
					"text",
					0,
				)
				ret.tw_application_presets.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererToggleNew()
				rend.SetActivatable(false)
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"BuiltIn?",
					rend,
					"active",
					1,
				)
				ret.tw_application_presets.AppendColumn(column)
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
				ret.tw_application_presets.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Status",
					rend,
					"text",
					3,
				)
				ret.tw_application_presets.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Checksum",
					rend,
					"text",
					4,
				)
				ret.tw_application_presets.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererTextNew()
				//rend.SetActivatable(false)
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Last ReKey Time",
					rend,
					"text",
					5,
				)
				ret.tw_application_presets.AppendColumn(column)
			}
		}

		{
			// setup columns in tw_application_modules
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Name",
					rend,
					"text",
					0,
				)
				ret.tw_application_modules.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererToggleNew()
				rend.SetActivatable(false)
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"BuiltIn?",
					rend,
					"active",
					1,
				)
				ret.tw_application_modules.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Checksum",
					rend,
					"text",
					2,
				)
				ret.tw_application_modules.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Path",
					rend,
					"text",
					3,
				)
				ret.tw_application_modules.AppendColumn(column)
			}
			{
				rend, _ := gtk.CellRendererTextNew()
				column, _ := gtk.TreeViewColumnNewWithAttribute(
					"Description",
					rend,
					"text",
					4,
				)
				ret.tw_application_modules.AppendColumn(column)
			}
		}
	}

	ret.button_refresh_application_presets.Connect(
		"clicked",
		func() {
			mdl := ret.application_presets

			{
				iter, ok := mdl.GetIterFirst()
				for {
					if !ok {
						break
					}
					ok = mdl.Remove(iter)
				}
			}

			for _, i := range ret.main_window.controller.application_presets {
				iter := mdl.Append()
				cs := "N/A"
				if !i.DBStatus.Builtin {
					cs = i.DBStatus.Checksum
				}
				mdl.Set(
					iter,
					[]int{0, 1, 2, 3, 4, 5},
					[]interface{}{
						i.Name,
						i.DBStatus.Builtin,
						i.DBStatus.Enabled,
						strings.Title(i.Instance.Status().String()),
						cs,
						i.DBStatus.LastDBReKey.String(),
					},
				)
			}

		},
	)

	ret.button_refresh_application_modules.Connect(
		"clicked",
		func() {
			// TODO: make this assinc
			mdl := ret.application_modules

			{
				iter, ok := mdl.GetIterFirst()
				for {
					if !ok {
						break
					}
					ok = mdl.Remove(iter)
				}
			}

			res := ret.main_window.controller.ModSearcher.ListModules()

			for _, i := range res {
				if i.builtin {
					iter := mdl.Append()
					mod, err := i.Mod()
					if err != nil {
						fmt.Printf("error executing .Mod(): %s\n", err.Error())
						mdl.Remove(iter)
						continue
					}

					mdl.Set(
						iter,
						[]int{0, 1, 2, 3, 4},
						[]interface{}{
							mod.Name(),
							true,
							"N/A",
							"N/A",
							mod.Description(),
						},
					)
				}
			}
		},
	)

	ret.button_accept_application.Connect(
		"clicked",
		func() {

			var (
				builtin  bool
				name     string
				checksum string
			)

			sel, _ := ret.tw_application_modules.GetSelection()
			model, iter, ok := sel.GetSelected()
			if ok {

				{
					val, _ := model.(*gtk.TreeModel).GetValue(iter, 0)
					name, _ = val.GetString()
				}

				{
					val, _ := model.(*gtk.TreeModel).GetValue(iter, 1)
					builtin_t, _ := val.GoValue()
					builtin, _ = builtin_t.(bool)
				}

				{
					val, _ := model.(*gtk.TreeModel).GetValue(iter, 2)
					checksum, _ = val.GetString()
				}

				/*
					fmt.Println("builtin", builtin)
					fmt.Println("name", name)
					fmt.Println("checksum", checksum)
				*/

				{

					checksum_obj := (*common_types.ModuleChecksum)(nil)
					if !builtin {
						var err error

						checksum_obj, err =
							common_types.ModuleChecksumNewFromString(checksum)

						if err != nil {
							panic("programming error: " + err.Error())
						}
					}

					name_obj, err :=
						common_types.ModuleNameNew(name)

					if err != nil {
						panic("programming error: " + err.Error())
					}

					w, err := UIWindowModuleAcceptorNew(
						ret.main_window,
						builtin,
						name_obj,
						checksum_obj,
					)

					if err != nil {
						panic(
							"programming error (this window should display" +
								" errors by it's self, not return them): " + err.Error(),
						)
					}

					w.Show()
				}
			}

		},
	)

	ret.button_enable_application_preset.Connect(
		"clicked",
		func() {
			value, ok := ret.GetSelectedPresetName()
			if ok {
				ret.main_window.controller.EnableModule(value, true)
			}
		},
	)

	ret.button_disable_application_preset.Connect(
		"clicked",
		func() {
			value, ok := ret.GetSelectedPresetName()
			if ok {
				ret.main_window.controller.EnableModule(value, false)
			}
		},
	)

	return ret, nil
}

func (self *UIWindowMainTabApplications) GetSelectedPresetName() (
	string,
	bool,
) {
	sel, _ := self.tw_application_presets.GetSelection()
	model, iter, ok := sel.GetSelected()
	if ok {
		val, _ := model.(*gtk.TreeModel).GetValue(iter, 0)
		val_str, _ := val.GetString()
		return val_str, true
	}
	return "", false
}

/*
func (self *UIWindowMainTabApplications) GetSelectedModuleName() (
	string,
	bool,
) {
	sel, _ := self.tw_application_modules.GetSelection()
	model, iter, ok := sel.GetSelected()
	if ok {
		val, _ := model.(*gtk.TreeModel).GetValue(iter, 0)
		val_str, _ := val.GetString()
		return val_str, true
	}
	return "", false
}
*/
