package builtin_net

import (
	"fmt"

	"github.com/AnimusPEXUS/dnet/cmd/dnetgtk/applications/builtin_net_ip"
	"github.com/AnimusPEXUS/dnet/common_types"
)

type Module struct {
	name *common_types.ModuleName

	network_modules map[string]common_types.NetworkApplicationModule
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

func (self *Module) HasUI() bool {
	return false
}

func (self *Module) Instance(com common_types.ApplicationCommunicator) (
	common_types.ApplicationModuleInstance,
	error,
) {
	ret := &Instance{}
	ret.com = com
	ret.db = &DB{db: com.GetDBConnection()}
	ret.mod = self

	if !ret.db.db.HasTable(&OwnData{}) {
		if err := ret.db.db.CreateTable(&OwnData{}).Error; err != nil {
			// TODO: this sort of error handeling shold be reworker her and in
			//       ither modules
			fmt.Println("builtin_net:", "Can't create table:", err.Error())
		}
	}

	ret.network_modules["builtin_net_ip"] = &builtin_net_ip.Module().Instance()

	if mod, ok := ret.com.GetOtheModuleInstance("builtin_net_ip"); ok {
		ret.ip_module = mod
	} else {
		return nil, errors.New(
			"builtin_net: module `builtin_net_ip' is required for work",
		)
	}

	return ret, nil
}
