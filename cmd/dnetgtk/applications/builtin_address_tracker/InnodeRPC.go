package builtin_address_tracker

type InnodeRPC struct {
	instance         *Instance
	calling_app_name string
}

func NewInnodeRPC(
	instance *Instance,
	calling_app_name string,
) *InnodeRPC {
	ret := new(InnodeRPC)
	ret.instance = instance
	ret.calling_app_name = calling_app_name
	return ret
}
