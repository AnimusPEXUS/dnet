package builtin_net_ip

import (
	"net"

	"github.com/AnimusPEXUS/dnet/common_types"
)

type WorkerTCP struct {
	w *common_types.WorkerStatus

	listener *net.TCPConn
}

func WorkerTCPNew() (*WorkerTCP, error) {
	ret := new(WorkerTCP)

	return ret, nil
}

func (self *WorkerTCP) listenerThread() {
	for {

	}
}

//func (self *WorkerTCP) Call() {}
