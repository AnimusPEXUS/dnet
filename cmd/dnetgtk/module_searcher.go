package main

import (
	"errors"
	"fmt"
	"plugin"

	"github.com/AnimusPEXUS/dnet/common_types"
)

type ApplicationModuleWrap struct {
}

func ApplicationModuleWrapNew(
	mod common_types.ApplicationModule,
) *ApplicationModuleWrap {
	ret := new(ApplicationModuleWrap)
	return ret
}

type ModuleSercherSearchResult struct {
	parent_searcher *ModuleSercher
	name            *common_types.ModuleName
	builtin         bool
	path            string
	checksum        string
}

/*
	Note: Name() returns valid value, only if .builtin == true.
				If .builtin == false, You have to use .Mod().Name()
*/
func (self *ModuleSercherSearchResult) Name() *common_types.ModuleName {
	return self.name
}

func (self *ModuleSercherSearchResult) Builtin() bool {
	return self.builtin
}

func (self *ModuleSercherSearchResult) Path() string {
	return self.path
}

func (self *ModuleSercherSearchResult) Checksum() string {
	return self.checksum
}

/*
 Warning: Using .Mod() if .builtin == false, presumes checkking .path's
 checksum consistency and opening it as go plugin, so use with caution!
*/
func (self *ModuleSercherSearchResult) Mod() (
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
		symb, err := plug.Lookup("ModuleReturner")
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

type ModuleSercher struct {
	builtin []common_types.ApplicationModule
}

func ModuleSercherNew(
	builtin []common_types.ApplicationModule,
) *ModuleSercher {
	ret := new(ModuleSercher)
	/*
		for _, i := range builtin {
			if !common_types.IsApplicationNameCorrect(i.Name()) {
				panic("found incorrect builtin application module name")
			}
		}
	*/
	ret.builtin = builtin
	return ret
}

func (self *ModuleSercher) ListModules() []*ModuleSercherSearchResult {
	ret := make([]*ModuleSercherSearchResult, 0)

	for _, i := range self.builtin {
		ret = append(
			ret,
			&ModuleSercherSearchResult{
				parent_searcher: self,
				name:            i.Name(),
				builtin:         true,
				path:            "-",
				checksum:        "-",
			},
		)
	}

	return ret
}

// Depending on builtin value,  GetMod() will use eather name or checksum
// NOTE: the type of checksum may change in the future, fo instance, to better
// describe checksum method desired
func (self *ModuleSercher) GetMod(
	builtin bool,
	name *common_types.ModuleName,
	checksum *common_types.ModuleChecksum,
) (
	common_types.ApplicationModule,
	error,
) {

	res, err := self.SearchMod(builtin, name, checksum)
	if err != nil {
		return nil, errors.New("error. module search error: " + err.Error())
	}

	res2, err := res.Mod()
	if err != nil {
		return nil, errors.New("error. module aquire error: " + err.Error())
	}

	return res2, nil
}

func (self *ModuleSercher) SearchMod(
	builtin bool,
	name *common_types.ModuleName,
	checksum *common_types.ModuleChecksum,
) (
	*ModuleSercherSearchResult,
	error,
) {

	res := self.ListModules()

	if builtin {
		for _, i := range res {
			if i.builtin {
				if i.name.Value() == name.Value() {
					return i, nil
				}
			}
		}
	} else {
		if !checksum.Valid() {
			return nil, errors.New("given checksum is invalid")
		} else {
			if checksum.Meth() != "md5" {
				return nil, errors.New("only md5 sums are supported")
			}
		}
		for _, i := range res {
			if !i.builtin {
				if i.checksum == checksum.Sum() {
					return i, nil
				}
			}
		}
	}

	return nil, errors.New("module not found")
}
