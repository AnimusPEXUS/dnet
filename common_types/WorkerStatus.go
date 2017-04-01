package common_types

type WorkerStatus struct {
	starting bool
	stopping bool
	working  bool
}

func NewWorkerStatus() *WorkerStatus {
	ret := &WorkerStatus{
		starting: false,
		stopping: false,
		working:  false,
	}
	return ret
}

func (self *WorkerStatus) Starting() bool {
	return self.starting
}

func (self *WorkerStatus) Stopping() bool {
	return self.stopping
}

func (self *WorkerStatus) Working() bool {
	return self.working
}

func (self *WorkerStatus) Stopped() bool {
	return !self.starting && !self.stopping && !self.working
}

func (self *WorkerStatus) String() string {

	if self.Starting() && self.Stopping() {
		return "invalid starting and stopping"

	} else if (self.Starting() && self.Stopping()) || self.Working() {
		return "invalid (starting or stopping) and working"

	} else if self.Starting() {
		return "starting"

	} else if self.Stopping() {
		return "stopping"

	} else if self.Working() {
		return "working"

	} else if self.Stopped() {
		return "stopped"

	} else {
		return "unknown"
	}

}
