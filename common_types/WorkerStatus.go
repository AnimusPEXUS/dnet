package common_types

type WorkerStatus struct {
	Starting bool
	Stopping bool
	Working  bool
}

func NewWorkerStatus() *WorkerStatus {
	ret := &WorkerStatus{
		Starting: false,
		Stopping: false,
		Working:  false,
	}
	return ret
}

func (self *WorkerStatus) Stopped() bool {
	return !self.Starting && !self.Stopping && !self.Working
}

func (self *WorkerStatus) String() string {

	if self.Starting && self.Stopping {
		return "invalid starting and stopping"

	} else if (self.Starting || self.Stopping) && self.Working {
		return "invalid (starting or stopping) and working"

	} else if self.Starting {
		return "starting"

	} else if self.Stopping {
		return "stopping"

	} else if self.Working {
		return "working"

	} else if self.Stopped() {
		return "stopped"

	} else {
		return "unknown"
	}

}
