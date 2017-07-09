package builtin_net_ip

import (
	"errors"
	"fmt"
	"net"

	"github.com/AnimusPEXUS/dnet/common_types"
	"github.com/AnimusPEXUS/worker"
)

type UDPLocator struct {
	*worker.Worker

	instance *Instance
}

func UDPLocatorNew(instance *Instance) (*UDPLocator, error) {
	ret := new(UDPLocator)

	ret.instance = instance

	ret.Worker = worker.New(ret.threadWorker)

	return ret, nil
}

func (self *UDPLocator) threadWorker(
	set_starting func(),
	set_working func(),
	set_stopping func(),
	set_stopped func(),

	set_error func(error),

	is_stop_flag func() bool,

	data interface{},
) {

	set_starting()
	defer set_stopped()

	addr, err := net.ResolveUDPAddr("udp", MULTICAST_ADDRESS)
	if err != nil {
		set_error(err)
		return
	}

	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		set_error(errors.New("error making listener: " + err.Error()))
		return
	}

	defer conn.Close()

	accept_buff := make([]byte, 1024)

	set_working()

	for !is_stop_flag() {
		for i := 0; i != len(accept_buff); i++ {
			accept_buff[i] = 0
		}

		length, addr, err := conn.ReadFromUDP(accept_buff)
		if err != nil {
			set_error(errors.New("error reading UDP message:" + err.Error()))
			// TODO: probably not every error should lead to thread termination
			return
		}

		if length == 1024 {
			set_error(errors.New("error reading UDP message: message too large"))
			return
		}

		parsed, err := common_types.ParseUDPBeaconMessage(accept_buff)
		if err != nil {
			continue
		}

		fmt.Println("accepted DNet UDP beacon message: '%s' from %v", parsed, addr)

		go func() {
			self.instance.IncommingUDPBeaconMessage(addr, parsed)
		}()

	}

}
