package tcpip

import (
	"fmt"
	"net"

	"bitbucket.org/AnimusPEXUS/dnet/common_types"
)

type Instance struct {
	config string

	module common_types.NetworkModule

	listener *Listener
	beacon   *Beacon
	explorer *Explorer
}

func (self *Instance) SetConnAcceptor(func(net.Conn)) {
	return
}

func (self *Instance) Start() {
	fmt.Println("Start() executed")
	return
}

func (self *Instance) Stop() {
	fmt.Println("Stop() executed")
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

func (self *Instance) Config() string {
	return self.config
}

func (self *Instance) SetConfig(value string) error {
	err := self.module.VerifyConfig(value)
	if err == nil {
		self.config = value
	}
	return err
}

func (self *Instance) Beacon() common_types.NetworkBeacon {
	return self.beacon
}

func (self *Instance) Explorer() common_types.NetworkExplorer {
	return self.explorer
}
