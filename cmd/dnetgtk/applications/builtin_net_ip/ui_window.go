package builtin_net_ip

import (
	//"github.com/gotk3/gotk3/glib"

	"strconv"
	"time"

	"github.com/AnimusPEXUS/goworker"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type UIWindow struct {
	instance *Instance

	worker_monitor *worker.Worker

	transient_for *gtk.Window

	window                         *gtk.Window
	button_open                    *gtk.Button
	tw                             *gtk.TreeView
	entry_probe                    *gtk.Entry
	button_probe                   *gtk.Button
	cb_tcp_listener_enabled        *gtk.CheckButton
	label_tcp_listener_status      *gtk.Label
	button_tcp_listener_start      *gtk.Button
	button_tcp_listener_stop       *gtk.Button
	entry_tcp_listener_port        *gtk.Entry
	button_tcp_listener_port_apply *gtk.Button
	entry_udp_port                 *gtk.Entry
	button_udp_port_apply          *gtk.Button
	cb_udp_beacon_enabled          *gtk.CheckButton
	label_udp_beacon_status        *gtk.Label
	button_udp_beacon_start        *gtk.Button
	button_udp_beacon_stop         *gtk.Button
	entry_udp_beacon_interval      *gtk.Entry
	button_beacon_interval_apply   *gtk.Button
	cb_udp_listener_enabled        *gtk.CheckButton
	label_udp_listener_status      *gtk.Label
	button_udp_listener_start      *gtk.Button
	button_udp_listener_stop       *gtk.Button
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
		t0, _ := builder.GetObject("cb_tcp_listener_enabled")
		t1, _ := t0.(*gtk.CheckButton)
		ret.cb_tcp_listener_enabled = t1
	}

	{
		t0, _ := builder.GetObject("label_tcp_listener_status")
		t1, _ := t0.(*gtk.Label)
		ret.label_tcp_listener_status = t1
	}

	{
		t0, _ := builder.GetObject("button_tcp_listener_start")
		t1, _ := t0.(*gtk.Button)
		ret.button_tcp_listener_start = t1
	}

	{
		t0, _ := builder.GetObject("button_tcp_listener_stop")
		t1, _ := t0.(*gtk.Button)
		ret.button_tcp_listener_stop = t1
	}

	{
		t0, _ := builder.GetObject("entry_tcp_listener_port")
		t1, _ := t0.(*gtk.Entry)
		ret.entry_tcp_listener_port = t1
	}

	{
		t0, _ := builder.GetObject("button_tcp_listener_port_apply")
		t1, _ := t0.(*gtk.Button)
		ret.button_tcp_listener_port_apply = t1
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
		t0, _ := builder.GetObject("cb_udp_listener_enabled")
		t1, _ := t0.(*gtk.CheckButton)
		ret.cb_udp_listener_enabled = t1
	}

	{
		t0, _ := builder.GetObject("label_udp_listener_status")
		t1, _ := t0.(*gtk.Label)
		ret.label_udp_listener_status = t1
	}

	{
		t0, _ := builder.GetObject("button_udp_listener_start")
		t1, _ := t0.(*gtk.Button)
		ret.button_udp_listener_start = t1
	}

	{
		t0, _ := builder.GetObject("button_udp_listener_stop")
		t1, _ := t0.(*gtk.Button)
		ret.button_udp_listener_stop = t1
	}

	{
		glib.IdleAdd(
			func() {
				ret.cb_tcp_listener_enabled.SetActive(ret.instance.cfg.TCPListenerEnabled)
				ret.cb_udp_beacon_enabled.SetActive(ret.instance.cfg.UDPBeaconEnabled)
				ret.cb_udp_listener_enabled.SetActive(ret.instance.cfg.UDPListenerEnabled)

				ret.entry_tcp_listener_port.SetText(
					strconv.Itoa(ret.instance.cfg.TCPListenerPort),
				)
				ret.entry_udp_port.SetText(
					strconv.Itoa(ret.instance.cfg.UDPPort),
				)
				ret.entry_udp_beacon_interval.SetText(
					strconv.Itoa(ret.instance.cfg.UDPBeaconInterval),
				)
			},
		)
	}

	ret.window.Connect(
		"destroy",
		func() {
			ret.worker_monitor.Stop()
		},
	)

	ret.cb_tcp_listener_enabled.Connect(
		"toggled",
		func() {
			ret.instance.cfg.TCPListenerEnabled =
				ret.cb_tcp_listener_enabled.GetActive()
			go ret.instance.SaveConfig()
		},
	)

	ret.cb_udp_beacon_enabled.Connect(
		"toggled",
		func() {
			ret.instance.cfg.UDPBeaconEnabled =
				ret.cb_udp_beacon_enabled.GetActive()
			go ret.instance.SaveConfig()
		},
	)

	ret.cb_udp_listener_enabled.Connect(
		"toggled",
		func() {
			ret.instance.cfg.UDPListenerEnabled =
				ret.cb_udp_listener_enabled.GetActive()
			go ret.instance.SaveConfig()
		},
	)

	ret.button_tcp_listener_port_apply.Connect(
		"clicked",
		func() {
			v1, _ := ret.entry_tcp_listener_port.GetText()
			if v, err := strconv.Atoi(v1); err != nil {
			} else {
				ret.instance.cfg.TCPListenerPort = v
				go ret.instance.SaveConfig()
			}
		},
	)

	ret.button_beacon_interval_apply.Connect(
		"clicked",
		func() {
			v1, _ := ret.entry_udp_beacon_interval.GetText()
			if v, err := strconv.Atoi(v1); err != nil {
			} else {
				ret.instance.cfg.UDPBeaconInterval = v
				go ret.instance.SaveConfig()
			}
		},
	)

	ret.button_udp_port_apply.Connect(
		"clicked",
		func() {
			v1, _ := ret.entry_udp_port.GetText()
			if v, err := strconv.Atoi(v1); err != nil {
			} else {
				ret.instance.cfg.UDPPort = v
				go ret.instance.SaveConfig()
			}
		},
	)

	ret.button_tcp_listener_start.Connect(
		"clicked",
		func() {
			ret.instance.tcp_listener.Start()
		},
	)

	ret.button_tcp_listener_stop.Connect(
		"clicked",
		func() {
			ret.instance.tcp_listener.Stop()
		},
	)

	ret.button_udp_beacon_start.Connect(
		"clicked",
		func() {
			ret.instance.udp_beacon.Start()
		},
	)

	ret.button_udp_beacon_stop.Connect(
		"clicked",
		func() {
			ret.instance.udp_beacon.Stop()
		},
	)

	ret.button_udp_listener_start.Connect(
		"clicked",
		func() {
			ret.instance.udp_listener.Start()
		},
	)

	ret.button_udp_listener_stop.Connect(
		"clicked",
		func() {
			ret.instance.udp_listener.Stop()
		},
	)

	ret.worker_monitor = worker.New(
		func(
			set_starting func(),
			set_working func(),
			set_stopping func(),
			set_stopped func(),

			is_stop_flag func() bool,

		) {
			for !is_stop_flag() {
				c := make(chan bool, 1)
				glib.IdleAdd(
					func() {
						if ret.instance.tcp_listener == nil {
							ret.label_tcp_listener_status.SetText("N/A")
						} else {
							ret.label_tcp_listener_status.SetText(
								ret.instance.tcp_listener.Status().StringT(),
							)
						}
						if ret.instance.udp_beacon == nil {
							ret.label_udp_beacon_status.SetText("N/A")
						} else {
							ret.label_udp_beacon_status.SetText(
								ret.instance.udp_beacon.Status().StringT(),
							)
						}
						if ret.instance.udp_listener == nil {
							ret.label_udp_listener_status.SetText("N/A")
						} else {
							ret.label_udp_listener_status.SetText(
								ret.instance.udp_listener.Status().StringT(),
							)
						}
						c <- true
					},
				)
				<-c
				time.Sleep(time.Second)
			}
		},
	)

	ret.worker_monitor.Start()

	return ret
}

func (self *UIWindow) Show() error {
	self.window.ShowAll()
	return nil
}
