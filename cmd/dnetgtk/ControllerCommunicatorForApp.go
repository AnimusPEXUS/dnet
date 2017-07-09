package main

import (
	"errors"
	"fmt"
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

	caller_name := self.wrap.Name.Value()

	/*
		fmt.Printf(
			"application `%s' tries to get `%s''s instance",
			caller_name,
			name,
		)
	*/

	for _, i := range self.controller.application_presets {
		if i.Name.Value() == name {
			//	fmt.Println("  success")
			return i.Instance.RequestInstance(caller_name)
		}
	}
	//fmt.Println("  failure")
	return nil, nil, errors.New("module not found")
}

func (self *ControllerCommunicatorForApp) ServeConnection(
	to_service string,
	who *common_types.Address,
	conn net.Conn,
) error {

	caller_name := self.wrap.Name.Value()

	if caller_name != "builtin_net" {
		fmt.Printf(
			"module %s tried to access it's communicator's "+
				"ServeConnection() method\n",
			caller_name,
		)
		return errors.New("only `builtin_net' module may access this method")
	}

	return self.controller.ServeConnection(to_service, who, conn)
}
