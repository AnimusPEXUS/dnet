package pm

import (
	"github.com/jinzhu/gorm"

	"bitbucket.org/AnimusPEXUS/dnet/common_types"
)

type Module struct {
}

func NewModule() *Module {
	ret := new(Module)
	return ret
}

func (self *Module) Name() string {
	return "pm"
}

func (self *Module) Title() string {
	return "PM"
}

func (self *Module) Description() string {
	return "Simple Private Messaging Application for DNet"
}

func (self *Module) PreferredPort() uint64 {
	return 27510 // little endian from "vk"
}

func (self *Module) Instance(db *gorm.DB) (
	common_types.ApplicationInstance,
	error,
) {
	return nil, nil
}
