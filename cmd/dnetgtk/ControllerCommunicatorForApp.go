package main

import (
	"errors"
	"fmt"
	"net"

	"github.com/jinzhu/gorm"

	"github.com/AnimusPEXUS/dnet/common_types"
	"github.com/AnimusPEXUS/workerstatus"
)

type ControllerCommunicatorForApp struct {
	name       *common_types.ModuleName
	controller *Controller
	wrap       *SafeApplicationModuleInstanceWrap
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
	common_types.ApplicationModuleInstance,
	common_types.ApplicationModule,
	error,
) {

	/*
		if !self.controller.module_instances_added {
			// NOTE: hitting this panic should be considered programming error
			// TODO: check what this can't be exploited somehow
			panic("programming error: modules not instantiated yet")
			return nil, nil, errors.New("modules not instantiated yet")
		}
	*/

	caller_name := self.name.Value()

	/*
		fmt.Printf(
			"application `%s' tries to get `%s''s instance",
			caller_name,
			name,
		)
	*/

	for key, val := range self.controller.application_controller.application_wrappers {
		if key == name {
			//	fmt.Println("  success")
			if val.Instance == nil {
				return nil, nil, errors.New("not instantiated")
			}

			return val.Instance.GetSelf(caller_name)
		}
	}
	//fmt.Println("  failure")
	return nil, nil, errors.New("module not found")
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

func (self *ControllerCommunicatorForApp) InstanceStatusChanged(
	data *workerstatus.WorkerStatus,
) {
	self.controller.application_controller.ModuleInstanceStatusChangeListener(
		self.name,
		data.String(),
	)
}
