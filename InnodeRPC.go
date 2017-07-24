package dnet

type InnodeRPC struct {
	controller       *Controller
	calling_app_name string
}

func NewInnodeRPC(
	controller *Controller,
	calling_app_name string,
) *InnodeRPC {
	ret := new(InnodeRPC)
	ret.controller = controller
	ret.calling_app_name = calling_app_name
	return ret
}

func (self *InnodeRPC) DiscoveredPossibleNodeMayBeProbe(
	arg interface{},
	res interface{},
) error {
	return nil
}
