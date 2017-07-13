package builtin_ownkeypair

import (
	"errors"
	"net"
	"sync"

	"github.com/AnimusPEXUS/dnet/common_types"
	"github.com/AnimusPEXUS/workerstatus"
)

type Instance struct {
	com common_types.ApplicationCommunicator
	db  *DB
	mod *Module

	window *UIWindow

	window_show_sync *sync.Mutex
}

func (self *Instance) Start() {}

func (self *Instance) Stop() {}

func (self *Instance) Status() *workerstatus.WorkerStatus {
	return &workerstatus.WorkerStatus{}
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

func (self *Instance) GetServeConn(calling_app_name string) func(
	bool,
	string,
	string,
	*common_types.Address,
	net.Conn,
) error {
	if calling_app_name != "localDNet" {
		return nil
	}
	return self.ServeConn
}

func (self *Instance) GetSelf(local_svc_name string) (
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

func (self *Instance) GetUI(interface{}) (interface{}, error) {
	self.window_show_sync.Lock()
	defer self.window_show_sync.Unlock()

	var err error

	if self.window == nil {
		self.window, err = UIWindowNew(self)
		if err != nil {
			return nil, errors.New("Error creating window for builtin_ownkeypair module")
		}
		self.window.window.Connect(
			"destroy",
			func() {
				self.window_show_sync.Lock()
				defer self.window_show_sync.Unlock()

				self.window = nil
			},
		)
	}

	return self.window, nil
}

func (self *Instance) GetOwnPrivKey() (string, error) {
	return self.db.GetOwnPrivKey()
}
