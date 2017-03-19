package tcpip

import (
	"bitbucket.org/AnimusPEXUS/dnet/common_types"
)

type Beacon struct {
	status *common_types.WorkerStatus
}

func NewBeeacon() *Beacon {
	ret := &Beacon{
		status: common_types.NewWorkerStatus(),
	}
	return ret
}

func (self *Beacon) Status() *common_types.WorkerStatus {
	return self.status
}

func (self *Beacon) Ping() {
}
