package dnet

import (
	"bitbucket.org/AnimusPEXUS/dnet/common_types"
	"bitbucket.org/AnimusPEXUS/dnet/network_modules/tcpip"
)

var BUILTIN_NETWORK_MODULES []common_types.NetworkModule

func init() {
	BUILTIN_NETWORK_MODULES = append(
		BUILTIN_NETWORK_MODULES,
		tcpip.NewModule(),
	)
}
