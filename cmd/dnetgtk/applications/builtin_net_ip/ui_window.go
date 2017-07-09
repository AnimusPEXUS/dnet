package builtin_net_ip

import (
	//"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type UIWindow struct {
	instance *Instance

	transient_for *gtk.Window

	window                       *gtk.Window
	button_open                  *gtk.Button
	tw                           *gtk.TreeView
	entry_probe                  *gtk.Entry
	button_probe                 *gtk.Button
	cb_tcp_listen_enabled        *gtk.CheckButton
	label_tcp_listen_status      *gtk.Label
	button_tcp_listen_start      *gtk.Button
	button_tcp_listen_stop       *gtk.Button
	entry_tcp_listen_port        *gtk.Entry
	button_tcp_listen_port_apply *gtk.Button
	entry_udp_port               *gtk.Entry
	button_udp_port_apply        *gtk.Button
	cb_udp_beacon_enabled        *gtk.CheckButton
	label_udp_beacon_status      *gtk.Label
	button_udp_beacon_start      *gtk.Button
	button_udp_beacon_stop       *gtk.Button
	entry_udp_beacon_interval    *gtk.Entry
	button_beacon_interval_apply *gtk.Button
	cb_udp_locator_enabled       *gtk.CheckButton
	label_udp_locator_status     *gtk.Label
	button_udp_locator_start     *gtk.Button
	button_udp_locator_stop      *gtk.Button
}

func UIWindowNew(instance *Instance) *UIWindow {

	ret := new(UIWindow)
	ret.instance = instance

	builder, err := gtk.BuilderNew()
	if err != nil {
		panic(err.Error())
	}

	data, err := uiWindowGladeBytes()
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
		t0, _ := builder.GetObject("button_open")
		t1, _ := t0.(*gtk.Button)
		ret.button_open = t1
	}

	{
		t0, _ := builder.GetObject("tw")
		t1, _ := t0.(*gtk.TreeView)
		ret.tw = t1
	}

	{
		t0, _ := builder.GetObject("entry_probe")
		t1, _ := t0.(*gtk.Entry)
		ret.entry_probe = t1
	}

	{
		t0, _ := builder.GetObject("button_probe")
		t1, _ := t0.(*gtk.Button)
		ret.button_probe = t1
	}

	{
		t0, _ := builder.GetObject("cb_tcp_listen_enabled")
		t1, _ := t0.(*gtk.CheckButton)
		ret.cb_tcp_listen_enabled = t1
	}

	{
		t0, _ := builder.GetObject("label_tcp_listen_status")
		t1, _ := t0.(*gtk.Label)
		ret.label_tcp_listen_status = t1
	}

	{
		t0, _ := builder.GetObject("button_tcp_listen_start")
		t1, _ := t0.(*gtk.Button)
		ret.button_tcp_listen_start = t1
	}

	{
		t0, _ := builder.GetObject("button_tcp_listen_stop")
		t1, _ := t0.(*gtk.Button)
		ret.button_tcp_listen_stop = t1
	}

	{
		t0, _ := builder.GetObject("entry_tcp_listen_port")
		t1, _ := t0.(*gtk.Entry)
		ret.entry_tcp_listen_port = t1
	}

	{
		t0, _ := builder.GetObject("button_tcp_listen_port_apply")
		t1, _ := t0.(*gtk.Button)
		ret.button_tcp_listen_port_apply = t1
	}

	{
		t0, _ := builder.GetObject("entry_udp_port")
		t1, _ := t0.(*gtk.Entry)
		ret.entry_udp_port = t1
	}

	{
		t0, _ := builder.GetObject("button_udp_port_apply")
		t1, _ := t0.(*gtk.Button)
		ret.button_udp_port_apply = t1
	}

	{
		t0, _ := builder.GetObject("cb_udp_beacon_enabled")
		t1, _ := t0.(*gtk.CheckButton)
		ret.cb_udp_beacon_enabled = t1
	}

	{
		t0, _ := builder.GetObject("label_udp_beacon_status")
		t1, _ := t0.(*gtk.Label)
		ret.label_udp_beacon_status = t1
	}

	{
		t0, _ := builder.GetObject("button_udp_beacon_start")
		t1, _ := t0.(*gtk.Button)
		ret.button_udp_beacon_start = t1
	}

	{
		t0, _ := builder.GetObject("button_udp_beacon_stop")
		t1, _ := t0.(*gtk.Button)
		ret.button_udp_beacon_stop = t1
	}

	{
		t0, _ := builder.GetObject("entry_udp_beacon_interval")
		t1, _ := t0.(*gtk.Entry)
		ret.entry_udp_beacon_interval = t1
	}

	{
		t0, _ := builder.GetObject("button_beacon_interval_apply")
		t1, _ := t0.(*gtk.Button)
		ret.button_beacon_interval_apply = t1
	}

	{
		t0, _ := builder.GetObject("cb_udp_locator_enabled")
		t1, _ := t0.(*gtk.CheckButton)
		ret.cb_udp_locator_enabled = t1
	}

	{
		t0, _ := builder.GetObject("label_udp_locator_status")
		t1, _ := t0.(*gtk.Label)
		ret.label_udp_locator_status = t1
	}

	{
		t0, _ := builder.GetObject("button_udp_locator_start")
		t1, _ := t0.(*gtk.Button)
		ret.button_udp_locator_start = t1
	}

	{
		t0, _ := builder.GetObject("button_udp_locator_stop")
		t1, _ := t0.(*gtk.Button)
		ret.button_udp_locator_stop = t1
	}

	return ret
}

func (self *UIWindow) Show() {
	self.window.ShowAll()
}
