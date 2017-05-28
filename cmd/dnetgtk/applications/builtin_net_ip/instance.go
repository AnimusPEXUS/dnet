package builtin_net_ip

import (
	"errors"
	"net"

	"github.com/AnimusPEXUS/dnet/common_types"
)

type Instance struct {
	com common_types.ApplicationCommunicator
	db  *DB
	mod *Module

	serveConnectionCB func(
		to_svc string,
		who *common_types.Address,
		conn net.Conn,
	) error

	// win *UIWindow
	//worker *NetworkWorker

	// ip_module common_types.NetworkApplicationModule

	tcp_worker *WorkerTCP
	udp_worker *WorkerUDP
}

func (self *Instance) Start() {
}

func (self *Instance) Stop() {
}

func (self *Instance) Status() *common_types.WorkerStatus {
	ret := new(common_types.WorkerStatus)

	ret.Starting = self.tcp_worker.w.Starting || self.udp_worker.w.Starting
	ret.Stopping = self.tcp_worker.w.Stopping || self.udp_worker.w.Stopping
	ret.Working = self.tcp_worker.w.Working && self.udp_worker.w.Working

	return ret
}

func (self *Instance) ServeConn(
	local bool,
	local_svc_name string,
	to_svc string,
	who *common_types.Address,
	conn net.Conn,
) error {
	if !local {
		return errors.New("this module does not accepts external connections")
	}

	return nil
}

func (self *Instance) RequestInstance(local_svc_name string) (
	common_types.ApplicationModuleInstance,
	common_types.ApplicationModule,
	error,
) {
	for _, i := range []string{"builtin_owntlscert"} {
		if local_svc_name == i {
			return self, self.mod, nil
		}
	}
	return nil, nil, errors.New("access denied")
}

func (self *Instance) ShowUI() error {
	return errors.New("not implimented")
}

func (self *Instance) Connect(
	to_who *common_types.Address,
	as_service string,
	to_service string,
) (
	*net.Conn,
	error,
) {
	return nil, nil
}
