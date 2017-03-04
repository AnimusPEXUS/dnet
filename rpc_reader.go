package dnet

import (
	"io"
)

type (
	DNetRPCServer struct {
		Controller DNetController
	}
)

func (self *DNetRPCServer) StartReading(input *io.Reader) error {

	var code RPCCommandCode

}
