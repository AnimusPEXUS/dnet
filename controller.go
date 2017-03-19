package dnet

type Controller struct {
}

func NewController() (*Controller, error) {
	return new(Controller), nil
}

func (self *Controller) Start() error {
	return nil
}

func (self *Controller) Stop() {
}
