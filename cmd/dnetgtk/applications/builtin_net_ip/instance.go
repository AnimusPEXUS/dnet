package builtin_net_ip

import (
	"errors"
	"net"
	"sync"
	"time"

	"github.com/AnimusPEXUS/dnet/common_types"
	"github.com/AnimusPEXUS/worker"
)

type Instance struct {
	*worker.Worker
	com common_types.ApplicationCommunicator
	db  *DB
	mod *Module

	/*
		serveConnectionCB func(
			to_svc string,
			who *common_types.Address,
			conn net.Conn,
		) error
	*/

	tcp_worker  *TCPWorker
	udp_beacon  *UDPBeacon
	udp_locator *UDPLocator

	window           *UIWindow
	window_show_sync *sync.Mutex
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

/*
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
*/

func (self *Instance) ShowUI() error {
	self.window_show_sync.Lock()
	defer self.window_show_sync.Unlock()

	if self.window == nil {
		self.window = UIWindowNew(self)
		self.window.window.Connect("destroy", self._OnWindowDestroy)
	}

	self.window.Show()

	return nil
}

func (self *Instance) _OnWindowDestroy() {
	self.window_show_sync.Lock()
	defer self.window_show_sync.Unlock()

	if self.window != nil {
		self.window = nil
	}
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

func (self *Instance) threadWorker(

	set_starting func(),
	set_working func(),
	set_stopping func(),
	set_stopped func(),

	set_error func(error),

	is_stop_flag func() bool,

	data interface{},

) {
}

func (self *Instance) AcceptTCP(conn net.Conn, err error) {

}

func (self *Instance) UDPBeaconMessage() string {
	return "test" // TODO
}

func (self *Instance) IncommingUDPBeaconMessage(
	conn *net.UDPAddr,
	value string) {
}

func (self *Instance) UDPBeaconSleepTime() time.Duration {
	return time.Duration(1 * time.Minute)
}

func (self *Instance) RequestInstance(calling_svc_name string) (
	common_types.ApplicationModuleInstance,
	common_types.ApplicationModule,
	error,
) {

	if calling_svc_name == "builtin_net" {
		return self, self.mod, nil
	}

	return nil, nil, errors.New("not allowed")

}
