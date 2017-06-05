package common_types

import (
	"net"
)

type NetworkModule interface {
	Name() *ModuleName

	Title() string
	Description() string

	DependsOn() []string

	IsWorker() bool

	HaveUI() bool

	Instance() (
		NetworkModuleInstance,
		error,
	)
}

type NetworkModuleInstance interface {

	/*
			Talking about "serving" connections. possibly in the furure DNet
			will be capable of forwarding connections somehow. But not yet!

			ServeConn(
				local bool,
			calling_svc_name string, // this is meaningfull only if `local' is true
			to_svc string,
			who *Address,
			conn net.Conn,
		) error
	*/

	ShowUI() error

	Start()
	Stop()
	Status() *WorkerStatus

	Connect(
		to_who *Address,
		as_service string,
		to_service string,
	) (
		*net.Conn,
		error,
	)
}
