package builtin_net_ip

import (
	"fmt"
	"net"
	"time"

	"github.com/AnimusPEXUS/goworker"
)

type TCPListener struct {
	*worker.Worker

	instance *Instance

	listener *net.TCPListener
}

func TCPListenerNew(instance *Instance) *TCPListener {
	ret := new(TCPListener)

	ret.Worker = worker.New(ret.threadWorker)

	ret.instance = instance

	return ret
}

func (self *TCPListener) threadWorker(
	set_starting func(),
	set_working func(),
	set_stopping func(),
	set_stopped func(),

	is_stop_flag func() bool,

) {

	set_starting()

	laddr, err := net.ResolveTCPAddr(
		"tcp",
		fmt.Sprintf(
			"0.0.0.0:%d",
			self.instance.cfg.TCPListenerPort,
		),
	)
	if err != nil {
		self.instance.logger.Error("tcp_listener: " + err.Error())
		return
	}

	if res, err := net.ListenTCP("tcp", laddr); err != nil {
		self.instance.logger.Error("tcp_listener: " + err.Error())
		return
	} else {
		self.listener = res
		defer self.listener.Close()
	}

	go func() {
		for {
			conn, err := self.listener.AcceptTCP()
			if is_stop_flag() {
				return
			}
			if err != nil {
				self.instance.logger.Error("tcp_listener AcceptTCP(): " + err.Error())
				<-self.Stop()
				return
			}
			go self.instance.AcceptTCP(conn)
		}
	}()

	set_working()

	for !is_stop_flag() {
		time.Sleep(time.Second)
	}

}
