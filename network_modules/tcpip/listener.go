package tcpip

import (
	"bitbucket.org/AnimusPEXUS/dnet/common_types"
)

type Listener struct {
	status *common_types.WorkerStatus
}

func NewListener() *Listener {
	ret := new(Listener)
	return ret
}

func (self *Listener) Start() error {
	return nil
}

func (self *Listener) Stop() {}

func (self *Listener) Status() *common_types.WorkerStatus {
	return common_types.NewWorkerStatus()
}
