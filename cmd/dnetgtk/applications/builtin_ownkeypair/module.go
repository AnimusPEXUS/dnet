package builtin_ownkeypair

import (
	"fmt"

	"github.com/AnimusPEXUS/dnet/common_types"
)

type Module struct {
	name *common_types.ModuleName
}

func (self *Module) Name() *common_types.ModuleName {
	if self.name == nil {
		t, err := common_types.ModuleNameNew("builtin_ownkeypair")
		if err != nil {
			panic("this shold not been happen")
		}
		self.name = t
	}

	return self.name
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
	ret.com = com
	ret.db = &DB{db: com.GetDBConnection()}
	ret.mod = self

	if !ret.db.db.HasTable(&OwnData{}) {
		if err := ret.db.db.CreateTable(&OwnData{}).Error; err != nil {
			fmt.Println("builtin_ownkeypair:", "Can't create table:", err.Error())
		}
	}

	return ret, nil
}
