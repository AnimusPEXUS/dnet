package builtin_ownkeypair

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

func (self *InnodeRPC) GetOwnPrivKey(arg bool, res *string) error {
	if t, err := self.instance.GetOwnPrivKey(); err != nil {
		return err
	} else {
		*res = t
	}
	return nil
}
