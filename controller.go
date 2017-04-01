package dnet

type Controller struct {
}

func NewController() (*Controller, error) {
	ret := new(Controller)
	return ret, nil
}

func (self *Controller) Start() error {
	return nil
}

func (self *Controller) Stop() {
}
