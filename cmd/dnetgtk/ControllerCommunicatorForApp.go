package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"

	"github.com/jinzhu/gorm"

	"github.com/AnimusPEXUS/dnet/common_types"
	"github.com/AnimusPEXUS/gologger"
)

type ControllerCommunicatorForApp struct {
	name       *common_types.ModuleName
	controller *Controller
	wrap       *SafeApplicationModuleInstanceWrap
	db         *gorm.DB // DB access
	logger     *ControllerCommunicatorForAppLogger
}

func (self *ControllerCommunicatorForApp) GetDBConnection() *gorm.DB {
	return self.db
}

func (self *ControllerCommunicatorForApp) GetLogger() gologger.LoggerI {
	return self.logger
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

func (self *ControllerCommunicatorForApp) ServeConnection(
	who *common_types.Address,
	conn net.Conn,
) error {

	caller_name := self.name.Value()

	if caller_name != "builtin_net" {
		fmt.Printf(
			"module %s tried to access it's communicator's "+
				"ServeConnection() method\n",
			caller_name,
		)
		return errors.New("only `builtin_net' module may access this method")
	}

	self.controller.dnet_controller.ServeConnection(who, conn)

	return nil
}

func (self *ControllerCommunicatorForApp) PossiblyNodeDiscovered(
	address common_types.NetworkAddress,
) error {
	if !self.wrap.Module.IsNetwork() {
		return errors.New(
			"Only network modules allowed to call NodeDiscovered",
		)
	}
	go self.controller.dnet_controller.
		PossiblyNodeDiscoveredNotificationReceptor(
			self.name,
			address,
		)
	return nil
}

func (self *ControllerCommunicatorForApp) GetInnodeRPC(
	target_app_name string,
) (*rpc.Client, error) {
	return self.controller.application_controller.GetInnodeRPC(
		self.name,
		common_types.ModuleNameNewF(target_app_name),
	)
}
