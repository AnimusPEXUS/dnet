package builtin_ownkeypair

import (
	"net"

	"github.com/AnimusPEXUS/dnet/common_types"
)

type Module struct {
}

func (self *Module) Name() *common_types.ModuleName {
	ret, err := common_types.ModuleNameNew("builtin_ownkeypair")
	if err != nil {
		panic("this shold not been happen")
	}
	return ret
}

func (self *Module) Title() string {
	return "Your Own Key Pair Editor"
}

func (self *Module) Description() string {
	return "Use this to create or load existing key pair to use as Your identity"
}

func (self *Module) DependsOn() []string {
	return []string{}
}

func (self *Module) HasWindow() bool {
	return true
}

func (self *Module) Instance(com common_types.ApplicationCommunicator) (
	common_types.ApplicationModuleInstance,
	error,
) {
	ret := &Instance{}
	return ret, nil
}

type Instance struct {
}

func (self *Instance) Start() {
}

func (self *Instance) Stop() {
}

func (self *Instance) Status() *common_types.WorkerStatus {
	return &WorkerStatus{}
}

func (self *Instance) AcceptConn(
	local bool,
	local_svc_name string,
	to_svc string,
	who *common_types.Address,
	conn net.Conn,
) error {
	if !local {
		return errors.New("this module does not accepts external connections")
	}

	return nil 
}

func (self *Instance) RequestInstance(local_svc_name string) (
	common_types.ApplicationModuleInstance,
	common_types.ApplicationModule,
	error,
) {
}

func (self *Instance) ShowWindow() error {
}

func test(a common_types.ApplicationModule) {
	test(&Module{})
}
