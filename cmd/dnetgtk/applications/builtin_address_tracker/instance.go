package builtin_address_tracker

import (
	"errors"
	"net"
	"net/rpc"
	"sync"
	"time"

	"github.com/AnimusPEXUS/dnet"
	"github.com/AnimusPEXUS/dnet/common_types"
	"github.com/AnimusPEXUS/goworker"
)

type Instance struct {
	*worker.Worker

	com common_types.ApplicationCommunicator
	mod *Module

	db *DB

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

	t := &AddressHistory{}
	if !ret.db.db.HasTable(t) {
		if err := ret.db.db.CreateTable(t).Error; err != nil {
			return nil, err
		}
	}

	ret.Worker = worker.New(ret.threadWorker)
	return ret, nil
}

func (self *Instance) threadWorker(

	set_starting func(),
	set_working func(),
	set_stopping func(),
	set_stopped func(),

	is_stop_flag func() bool,

) {
	for !is_stop_flag() {
		time.Sleep(time.Second)
	}

}

func (self *Instance) Connect(
	address common_types.NetworkAddress,
) (*net.Conn, error) {
	return nil, errors.New("not implimented")
}

func (self *Instance) GetUI(interface{}) (interface{}, error) {
	return nil, errors.New("not implimented")
}

func (self *Instance) ServeConn(
	local bool,
	calling_app_name string, // this is meaningfull only if `local' is true
	to_svc string,
	who *common_types.Address,
	conn net.Conn,
) error {
	return errors.New("not implimented")
}

func (self *Instance) GetInnodeRPC(calling_app_name string) (
	*rpc.Client, error,
) {
	found := false
	for _, i := range []string{
		dnet.DNET_UNIVERSAL_APPLICATION_NAME,
		"builtin_net_ip",
	} {
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
