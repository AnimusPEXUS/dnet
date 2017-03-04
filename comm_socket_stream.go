package dnet

import (
	"errors"
	"net"
	"sync"
)

type CommSocketStreamServerPool struct {
	Pool           []*CommSocketStreamServer
	dnetcontroller *DNetController
}

func NewCommSocketStreamServerPool(
	dnetcontroller *DNetController,
) *CommSocketStreamServerPool {

	ret := new(CommSocketStreamServerPool)
	ret.dnetcontroller = dnetcontroller

	return ret
}

func (self *CommSocketStreamServerPool) Stop() {
	for _, i := range self.Pool {
		i.Stop()
	}
}

func (self *CommSocketStreamServerPool) IsStoped() bool {
	for _, i := range self.Pool {
		if i.GetStatus() != "stopped" {
			return false
		}
	}
	return true
}

type CommSocketStreamServer struct {
	listener net.Listener
	mut      *sync.Mutex

	net, laddr string

	starting bool
	working  bool
	stopping bool
}

func NewCommSocketStreamServer(
	net, laddr string,
	rpc_comm *DNetController,
) *CommSocketStreamServer {
	ret := new(CommSocketStreamServer)
	ret.mut = new(sync.Mutex)

	ret.net = net
	ret.laddr = laddr
	return ret
}

func (self *CommSocketStreamServer) GetListener() net.Listener {
	return self.listener
}

func (self *CommSocketStreamServer) LocalAddr() net.Addr {
	return self.listener.LocalAddr()
}

func (self *CommSocketStreamServer) Start() error {

	if self.GetStatus() != "stopped" {
		return errors.New("not stopped")
	}

	ret.mut.Lock()
	defer ret.mut.Unlock()

	self.starting = true
	defer func() { self.starting = false }()

	lstnr, err := net.Listen(self.net, self.laddr)
	if err != nil {
		return err
	}

	self.listener = lstnr

	self.working = true

	return nil
}

func (self *CommSocketStreamServer) Stop() error {

	ret.mut.Lock()
	defer ret.mut.Unlock()

	self.stopping = true
	defer func() { self.stopping = false }()

	self.listener.Close()

	self.working = false

	return nil
}

func (self *CommSocketStreamServer) GetStatus() string {
	ret.mut.Lock()
	defer ret.mut.Unlock()
	if self.working && !self.starting && !self.stopping {
		return "working"
	} else if !self.working && !self.starting && !self.stopping {
		return "stopped"
	} else if !self.working && self.starting && !self.stopping {
		return "starting"
	} else if !self.working && !self.starting && self.stopping {
		return "stopping"
	} else {
		return "unknown"
	}
}
