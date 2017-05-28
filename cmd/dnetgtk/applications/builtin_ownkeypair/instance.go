package builtin_ownkeypair

import (
	"errors"
	"net"

	"github.com/AnimusPEXUS/dnet/common_types"
)

type Instance struct {
	com common_types.ApplicationCommunicator
	db  *DB
	mod *Module

	win *UIWindow
}

func (self *Instance) ServeConn(
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
	for _, i := range []string{"builtin_owntlscert"} {
		if local_svc_name == i {
			return self, self.mod, nil
		}
	}
	return nil, nil, errors.New("access denied")
}

func (self *Instance) ShowUI() error {
	if self.win == nil {
		w, err := UIWindowNew(self)
		if err != nil {
			return err
		}
		self.win = w
	}
	self.win.window.ShowAll()
	return nil
}

func (self *Instance) GetOwnPrivKey() (string, error) {
	return self.db.GetOwnPrivKey()
}
