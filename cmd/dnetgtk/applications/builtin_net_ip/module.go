package builtin_net_ip

import (
	//"fmt"
	"sync"

	"github.com/AnimusPEXUS/worker"

	"github.com/AnimusPEXUS/dnet/common_types"
)

var (
	MULTICAST_ADDRESS = "224.0.0.1:5555"
	DESIGNATED_PORT   = 5555
)

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

func (self *Module) IsNetwork() bool {
	return false
}

func (self *Module) HaveUI() bool {
	return true
}

func (self *Module) Instantiate(com common_types.ApplicationCommunicator) (
	common_types.ApplicationModuleInstance,
	error,
) {

	//net_mod, ok := application_net.(*builtin_net.Module)

	//ret, err := InstanceNew(application_net)

	ret := &Instance{}
	ret.com = com
	ret.mod = self
	ret.window_show_sync = new(sync.Mutex)

	ret.db = com.GetDBConnection()
	ret.logger = com.GetLogger()

	ret.tcp_listener = TCPListenerNew(ret)
	ret.udp_beacon = UDPBeaconNew(ret)
	ret.udp_listener = UDPListenerNew(ret)

	ret.cfg = &InstanceConfig{}
	//ret.cfg.SetDefaults()
	ret.LoadConfig()

	//ret.db = &DB{db: com.GetDBConnection()}

	ret.window_show_sync = &sync.Mutex{}

	ret.Worker = worker.New(ret.threadWorker)

	return ret, nil
}
