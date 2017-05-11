package builtin_owntlscert

import (
	"fmt"

	"github.com/AnimusPEXUS/dnet/common_types"
)

type Module struct {
	name *common_types.ModuleName
}

func (self *Module) Name() *common_types.ModuleName {
	if self.name == nil {
		t, err := common_types.ModuleNameNew("builtin_owntlscert")
		if err != nil {
			panic("this shold not been happen")
		}
		self.name = t
	}

	return self.name
}

func (self *Module) Title() string {
	return "Your Own TLS Certificate Editor"
}

func (self *Module) Description() string {
	return "Use this to create or load existing TLS Certificate" +
		" to use with untrusted networks"
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
	ret.com = com
	ret.db = &DB{db: com.GetDBConnection()}
	ret.mod = self

	if !ret.db.db.HasTable(&OwnData{}) {
		if err := ret.db.db.CreateTable(&OwnData{}).Error; err != nil {
			fmt.Println("builtin_owntlscert:", "Can't create table:", err.Error())
		}
	}

	return ret, nil
}
