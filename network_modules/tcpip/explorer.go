package tcpip

import (
	"bitbucket.org/AnimusPEXUS/dnet/common_types"
)

type Explorer struct {
	status *common_types.WorkerStatus
}

func NewExplorer() *Explorer {
	ret := &Explorer{
		status: common_types.NewWorkerStatus(),
	}
	return ret
}

func (self *Explorer) Status() *common_types.WorkerStatus {
	return self.status
}
