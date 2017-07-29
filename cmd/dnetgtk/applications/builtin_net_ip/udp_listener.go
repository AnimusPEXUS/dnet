package builtin_net_ip

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/AnimusPEXUS/dnet/common_types"
	"github.com/AnimusPEXUS/goworker"
)

type UDPListener struct {
	*worker.Worker

	instance *Instance
}

func UDPListenerNew(instance *Instance) *UDPListener {
	ret := new(UDPListener)

	ret.instance = instance

	ret.Worker = worker.New(ret.threadWorker)

	return ret
}

func (self *UDPListener) threadWorker(
	set_starting func(),
	set_working func(),
	set_stopping func(),
	set_stopped func(),

	is_stop_flag func() bool,

) {

	addr, err := net.ResolveUDPAddr(
		"udp",
		net.JoinHostPort(MULTICAST_IP, strconv.Itoa(self.instance.cfg.UDPPort)),
	)
	if err != nil {
		self.instance.logger.Error(err.Error())
		return
	}

	listener, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		self.instance.logger.Error("error making listener: " + err.Error())
		return
	}
	defer listener.Close()

	go func() {
		defer self.Stop()
		for !is_stop_flag() {
			accept_buff := make([]byte, 1024)

			for i := 0; i != len(accept_buff); i++ {
				accept_buff[i] = 0
			}

			length, addr, err := listener.ReadFromUDP(accept_buff)
			if err != nil {
				// NOTE: probably not every error should lead to thread termination
				self.instance.logger.Error("error reading UDP message:" + err.Error())
				continue
			}

			if length == 1024 {
				self.instance.logger.Error(
					"error reading UDP message: message too long",
				)
				return
			}

			parsed, err := common_types.ParseUDPBeaconMessage(accept_buff)
			if err != nil {
				self.instance.logger.Error(err.Error())
				continue
			}

			self.instance.logger.Info(
				fmt.Sprintf(
					"accepted DNet UDP beacon message: '%s' from %v",
					parsed,
					addr,
				),
			)

			go self.instance.IncommingUDPBeaconMessage(addr, parsed)
		}

	}()

	set_working()

	for !is_stop_flag() {
		time.Sleep(time.Second)
	}
}
