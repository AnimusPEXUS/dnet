package pm

import (
	"net"

	"bitbucket.org/AnimusPEXUS/dnet/common_types"
)

type Instance struct {
}

func (self *Instance) SetCoonectionRequestCB(
	func(credentials string) (net.Conn, error),
) {
	return
}

func (self *Instance) Start() {
}

func (self *Instance) Stop() {
}

func (self *Instance) Status() *common_types.ApplicationInstanceStatus {
	return new(common_types.ApplicationInstanceStatus)
}

func (self *Instance) AcceptConn(net.Conn) error {
	return nil
}
