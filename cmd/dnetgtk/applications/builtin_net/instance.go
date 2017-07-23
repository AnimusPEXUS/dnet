package builtin_net

import (
	"errors"
	"net"
	"sync"
	"time"

	//"github.com/AnimusPEXUS/dnet/cmd/dnetgtk/applications/builtin_net_ip"
	"github.com/AnimusPEXUS/dnet/common_types"
	"github.com/AnimusPEXUS/worker"
	"github.com/AnimusPEXUS/workerstatus"
)

type Instance struct {
	*worker.Worker
	com common_types.ApplicationCommunicator
	db  *DB
	mod *Module

	w *workerstatus.WorkerStatus

	module_instances map[string]common_types.ApplicationModuleInstance

	//window           *UIWindow
	window_show_sync *sync.Mutex
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

func (self *Instance) GetUI(interface{}) (interface{}, error) {
	return nil, errors.New("not implimented")
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

func (self *Instance) GetSelf(calling_app_name string) (
	common_types.ApplicationModuleInstance,
	common_types.ApplicationModule,
	error,
) {
	if calling_app_name == "localDNet" {
		return self, self.mod, nil
	}

	return nil, nil, errors.New("not allowed")
}

func (self *Instance) GetServeConn(calling_app_name string) func(
	local bool,
	calling_app_name string,
	to_svc string,
	who *common_types.Address,
	conn net.Conn,
) error {
	if calling_app_name != "localDNet" {
		return nil
	}
	return self.ServeConn
}

func (self *Instance) threadWorker(

	set_starting func(),
	set_working func(),
	set_stopping func(),
	set_stopped func(),

	is_stop_flag func() bool,

	data interface{},

) {

	for !is_stop_flag() {
		time.Sleep(time.Second)
	}
}
