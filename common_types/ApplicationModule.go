package common_types

import (
	"net"
	"regexp"
)

func IsApplicationNameCorrect(text string) bool {
	if ok, err :=
		regexp.Match(`^[a-z][a-z0-9_\-]*$`, []byte(text)); err != nil {
		panic(err.Error)
	} else {
		return ok
	}
}

type ApplicationModule interface {
	Name() *ModuleName

	Title() string
	Description() string

	DependsOn() []string // module names which required to be enabled

	// If instance can be called to show it's window
	HasWindow() bool

	//////////////////

	// Single ApplicationInstance should serve all and any requests to it.
	// DNet does not creates separate instances to each request.

	// DNet automatically creates and passes DB connection along with
	// ApplicationCommunicator structure

	// Database connection should be considered sqlcipher (or sqlite3, while
	// developping) GORM instance.

	// DNet uses  Key and ReKey sqlcipher's commands by it's own means and only 
	// DNet should know and only DNet may change DB keyphrase. Application shold 
	// work with db as with regular GORM sqlite3 connection, except Application 
	// should not perform closing of DB. otherwise DB misconsistencies may 
	// happen, leading to database reinitialization or inconsistency (behavior 
	// is not specified).

	// DNet will automatically do ReKey command to DB over some time intervals.
	// presumably 30 days.
	Instance(com ApplicationCommunicator) (ApplicationModuleInstance, error)
}

type ApplicationModuleInstance interface {
	Start()
	Stop()
	Status() *WorkerStatus

	AcceptConn(
		local bool,
		from_local_svc string, // this is meaningfull only if `local' is true
		to_svc string,
		who *Address,
		conn net.Conn,
	) error
}
