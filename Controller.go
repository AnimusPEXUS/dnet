package dnet

import (
	"errors"
	"fmt"
	"net/rpc"
	//	"crypto/rsa"
	//"crypto/tls"
	//"errors"
	"net"

	"golang.org/x/crypto/ssh"

	"github.com/AnimusPEXUS/dnet/common_types"
	"github.com/AnimusPEXUS/worker"
)

type Controller struct {
	*worker.Worker

	module_controller common_types.ModuleController
}

func NewController(module_controller common_types.ModuleController) (
	*Controller,
	error,
) {
	if module_controller == nil {
		return nil, errors.New("module_controller must be specified")
	}
	ret := new(Controller)
	ret.module_controller = module_controller
	return ret, nil
}

func (self *Controller) TestOutgoingConnection(conn net.Conn) error {

	repl := &struct{ Text string }{}

	cli := rpc.NewClient(conn)

	err := cli.Call(
		"Node.IsReallyDNetNode",
		&struct{ Text string }{Text: "Hi!"},
		repl,
	)
	if err != nil {
		fmt.Println("debug error TestOutgoingConnection", err)
		return err
	}

	if repl.Text != "yes it is" {
		return errors.New("RPC challange failed")
	}

	return nil
}

func (self *Controller) ServeConnection(
	who *common_types.Address,
	conn net.Conn,
) {

	config := &ssh.ServerConfig{
		NoClientAuth:                true,
		MaxAuthTries:                2,
		PasswordCallback:            self._ISSHH_PasswordCallback,
		PublicKeyCallback:           self._ISSHH_PublicKeyCallback,
		KeyboardInteractiveCallback: self._ISSHH_KeyboardInteractiveCallback,
		AuthLogCallback:             self._ISSHH_AuthLogCallback,
		ServerVersion:               "",
	}

	ssh_conn, new_chan, request, err := ssh.NewServerConn(conn, config)
	if err != nil {
		fmt.Println("error", err)
	}

	return
}

func (self *Controller) _ISSHH_PasswordCallback(
	conn ssh.ConnMetadata,
	password []byte,
) (
	*ssh.Permissions,
	error,
) {

}

func (self *Controller) _ISSHH_PublicKeyCallback(
	conn ssh.ConnMetadata,
	key ssh.PublicKey,
) (
	*ssh.Permissions,
	error,
) {
	ret := &ssh.Permissions{}
	return ret, nil
}

func (self *Controller) _ISSHH_KeyboardInteractiveCallback(
	conn ssh.ConnMetadata,
	client ssh.KeyboardInteractiveChallenge,
) (
	*ssh.Permissions,
	error,
) {
	ret := &ssh.Permissions{}
	return ret, nil
}

func (self *Controller) _ISSHH_AuthLogCallback(
	conn ssh.ConnMetadata,
	method string,
	err error,
) {
	fmt.Println("authentication attempted", conn, method, err)
}

func (self *Controller) FoundPossibleNode(*common_type.TransportAddress) {
}

func (self *Controller) ProbeAddress(
	sync bool,
	callback func(
		success bool,
		dnet_address *common_type.Address,
		transport_address *common_type.TransportAddress,
		arg interface{},
	),
	arg interface{},
) (
	bool,
	*common_type.Address,
	*common_type.TransportAddress,
) {
}
