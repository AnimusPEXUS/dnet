package dnet

import "github.com/AnimusPEXUS/dnet/common_types"

type InnodeRPC struct {
	controller       *Controller
	calling_app_name *common_types.ModuleName
}

func NewInnodeRPC(
	controller *Controller,
	calling_app_name *common_types.ModuleName,
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
