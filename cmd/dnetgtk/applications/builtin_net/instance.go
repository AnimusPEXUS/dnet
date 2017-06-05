package builtin_net

import (
	"errors"
	"net"

	//"github.com/AnimusPEXUS/dnet/cmd/dnetgtk/applications/builtin_net_ip"
	"github.com/AnimusPEXUS/dnet/common_types"
)

type Instance struct {
	com common_types.ApplicationCommunicator
	db  *DB
	mod *Module

	w *common_types.WorkerStatus

	module_instances map[string]common_types.NetworkModuleInstance
}

func (self *Instance) Start() {
	go func() {
		for _, value := range self.module_instances {
			go value.Start()
		}
	}()
}

func (self *Instance) Stop() {
	go func() {
		for _, value := range self.module_instances {
			go value.Stop()
		}
	}()
}

func (self *Instance) Status() *common_types.WorkerStatus {
	t := make([]*common_types.WorkerStatus, 0)
	for _, value := range self.module_instances {
		t = append(t, value.Status())
	}
	self.w.Sum(t)
	return self.w
}

func (self *Instance) ServeConn(
	local bool,
	local_svc_name string,
	to_svc string,
	who *common_types.Address,
	conn net.Conn,
) error {
	return errors.New("this module does not accepts external connections")
}

func (self *Instance) RequestInstance(local_svc_name string) (
	common_types.ApplicationModuleInstance,
	common_types.ApplicationModule,
	error,
) {
	return nil, nil, errors.New("any access is denied to this module")
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
