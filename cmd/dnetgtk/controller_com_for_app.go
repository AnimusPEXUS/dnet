package main

import (
	"net"

	"github.com/jinzhu/gorm"

	"bitbucket.org/AnimusPEXUS/dnet/common_types"
)

type ControllerCommunicatorForApp struct {
	db *gorm.DB // DB access
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
