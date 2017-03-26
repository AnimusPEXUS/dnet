package common_types

import (
	"net"
)

type NetworkModule interface {
	Name() string
	Description() string

	WorkingName() string

	SampleConfig() string

	CanExplore() bool
	CanBeacon() bool
	CanListen() bool

	NetworkInterfaceList() ([]string, error)
	NetworkTypeList(interface_name string) []string
	NetworkList(interface_name, network_type_name string) []string

	Instance(network_type, network, settings string) (NetworkInstance, error)
}

type NetworkInstance interface {
	Stop()
	Status() *NetworkInstanceStatus

	Listen(
		func(
			*NetworkModule,
			*NetworkInstance,
			*net.Conn,
		),
	) NetworkListener

	Settings() string
	SetSettings(string) error

	Beacon() NetworkBeacon
	Explorer() NetworkExplorer
}

type NetworkListener interface {
	Start() error
	Stop()
	Status() *WorkerStatus
}

type NetworkBeacon interface {
	Status() *WorkerStatus
	Ping()
}

type NetworkExplorer interface {
	Status() *WorkerStatus
	// ExplorePeers()
}

type NetworkInstanceStatus struct {
	WorkerStatus
	sent_bytes  uint64
	recvd_bytes uint64
	is_online   bool
}

func NewNetworkInstanceStatus() *NetworkInstanceStatus {
	return new(NetworkInstanceStatus)
}

func (self *NetworkInstanceStatus) SentBytes() uint64 {
	return self.sent_bytes
}

func (self *NetworkInstanceStatus) RecvdBytes() uint64 {
	return self.recvd_bytes
}

func (self *NetworkInstanceStatus) IsOnline() bool {
	return self.is_online
}

