package builtin_net

import (
	//	"fmt"

	//"github.com/AnimusPEXUS/dnet/cmd/dnetgtk/applications/builtin_net_ip"
	"sync"

	"github.com/AnimusPEXUS/dnet/common_types"
)

type Module struct {
	name *common_types.ModuleName

	//network_modules map[string]common_types.ApplicationModule
}

func (self *Module) Name() *common_types.ModuleName {
	if self.name == nil {
		t, err := common_types.ModuleNameNew("builtin_net")
		if err != nil {
			panic("this shold not been happen")
		}
		self.name = t
	}

	return self.name
}

func (self *Module) Title() string {
	return "net"
}

func (self *Module) Description() string {
	return ("Centeral Network communications module. It can be start/stopped" +
		" not only in applications window, but also directly with centaral" +
		" `Online' button (which is positioned on the main GtkDNet window)")
}

func (self *Module) DependsOn() []string {
	return []string{}
}

func (self *Module) IsWorker() bool {
	return true
}

func (self *Module) IsNetwork() bool {
	return true
}

func (self *Module) HaveUI() bool {
	return true
}

func (self *Module) Instance(com common_types.ApplicationCommunicator) (
	common_types.ApplicationModuleInstance,
	error,
) {
	ret := &Instance{}
	ret.com = com
	ret.db = &DB{db: com.GetDBConnection()}
	ret.mod = self
	ret.window_show_sync = new(sync.Mutex)

	ret.module_instances = make(
		map[string]common_types.ApplicationModuleInstance,
	)

	/*
		// this is not a good time to populate ret.module_instances,
		// as DNet may be didn't yet loaded all modules

		if mod, ok := ret.com.GetOtherModuleInstance("builtin_net_ip"); !ok {
			panic("error: module builtin_net_ip is required")
		} else {
			ret.module_instances = append(
				ret.module_instances,
				mod,
			)
		}
	*/

	return ret, nil
}
