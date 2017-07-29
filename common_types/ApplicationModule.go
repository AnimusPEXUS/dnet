package common_types

import (
	"net"
	"net/rpc"
	"regexp"

	"github.com/AnimusPEXUS/workerstatus"
)

type (
	ApplicationModuleMap         map[string]ApplicationModule
	ApplicationModuleInstanceMap map[string]ApplicationModuleInstance
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

	// If calling instance's Start() Stop() and Status() methods have sence
	IsWorker() bool

	IsNetwork() bool

	// If instance can be called to show it's window
	HaveUI() bool

	//////////////////

	// Single ApplicationInstance should serve all and any requests to it. DNet
	// does not creates separate instances to each request.

	// DNet automatically creates and passes DB connection along with
	// ApplicationCommunicator structure

	// Database connection should be considered sqlcipher (or sqlite3, while
	// developping) GORM instance.

	// DNet uses  Key and ReKey sqlcipher's commands by it's own means and only
	// DNet should know and only DNet may change DB keyphrase. Application shold
	// work with db as with regular GORM sqlite3 connection, except Application
	// should not perform closing of DB. otherwise DB misconsistencies may happen,
	// leading to database reinitialization or other unspecified behavior.

	// DNet will not do automatic rekeys to not try to gues better time. Instead,
	// UI should provide user with flexible tools for user to decide on he's own
	// when he's need to do rekeys
	Instantiate(com ApplicationCommunicator) (ApplicationModuleInstance, error)
}

type ApplicationModuleInstance interface {

	// Those may be called only if Module's IsWorker() returns true.
	// If Module's IsWorker() returns false, instance have simply
	// impliment them as noop.
	Start() chan bool
	Stop() chan bool
	Restart() chan bool
	Status() *workerstatus.WorkerStatus

	// This is for modules with isNetwork() == true. Controller calling this
	// than wants to connect to some particular node in this network.
	Connect(
		address NetworkAddress,
	) (*net.Conn, error)

	ServeConn(
		local bool,
		calling_app_name string, // this is meaningfull only if `local' is true
		to_svc string,
		who *Address,
		conn net.Conn,
	) error

	// // for usage via ApplicationCommunicator
	// GetServeConn(calling_app_name string) func(
	// 	bool, string, string, *Address, net.Conn,
	// ) error

	// this method may be called only by local services.
	// this method is for direct access for trusted modules (services), as this
	// is mutch faster than socket connection.
	// ApplicationModuleInstance, using calling_svc_name, should decide, to
	// return valid values, or to return nils and error.

	// GetSelf(calling_app_name string) (
	// 	ApplicationModuleInstance,
	// 	ApplicationModule,
	// 	error,
	// )

	// should return module's ui of some sort. If module has no window and it's
	// HaveUI() result's to false, then GetUI() should return non-nil error
	// stating so. anyway, DNet Implimenting Software should not allow user to
	// call GetUI() if HaveUI() results to false
	GetUI(interface{}) (interface{}, error)

	GetInnodeRPC(calling_app_name string) (*rpc.Client, error)
}
