package builtin_net

import (
	"errors"
	"net"

	"github.com/AnimusPEXUS/dnet/cmd/dnetgtk/applications/builtin_net_ip"
	"github.com/AnimusPEXUS/dnet/common_types"
)

type Instance struct {
	com common_types.ApplicationCommunicator
	db  *DB
	mod *Module

	// win *UIWindow
	//worker *NetworkWorker

	module_instances map[string]common_types.NetworkApplicationModuleInstance
}

func (self *Instance) Start() {

	go func() {
		for _, i := range self.mod.network_modules {
			ins := i.Instance()
			self.network_module_instances = append(
				self.network_module_instances,
				ins,
			)
		}
	}()
}

func (self *Instance) Stop() {
	go self.ip_module.Stop()
}

func (self *Instance) Status() *common_types.WorkerStatus {
	return self.ip_module.Status()
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

func (self *Instance) GetOwnPrivKey() (string, error) {
	return self.db.GetOwnPrivKey()
}

func (self *Instance) Connect(
	to_who *common_types.Address,
	as_service string,
	to_service string,
) (
	*net.Conn,
	error,
) {

}

func (self *Instance) ServeConn(
	local bool,
	calling_svc_name string, // this is meaningfull only if `local' is true
	to_svc string,
	who *Address,
	conn net.Conn,
) error {
}
