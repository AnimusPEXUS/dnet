package main

import (
	"errors"
	"fmt"
	"plugin"

	"github.com/AnimusPEXUS/dnet/common_types"
)

type ModuleSearcherSearchResult struct {
	parent_searcher *ModuleSearcher
	name            *common_types.ModuleName
	builtin         bool
	path            string
	checksum        string
}

/*
	Note: Name() returns valid value, only if .builtin == true.
				If .builtin == false, You have to use .Mod().Name()
*/
func (self *ModuleSearcherSearchResult) Name() *common_types.ModuleName {
	return self.name
}

func (self *ModuleSearcherSearchResult) Builtin() bool {
	return self.builtin
}

func (self *ModuleSearcherSearchResult) Path() string {
	return self.path
}

func (self *ModuleSearcherSearchResult) Checksum() string {
	return self.checksum
}

/*
 Warning: Using .Mod() if .builtin == false, presumes checking .path's
 checksum consistency and opening it as go plugin, so use with caution!
*/
func (self *ModuleSearcherSearchResult) Mod() (
	common_types.ApplicationModule,
	error,
) {
	if self.builtin {
		for _, i := range self.parent_searcher.builtin {
			if i.Name().Value() == self.name.Value() {
				return i, nil
			}
		}
	} else {
		// TODO: checksum check
		plug, err := plugin.Open(self.path)
		if err != nil {
			return nil, errors.New(
				fmt.Sprintf(
					"couldn't open file (%s) as golang plugin: %s",
					self.path,
					err.Error(),
				),
			)
		}
		symb, err := plug.Lookup("GttDNetGtkModule")
		if err != nil {
			return nil, errors.New(
				fmt.Sprintf(
					"plugin file (%s) ModuleReturner symbol lookup error: %s",
					self.path,
					err.Error(),
				),
			)
		}

		your_little_func, ok := symb.(func() common_types.ApplicationModule)
		if !ok {
			return nil, errors.New(
				fmt.Sprintf(
					"could not use returned symbol as "+
						"(func() common_types.ApplicationModule)",
					self.path,
					err.Error(),
				),
			)
		}

		mod := your_little_func()

		return mod, nil

	}
	return nil, errors.New("some programming error. report if You got it")
}
