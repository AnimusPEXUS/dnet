package builtin_net_ip

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/AnimusPEXUS/dnet/common_types"
	"github.com/AnimusPEXUS/worker"
)

type UDPBeacon struct {
	*worker.Worker

	instance *Instance
}

func UDPBeaconNew(instance *Instance) *UDPBeacon {
	ret := new(UDPBeacon)

	ret.instance = instance

	ret.Worker = worker.New(ret.threadWorker)

	return ret
}

func (self *UDPBeacon) threadWorker(
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

	set_working()

	for !is_stop_flag() {

		// TODO: probably it's better to specify something more specific for
		//			 second parameter, but looks like it is works well enough with nil.
		conn, err := net.DialUDP("udp", nil, addr)
		if err != nil {
			//TODO: should it bailout on error?
			fmt.Println("error Dialing UDP:", err)
			continue
		}
		defer conn.Close()

		msg_to_write :=
			common_types.RenderUDPBeaconMessage(self.instance.UDPBeaconMessage())

		if len(msg_to_write) > 1024 {
			set_error(
				errors.New(
					"rendered beacon message exceeds sane maximum length",
				),
			)
			return
		}

		_, err = conn.Write(msg_to_write)
		//sent_len, err := conn.Write(msg_to_write)
		// TODO: maybe need something to do with sent_len
		if err != nil {
			//TODO: should it bailout on error?
			fmt.Println("error sending UDP message:", err)
			continue
		}

		time.Sleep(self.instance.UDPBeaconSleepTime())

	}

}
