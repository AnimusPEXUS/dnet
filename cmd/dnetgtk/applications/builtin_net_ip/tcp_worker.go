package builtin_net_ip

import (
	"net"
	"sync"

	"github.com/AnimusPEXUS/dnet/common_types"
)

type TCPWorker struct {
	w *common_types.WorkerStatus

	instance *Instance

	listener *net.TCPListener

	stop_flag bool
	err       error

	start_stop_mutex *sync.Mutex
}

func TCPWorkerNew(instance *Instance) (*TCPWorker, error) {
	ret := new(TCPWorker)

	ret.instance = instance
	ret.w = common_types.NewWorkerStatus()
	ret.err = nil

	return ret, nil
}

func (self *TCPWorker) Error() error {
	return self.err
}

func (self *TCPWorker) Start() {
	go func() {
		self.start_stop_mutex.Lock()
		defer self.start_stop_mutex.Unlock()

		if self.w.Stopped() {
			self.stop_flag = false
			go self.threadWorker()
		}
	}()
}

func (self *TCPWorker) Stop() {
	go func() {
		self.start_stop_mutex.Lock()
		defer self.start_stop_mutex.Unlock()

		self.stop_flag = true
	}()
}

func (self *TCPWorker) threadWorker() {

	defer self.w.Reset()

	self.w.Starting = true

	laddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:5555")
	if err != nil {
		self.err = err
		return
	}

	self.w.Working = true
	self.w.Starting = false

	if res, err := net.ListenTCP("tcp", laddr); err != nil {
		self.err = err
		return
	} else {
		self.listener = res
	}

	for !self.stop_flag {
		conn, err := self.listener.AcceptTCP()
		go self.instance.AcceptTCP(conn, err)
	}

}
