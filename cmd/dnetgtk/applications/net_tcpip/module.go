package tcpip

import (
	"bytes"

	"github.com/go-ini/ini"

	"bitbucket.org/AnimusPEXUS/dnet/common_types"
)

type Module struct {
}

func NewModule() *Module {
	ret := new(Module)
	return ret
}

func (self *Module) Name() string {
	return "tcpip"
}

func (self *Module) Title() string {
	return "TCP/IP"
}

func (self *Module) Description() string {
	return "DNet module for working with TCP/IP networks using TLS encryption"
}

func (self *Module) CanExplore() bool {
	return true
}

func (self *Module) CanBeacon() bool {
	return true
}

func (self *Module) CanListen() bool {
	return true
}

func (self *Module) Instance() (common_types.NetworkInstance, error) {

	ret := new(Instance)
	ret.module = self

	return ret, nil
}

func (self *Module) SampleConfig() string {
	// This is My first experience with ini module and I'm now only experimenting,
	// so this section and config for this module will be changed
	f := ini.Empty()

	s := f.Section("binding")
	s.Comment = "main configuration section of TCPIP module instance"

	k := s.Key("bind_by")
	k.Comment = `possible values: "interface", "address", "special"`
	k.SetValue("address")

	k = s.Key("bind")
	k.Comment =
		"depending on bind_by value, set this to interface name\n" +
			"; (e.g. lo, eth0, enp0s31f6 etc.) or to ip address.\n" +
			"; 'special' values under development now, please wait.."
	k.SetValue("xxx")

	b := bytes.Buffer{}
	f.WriteTo(&b)
	return b.String()
}

func (self *Module) VerifyConfig(value string) error {
	return nil
}
