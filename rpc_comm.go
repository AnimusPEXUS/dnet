package dnet

type RPCComm struct {
}

func (self *RPCComm) Read(p []byte) (n int, err error) {

}

// Method for Writting some sort of stream which writes RPC commands and
// wants from Node something
// this should be of standard io.Reader interface
func (self *RPCComm) Write(p []byte) (n int, err error) {
}
