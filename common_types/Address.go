package common_types

type Address struct {
}

func AddressNew(value string) (*Address, error) {
	ret := new(Address)

	return ret, nil
}
