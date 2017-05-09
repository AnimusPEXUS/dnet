package builtin_ownkeypair

type Instance struct {
}

func (self *Instance) Start() {
}

func (self *Instance) Stop() {
}

func (self *Instance) Status() *common_types.WorkerStatus {
	return &WorkerStatus{}
}

func (self *Instance) AcceptConn(
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

func (self *Instance) RequestInstance(local_svc_name string) (
	common_types.ApplicationModuleInstance,
	common_types.ApplicationModule,
	error,
) {
}

func (self *Instance) ShowWindow() error {
}

func test(a common_types.ApplicationModule) {
	test(&Module{})
}

