package common_types

type Address struct {
	value string
}

func AddressNew(value string) (*Address, error) {
	ret := new(Address)
	ret.value = value
	return ret, nil
}

func (self *Address) String() string {
	return self.value
}
