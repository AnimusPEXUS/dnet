package builtin_ownkeypair

import (
	"errors"
	"net"
	"net/rpc"
	"sync"

	"github.com/AnimusPEXUS/dnet/common_types"
	"github.com/AnimusPEXUS/goworker"
)

type Instance struct {
	*worker.Worker
	com common_types.ApplicationCommunicator
	mod *Module
	db  *DB

	window *UIWindow

	window_show_sync *sync.Mutex
}

func NewInstance(
	mod *Module,
	com common_types.ApplicationCommunicator,
) (*Instance, error) {
	ret := &Instance{}
	ret.com = com
	ret.db = &DB{db: com.GetDBConnection()}
	ret.mod = mod
	ret.window_show_sync = new(sync.Mutex)

	t := &OwnData{}
	if !ret.db.db.HasTable(t) {
		if err := ret.db.db.CreateTable(t).Error; err != nil {
			return nil, err
		}
	}

	ret.Worker = worker.New(ret.threadWorker)

	return ret, nil
}

func (self *Instance) Connect(
	address common_types.NetworkAddress,
) (*net.Conn, error) {
	return nil, errors.New("not implimented")
}

func (self *Instance) threadWorker(

	set_starting func(),
	set_working func(),
	set_stopping func(),
	set_stopped func(),

	is_stop_flag func() bool,

) {
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

func (self *Instance) GetInnodeRPC(calling_app_name string) (
	*rpc.Client, error,
) {
	found := false
	for _, i := range []string{"builtin_owntlscert"} {
		if i == calling_app_name {
			found = true
			break
		}
	}
	if !found {
		return nil, errors.New("not allowed")
	}
	p1, p2 := net.Pipe()
	serv := rpc.NewServer()
	serv.RegisterName("RPC", NewInnodeRPC(self, calling_app_name))
	go serv.ServeConn(p1)
	return rpc.NewClient(p2), nil
}
