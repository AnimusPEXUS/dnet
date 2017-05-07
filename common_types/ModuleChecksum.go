package common_types

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func isallhex(value string) bool {
	ok, err := regexp.MatchString(`^[0-9a-fA-F]+$`, value)
	if err != nil {
		panic(err.Error())
	}
	return ok
}

type ModuleChecksum struct {
	meth string
	sum  string
}

func ModuleChecksumNewFromString(checksum string) (*ModuleChecksum, error) {

	ret := new(ModuleChecksum)

	checksum = strings.ToLower(checksum)
	ind := strings.Index(checksum, ":")

	if ind == -1 {
		if isallhex(checksum) {
			ret.meth = "md5"
			ret.sum = checksum
			return ret, nil
		}
	} else {
		meth := checksum[0:ind]
		sum := checksum[ind:len(checksum)]

		if isallhex(checksum) {
			ret.meth = meth
			ret.sum = sum
			return ret, nil
		}
	}
	return nil, errors.New("invalid input string format")
}

func (self *ModuleChecksum) Valid() bool {
	if self.meth != "md5" {
		return false
	}
	if len(self.sum) != 32 {
		return false
	}
	if !isallhex(self.sum) {
		return false
	}
	return true
}

func (self *ModuleChecksum) String() string {
	/*
		if !self.Valid() {
			panic("invalid. can't be used")
		}
	*/
	return fmt.Sprintf("%s:%s", self.meth, self.sum)
}

func (self *ModuleChecksum) Meth() string {
	/*
		if !self.Valid() {
			panic("invalid. can't be used")
		}
	*/
	return self.meth
}

func (self *ModuleChecksum) Sum() string {
	/*
		if !self.Valid() {
			panic("invalid. can't be used")
		}
	*/
	return self.sum
}
