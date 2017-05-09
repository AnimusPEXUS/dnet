package common_types

import (
	"net"

	"github.com/jinzhu/gorm"
)

// This communicator is the standard way for application (module)
// to interact with DNet and other applications
type ApplicationCommunicator interface {
	GetDBConnection() *gorm.DB // Application's own db connection

	Connect(
		// depending on Address, DNet will decide if connect local or remote
		to_who *Address,

		// Select own service name
		// in case of remote connection, this value does no sence. but in case of
		// local connection, DNet introduce caller by this name.

		// This value must be one of
		as_service string,

		// select destination service
		to_service string,
	) (
		*net.Conn,
		error,
	)

	GetOtherApplicationInstance(name string) (
		ApplicationModuleInstance,
		ApplicationModule,
		error,
	)
}
