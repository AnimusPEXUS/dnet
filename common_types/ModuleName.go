package common_types

import (
	"errors"
	"regexp"
)

var VALID_MODULE_NAME_RE *regexp.Regexp = regexp.MustCompile(
	"^[A-Za-z][0-9A-Za-z_]*$",
)

func IsValidModuleName(value string) bool {
	return VALID_MODULE_NAME_RE.Match([]byte(value))
}

type ModuleName struct {
	value string
}

func ModuleNameNew(value string) (*ModuleName, error) {
	if !IsValidModuleName(value) {
		return nil, errors.New("invalid value")
	}
	ret := new(ModuleName)
	ret.value = value
	return ret, nil
}

func (self *ModuleName) Value() string {
	return self.value
}
