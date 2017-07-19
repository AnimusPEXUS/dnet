package builtin_net_ip

import (
	"net"

	"github.com/AnimusPEXUS/worker"
)

type TCPWorker struct {
	*worker.Worker

	instance *Instance

	listener *net.TCPListener
}

func TCPWorkerNew(instance *Instance) *TCPWorker {
	ret := new(TCPWorker)

	ret.Worker = worker.New(ret.threadWorker)

	ret.instance = instance

	return ret
}

func (self *TCPWorker) threadWorker(
	set_starting func(),
	set_working func(),
	set_stopping func(),
	set_stopped func(),

	set_error func(error),

	is_stop_flag func() bool,

	defer_me func(),

	data interface{},

) {
	defer defer_me()

	set_starting()

	laddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:5555")
	if err != nil {
		set_error(err)
		return
	}

	set_working()

	if res, err := net.ListenTCP("tcp", laddr); err != nil {
		set_error(err)
		return
	} else {
		self.listener = res
	}

	for !is_stop_flag() {
		conn, err := self.listener.AcceptTCP()
		go self.instance.AcceptTCP(conn, err)
	}

}
