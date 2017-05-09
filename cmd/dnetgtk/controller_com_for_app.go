package main

import (
	"net"

	"github.com/jinzhu/gorm"

	"github.com/AnimusPEXUS/dnet/common_types"
)

type ControllerCommunicatorForApp struct {
	controller *Controller
	wrap       *ControllerApplicationWrap
	db         *gorm.DB // DB access
}

func (self *ControllerCommunicatorForApp) GetDBConnection() *gorm.DB {
	return self.db
}

// returns socket-like connection to local or remote service
func (self *ControllerCommunicatorForApp) Connect(
	to_who *common_types.Address,
	as_service string,
	to_service string,
) (
	*net.Conn,
	error,
) {

	return nil, nil
}

func (self *ControllerCommunicatorForApp) GetOtherApplicationInstance(
	name string,
) (
	ApplicationModuleInstance,
	ApplicationModule,
	error,
) {
	for _, i := range self.controller.application_presets {
		if i.Name.Value() == name {
			return i.Instance.RequestInstance(self.wrap.Name.Value())
		}
	}
	return nil, nil, errors.New("module not found")
}
