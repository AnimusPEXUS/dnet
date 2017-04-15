package common_types

import (
	"net"

	"github.com/jinzhu/gorm"
)

type ApplicationModule interface {
	// name should be in lower case letters without spaces.
	// re is "^[a-z][a-z0-9_\-*]$". upper case used by special values
	Name() string

	Title() string
	Description() string

	DependsOn() []string // module names which required to be enabled

	// watever instance can be contacted through AcceptConn() call
	// (and accessed by other DNet nodes)
	ListensConnections() bool

	// list of applications on same DNet instance, which allowed to get access
	// to instance of this Application module
	AllowLocalInstanceAccessTo() []string

	// NOTE: instance may be contacted through the network from any service
	//       on any DNet node by AcceptConn, and this is not the same access
	//       level, which is grunted by AllowLocalInstanceAccessTo() list.

	//////////////////

	// Single ApplicationInstance should serve all and any requests to it.
	// DNet does not creates separate instances to each request.

	// DNet creates (if isn't already exists) database connection for starting
	// application instance and passes db connection to it.

	// db should be considered sqlcipher or sqlite3 GORM instance.
	// DNet uses Key and ReKey commands by it's own means and only DNet
	// should know and only DNet may change DB keyphrase. Application shold work
	// with db as with regular GORM sqlite3 connection, except Application should
	// not perform closing of DB.

	// DNet will automatically do ReKey command to DB over some time intervals.
	// presumably 30 days.
	Instance(db *gorm.DB) (ApplicationInstance, error)
}

type ApplicationInstance interface {
	SetCoonectionRequestCB(func(credentials string) (net.Conn, error))

	Start()
	Stop()
	Status() *ApplicationInstanceStatus

	AcceptConn(net.Conn) error
}

type ApplicationInstanceStatus struct {
	WorkerStatus
}
