package tcpip

import (
	"bytes"
	"net"

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
	return "TCP/IP"
}

func (self *Module) Description() string {
	return "DNet module for working with TCP/IP networks using TLS encryption"
}

func (self *Module) WorkingName() string {
	return "tcpip"
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

func (self *Module) NetworkInterfaceList() ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return []string{}, err
	}
	ret := []string{}
	for _, i := range ifaces {
		ret = append(ret, i.Name)
	}
	return ret, nil
}

func (self *Module) NetworkTypeList(
	interface_name string,
) []string {
	return []string{}
}

func (self *Module) NetworkList(
	interface_name string,
	network_type_name string,
) []string {
	return []string{}
}

func (self *Module) Instance(
	network_type string,
	network string,
	config string,
) (common_types.NetworkInstance, error) {

	ret := NewInstance()

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
		"depending on bind_by value, set this to interface name" +
			" (e.g. lo, eth0,enp0s31f6 etc.)" +
			" or to ip address. special values under development now, please wait.."
	k.SetValue("xxx")

	b := bytes.Buffer{}
	f.WriteTo(&b)
	return b.String()
}
