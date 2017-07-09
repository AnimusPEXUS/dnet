package main

import "github.com/AnimusPEXUS/dnet/common_types"

type ApplicationModuleWrap struct {
}

func NewApplicationModuleWrap(
	mod common_types.ApplicationModule,
) *ApplicationModuleWrap {
	ret := new(ApplicationModuleWrap)
	return ret
}
