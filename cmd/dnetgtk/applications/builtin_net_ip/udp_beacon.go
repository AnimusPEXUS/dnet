package builtin_net_ip

import (
	"net"
	"strconv"
	"time"

	"github.com/AnimusPEXUS/dnet/common_types"
	"github.com/AnimusPEXUS/goworker"
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

	is_stop_flag func() bool,

) {

	addr, err := net.ResolveUDPAddr(
		"udp",
		net.JoinHostPort(MULTICAST_IP, strconv.Itoa(self.instance.cfg.UDPPort)),
	)
	if err != nil {
		self.instance.logger.Error("udp_beacon: " + err.Error())
		return
	}

	go func() {
		for !is_stop_flag() {

			time.Sleep(self.instance.UDPBeaconSleepTime())
			// TODO: probably it's better to specify something more specific for
			// second parameter, but looks like it is works well enough with nil.
			conn, err := net.DialUDP("udp", nil, addr)
			if err != nil {
				//TODO: should it bailout on error?
				self.instance.logger.Error(err.Error())
				continue
			}
			defer conn.Close()

			msg_to_write := make([]byte, 0)
			if msg := self.instance.UDPBeaconMessage(); len(msg) != 0 {
				msg_to_write = common_types.RenderUDPBeaconMessage(msg)
			} else {
				continue
			}

			if len(msg_to_write) > 1024 {
				self.instance.logger.Error(
					"rendered beacon message exceeds sane maximum length",
				)
				<-self.Stop()
				return
			}

			_, err = conn.Write(msg_to_write)
			//sent_len, err := conn.Write(msg_to_write)
			// TODO: maybe need something to do with sent_len
			if err != nil {
				self.instance.logger.Error("error sending UDP message:" + err.Error())
				<-self.Stop()
				return
			}

		}

	}()

	set_working()

	for !is_stop_flag() {
		time.Sleep(time.Second)
	}

}
