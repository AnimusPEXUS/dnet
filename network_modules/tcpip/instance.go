package tcpip

import (
	"net"

	"bitbucket.org/AnimusPEXUS/dnet/common_types"
)

type Instance struct {
	listener *Listener
	beacon   *Beacon
	explorer *Explorer
}

func NewInstance() *Instance {
	return new(Instance)
}

func (self *Instance) Stop() {
	return
}

func (self *Instance) Status() *common_types.NetworkInstanceStatus {
	return common_types.NewNetworkInstanceStatus()
}

func (self *Instance) Listen(
	func(
		*common_types.NetworkModule,
		*common_types.NetworkInstance,
		*net.Conn,
	),
) common_types.NetworkListener {
	return self.listener
}

func (self *Instance) Settings() string {
	return ""
}

func (self *Instance) SetSettings(string) error {
	return nil
}

func (self *Instance) Beacon() common_types.NetworkBeacon {
	return self.beacon
}

func (self *Instance) Explorer() common_types.NetworkExplorer {
	return self.explorer
}
