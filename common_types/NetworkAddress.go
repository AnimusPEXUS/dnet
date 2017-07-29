package common_types

type NetworkAddress string

func NetworkAddressNewFromString(value string) NetworkAddress {
	return NetworkAddress(value)
}

func (self NetworkAddress) String() string {
	return string(self)
}
