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
