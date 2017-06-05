package builtin_net_ip

import (
	//"fmt"

	"github.com/AnimusPEXUS/dnet/common_types"
)

var MULTICAST_ADDRESS = "224.0.0.1:5555"
var DESIGNATED_PORT = 5555

type Module struct {
	name *common_types.ModuleName
}

func (self *Module) Name() *common_types.ModuleName {
	if self.name == nil {
		t, err := common_types.ModuleNameNew("builtin_net_ip")
		if err != nil {
			panic("this shold not been happen")
		}
		self.name = t
	}

	return self.name
}

func (self *Module) Title() string {
	return "net_ip"
}

func (self *Module) Description() string {
	return ("IP transport mechanism")
}

func (self *Module) DependsOn() []string {
	return []string{}
}

func (self *Module) IsWorker() bool {
	return true
}

func (self *Module) HaveUI() bool {
	return true
}

func (self *Module) Instance() (
	common_types.NetworkModuleInstance,
	error,
) {

	//net_mod, ok := application_net.(*builtin_net.Module)

	//ret, err := InstanceNew(application_net)

	ret := &Instance{}
	//ret.com = com
	//ret.db = &DB{db: com.GetDBConnection()}
	//ret.mod = self

	return ret, nil
}
