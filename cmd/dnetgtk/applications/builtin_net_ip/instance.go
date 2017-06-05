package builtin_net_ip

import (
	"errors"
	"net"
	"sync"
	"time"

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

	tcp_worker  *TCPWorker
	udp_beacon  *UDPBeacon
	udp_locator *UDPLocator

	w *common_types.WorkerStatus

	stop_flag bool
	err       error

	start_stop_mutex *sync.Mutex
}

func (self *Instance) Start() {
	go func() {
		self.start_stop_mutex.Lock()
		defer self.start_stop_mutex.Unlock()

		if self.w.Stopped() {
			self.stop_flag = false
			go self.threadWorker()
		}
	}()
}

func (self *Instance) Stop() {
	go func() {
		self.start_stop_mutex.Lock()
		defer self.start_stop_mutex.Unlock()

		self.stop_flag = true
	}()
}

func (self *Instance) Status() *common_types.WorkerStatus {
	return self.w
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

func (self *Instance) threadWorker() {
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
