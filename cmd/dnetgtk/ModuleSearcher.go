package main

import (
	"errors"

	"github.com/AnimusPEXUS/dnet/common_types"
)

type ModuleSearcher struct {
	builtin common_types.ApplicationModuleMap
}

func ModuleSearcherNew(
	builtin common_types.ApplicationModuleMap,
) *ModuleSearcher {
	ret := new(ModuleSearcher)
	ret.builtin = builtin
	return ret
}

//func (self *ModuleSearcher) GetBuiltinModules() common_types.ApplicationModuleMap {
//	return self.builtin
//}

func (self *ModuleSearcher) ListModules() []*ModuleSearcherSearchResult {
	ret := make([]*ModuleSearcherSearchResult, 0)

	for _, i := range self.builtin {
		ret = append(
			ret,
			&ModuleSearcherSearchResult{
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
func (self *ModuleSearcher) GetMod(
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

// NOTE:  name does not mean if builtin == false
// NOTE:  checksum does not mean if builtin == true
func (self *ModuleSearcher) SearchMod(
	builtin bool,
	name *common_types.ModuleName,
	checksum *common_types.ModuleChecksum,
) (
	*ModuleSearcherSearchResult,
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
