package builtin_net_ip

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/AnimusPEXUS/dnet/common_types"
)

type UDPBeacon struct {
	w *common_types.WorkerStatus

	instance *Instance

	stop_flag bool
	err       error

	start_stop_mutex *sync.Mutex
}

func UDPBeaconNew(instance *Instance) (*UDPBeacon, error) {
	ret := new(UDPBeacon)

	ret.instance = instance
	ret.w = common_types.NewWorkerStatus()
	ret.err = nil

	return ret, nil
}

func (self *UDPBeacon) Error() error {
	return self.err
}

func (self *UDPBeacon) Start() {
	go func() {
		self.start_stop_mutex.Lock()
		defer self.start_stop_mutex.Unlock()

		if self.w.Stopped() {
			self.stop_flag = false
			go self.threadWorker()
		}
	}()
}

func (self *UDPBeacon) Stop() {
	go func() {
		self.start_stop_mutex.Lock()
		defer self.start_stop_mutex.Unlock()

		self.stop_flag = true
	}()
}

func (self *UDPBeacon) threadWorker() {

	defer self.w.Reset()

	self.w.Starting = true

	addr, err := net.ResolveUDPAddr("udp", MULTICAST_ADDRESS)
	if err != nil {
		self.err = err
		return
	}

	self.w.Working = true
	self.w.Starting = false

	for !self.stop_flag {

		// TODO: probably it's better to specify something more specific for
		//			 second parameter, but looks like it is works well enough with nil.
		conn, err := net.DialUDP("udp", nil, addr)
		if err != nil {
			//TODO: should it bailout on error?
			fmt.Println("error Dialing UDP:", err)
			continue
		}
		defer conn.Close()

		msg_to_write :=
			common_types.RenderUDPBeaconMessage(self.instance.UDPBeaconMessage())

		if len(msg_to_write) > 1024 {
			self.err = errors.New(
				"rendered beacon message exceeds sane maximum length",
			)
			return
		}

		_, err = conn.Write(msg_to_write)
		//sent_len, err := conn.Write(msg_to_write)
		// TODO: maybe need something to do with sent_len
		if err != nil {
			//TODO: should it bailout on error?
			fmt.Println("error sending UDP message:", err)
			continue
		}

		time.Sleep(self.instance.UDPBeaconSleepTime())

	}

}
