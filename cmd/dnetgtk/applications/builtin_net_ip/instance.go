package builtin_net_ip

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"strconv"
	"sync"
	"time"

	"github.com/AnimusPEXUS/dnet/common_types"
	"github.com/AnimusPEXUS/gologger"
	"github.com/AnimusPEXUS/goworker"
	"github.com/jinzhu/gorm"
)

type Instance struct {
	*worker.Worker
	com    common_types.ApplicationCommunicator
	db     *gorm.DB
	logger gologger.LoggerI
	mod    *Module

	/*
		serveConnectionCB func(
			to_svc string,
			who *common_types.Address,
			conn net.Conn,
		) error
	*/

	tcp_listener *TCPListener
	udp_beacon   *UDPBeacon
	udp_listener *UDPListener

	window           *UIWindow
	window_show_sync *sync.Mutex

	cfg *InstanceConfig
}

func NewInstance(
	mod *Module,
	com common_types.ApplicationCommunicator,
) (*Instance, error) {
	ret := &Instance{}
	ret.com = com
	ret.mod = mod
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

func (self *Instance) Connect(
	address common_types.NetworkAddress,
) (*net.Conn, error) {
	return nil, errors.New("not implimented")
}

func (self *Instance) LoadConfig() error {
	t := &Data{}
	if !self.db.HasTable(t) {
		if err := self.db.CreateTable(t).Error; err != nil {
			return err
		}
	}

	ret := new(Data)
	if self.db.Where(
		&Data{ValueName: "config"},
	).Find(ret).Error == nil {
		json.Unmarshal([]byte(ret.Value), self.cfg)
		fmt.Println("LoadConfig result ok", self.cfg)
	} else {
		self.cfg = &InstanceConfig{
			TCPListenerEnabled: false,
			UDPBeaconEnabled:   false,
			UDPListenerEnabled: false,
			TCPListenerPort:    5556,
			UDPPort:            5556,
			UDPBeaconInterval:  60,
		}
	}

	return nil
}

func (self *Instance) SaveConfig() error {

	set := &Data{ValueName: "config"}
	if d, err := json.Marshal(*self.cfg); err != nil {
		return err
	} else {
		set.Value = string(d)
	}
	fmt.Println("self.cfg", self.cfg)
	fmt.Println("set.Value", set.Value)
	if self.db.NewRecord(set) {
		return self.db.Create(set).Error
	} else {
		return self.db.Save(set).Error
	}
	return nil
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
	if calling_app_name != "builtin_net" {
		return nil
	}
	return self.ServeConn
}

func (self *Instance) GetUI(interface{}) (interface{}, error) {
	self.window_show_sync.Lock()
	defer self.window_show_sync.Unlock()

	if self.window == nil {
		self.window = UIWindowNew(self)
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

func (self *Instance) threadWorker(

	set_starting func(),
	set_working func(),
	set_stopping func(),
	set_stopped func(),

	is_stop_flag func() bool,

) {
	defer func() {

		set_stopping()

		if self.tcp_listener != nil {
			self.tcp_listener.Stop()
		}

		if self.udp_beacon != nil {
			self.udp_beacon.Stop()
		}

		if self.udp_listener != nil {
			self.udp_listener.Stop()
		}

	}()

	set_starting()

	if self.tcp_listener != nil {
		if self.cfg.TCPListenerEnabled {
			self.tcp_listener.Start()
		}
	}

	if self.udp_beacon != nil {
		if self.cfg.UDPBeaconEnabled {
			self.udp_beacon.Start()
		}
	}

	if self.udp_listener != nil {
		if self.cfg.UDPListenerEnabled {
			self.udp_listener.Start()
		}
	}

	set_working()

	for !is_stop_flag() {
		time.Sleep(time.Second)
	}

}

func (self *Instance) AcceptTCP(conn net.Conn) {

}

func (self *Instance) IsNetwork() bool {
	return false
}

func (self *Instance) UDPBeaconMessage() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		self.logger.Error(err.Error())
		return ""
	}
	addrsl := make([]string, 0)
	for _, addr := range addrs {
		ip, _, err := net.ParseCIDR(addr.String())
		if err != nil {
			continue
		}
		addrsl = append(
			addrsl,
			net.JoinHostPort(
				ip.String(),
				strconv.Itoa(self.cfg.TCPListenerPort),
			),
		)
	}
	ret, err := json.Marshal(addrsl)
	if err != nil {
		self.logger.Error(err.Error())
		return ""
	}
	return string(ret)
}

func (self *Instance) IncommingUDPBeaconMessage(
	conn *net.UDPAddr,
	value string,
) {
	con := common_types.NetworkAddressNewFromString(value)
	go self.com.PossiblyNodeDiscovered(con)
}

func (self *Instance) UDPBeaconSleepTime() time.Duration {
	// TODO: use/create constants
	ret := 5
	if self.cfg.UDPBeaconInterval < 5 || self.cfg.UDPBeaconInterval > 3600 {
		ret = 60
	} else {
		ret = self.cfg.UDPBeaconInterval
	}
	return time.Duration(ret) * time.Second
}

func (self *Instance) GetSelf(calling_svc_name string) (
	common_types.ApplicationModuleInstance,
	common_types.ApplicationModule,
	error,
) {

	if calling_svc_name == "builtin_net" {
		return self, self.mod, nil
	}

	return nil, nil, errors.New("not allowed")

}

func (self *Instance) GetInnodeRPC(calling_app_name string) (
	*rpc.Client, error,
) {
	return nil, errors.New("not implimented")
}
