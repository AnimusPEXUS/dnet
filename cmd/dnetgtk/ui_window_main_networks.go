package main

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"bitbucket.org/AnimusPEXUS/dnet"
	"bitbucket.org/AnimusPEXUS/dnet/common_types"
)

type UIWindowMainTabNetworks struct {
	main_window *UIWindowMain

	button_refresh_network_presets *gtk.Button
	button_refresh_network_modules *gtk.Button
	button_create_new_preset       *gtk.Button
	tw_networks_presets            *gtk.TreeView
	tw_networks_modules            *gtk.TreeView
	networks_presets               *gtk.ListStore
	networks_modules               *gtk.ListStore
}

func UIWindowMainTabNetworksNew(
	builder *gtk.Builder,
	main_window *UIWindowMain,
) (*UIWindowMainTabNetworks, error) {

	ret := new(UIWindowMainTabNetworks)

	ret.main_window = main_window

	{
		t0, _ := builder.GetObject("button_refresh_network_modules")
		t1, _ := t0.(*gtk.Button)
		ret.button_refresh_network_modules = t1
	}

	{
		t0, _ := builder.GetObject("button_refresh_network_presets")
		t1, _ := t0.(*gtk.Button)
		ret.button_refresh_network_presets = t1
	}

	{
		t0, _ := builder.GetObject("button_create_new_preset")
		t1, _ := t0.(*gtk.Button)
		ret.button_create_new_preset = t1
	}

	{
		t0, _ := builder.GetObject("tw_networks_presets")
		t1, _ := t0.(*gtk.TreeView)
		ret.tw_networks_presets = t1
	}

	{
		t0, _ := builder.GetObject("tw_networks_modules")
		t1, _ := t0.(*gtk.TreeView)
		ret.tw_networks_modules = t1
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

	ret.button_refresh_network_modules.Connect(
		"clicked",
		func() {

			model := ret.networks_modules
			for {
				iter, _ := model.GetIterFirst()
				if iter != nil {
					model.Remove(iter)
				} else {
					break
				}
			}

			for _, i := range dnet.BUILTIN_NETWORK_MODULES {
				iter := model.Append()
				model.Set(
					iter,
					[]int{0, 1, 2},
					[]interface{}{i.Name(), i.WorkingName(), i.Description()},
				)
			}
		},
	)

	ret.button_create_new_preset.Connect(
		"clicked",
		func() {

			sel, _ := ret.tw_networks_modules.GetSelection()

			model, iter, ok := sel.GetSelected()

			if ok {
				// model := self.networks_modules

				val, _ := model.(*gtk.TreeModel).GetValue(iter, 1)
				val_str, _ := val.GetString()
				// fmt.Println("Value", val_str)

				var ext_a common_types.NetworkModule

				for _, i := range dnet.BUILTIN_NETWORK_MODULES {
					if i.WorkingName() == val_str {
						ext_a = i
						goto succ
					}
				}
				goto err

			succ:
				{
					w := UINetworkModuleConfigEditorNew(
						ret.main_window.controller,
						ret.main_window.win,
						false,
						"",
						val_str,
						ext_a.SampleConfig(),
					)
					w.Show()

					goto exit
				}

			err:
				glib.IdleAdd(
					func() {
						d := gtk.MessageDialogNew(
							ret.main_window.win,
							0,
							gtk.MESSAGE_ERROR,
							gtk.BUTTONS_OK,
							"Something wrong here, this shold not bin hapened. terminating..",
						)
						d.Run()
						d.Destroy()
						panic("programming error")
					},
				)

			exit:
			}
		},
	)

	ret.button_refresh_network_presets.Connect(
		"clicked",
		func() {

			{

				to_add := []string{}
				to_remove := []string{}

				{
					dnet_s_names := ret.main_window.controller.NetworkPresetList()
					model_s_names := []string{}

					{
						mdl := ret.networks_presets
						iter, ok := mdl.GetIterFirst()
						for ok {
							val, _ := mdl.GetValue(iter, 0)
							val_str, _ := val.GetString()
							model_s_names = append(model_s_names, val_str)
							ok = mdl.IterNext(iter)
						}
					}

					{
					searching1:
						for _, i := range dnet_s_names {
							for _, j := range model_s_names {
								if i == j {
									continue searching1
								}
							}
							to_add = append(to_add, i)
						}
					}

					{
					searching2:
						for _, i := range model_s_names {
							for _, j := range dnet_s_names {
								if i == j {
									continue searching2
								}
							}
							to_remove = append(to_remove, i)
						}
					}
				}

				{
					mdl := ret.networks_presets
					for _, i := range to_add {
						iter := mdl.Append()

						/*
							glib.TYPE_STRING,  // preset name
							glib.TYPE_STRING,  // module name
							glib.TYPE_BOOLEAN, // enabled?
							glib.TYPE_STRING,  // status
							glib.TYPE_BOOLEAN, // has errors?
							glib.TYPE_STRING,  // brief info
						*/

						mdl.Set(
							iter,
							[]int{0, 1, 2, 3, 4, 5},
							[]interface{}{
								i,
								"",
								false,
								"",
								false,
								"",
							},
						)
					}
				}

				{
					mdl := ret.networks_presets
					iter, ok := mdl.GetIterFirst()
				searching3:
					for ok {
						val, _ := mdl.GetValue(iter, 0)
						val_str, _ := val.GetString()
						for _, i := range to_remove {
							if val_str == i {
								ok = mdl.Remove(iter)
								continue searching3
							}
						}
						ok = mdl.IterNext(iter)
					}
				}

				{
					/*
						dnet_s_names := dnet.NetworkPresetList()

								mdl := ret.networks_presets
								iter, ok := mdl.GetIterFirst()
								for ok {
									val, _ := mdl.GetValue(iter, 0)
									val_str, _ := val.GetString()


									ok = mdl.IterNext(iter)
								}
					*/
				}

			}

		},
	)

	return ret, nil
}
