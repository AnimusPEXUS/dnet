package builtin_net_ip

import (
	"fmt"
	"net"
	"sync"

	"github.com/AnimusPEXUS/dnet/common_types"
)

type UDPLocator struct {
	w *common_types.WorkerStatus

	instance *Instance

	stop_flag bool
	err       error

	start_stop_mutex *sync.Mutex
}

func UDPLocatorNew(instance *Instance) (*UDPLocator, error) {
	ret := new(UDPLocator)

	ret.instance = instance
	ret.w = common_types.NewWorkerStatus()
	ret.err = nil

	return ret, nil
}

func (self *UDPLocator) Error() error {
	return self.err
}

func (self *UDPLocator) Start() {
	go func() {
		self.start_stop_mutex.Lock()
		defer self.start_stop_mutex.Unlock()

		if self.w.Stopped() {
			self.stop_flag = false
			go self.threadWorker()
		}
	}()
}

func (self *UDPLocator) Stop() {
	go func() {
		self.start_stop_mutex.Lock()
		defer self.start_stop_mutex.Unlock()

		self.stop_flag = true
	}()
}

func (self *UDPLocator) threadWorker() {

	defer self.w.Reset()

	self.w.Starting = true

	addr, err := net.ResolveUDPAddr("udp", MULTICAST_ADDRESS)
	if err != nil {
		self.err = err
		return
	}

	conn, err := net.ListenMulticastUDP("udp", addr)
	if err != nil {
		self.err = errors.New("error making listener: " + err.Error())
		return
	}

	defer conn.Close()

	var accept_buff [1024]byte

	self.w.Working = true
	self.w.Starting = false

	for !stop_flag {
		for i := 0; i != len(accept_buff); i++ {
			accept_buff[i] = 0
		}

		length, addr, err := conn.ReadFromUDP(accept_buff)
		if err != nil {
			self.err = errors.New("error reading UDP message:" + err.Error())
			return // TODO: probably not every error should lead to thread termination
		}

		parsed, err := common_types.ParseUDPBeaconMessage(accept_buff)
		if err != nil {
			continue
		}

		fmt.Println("accepted DNet UDP beacon message: '%s' from %v", parsed, addr)

		go func() {
			self.instance.UDPBeaconSpotted(parsed)
		}()

	}

}
