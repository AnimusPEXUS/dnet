package dnet

import (
	"errors"
	"fmt"
	"net/rpc"
	"time"
	//	"crypto/rsa"
	//"crypto/tls"
	//"errors"
	"net"

	"golang.org/x/crypto/ssh"

	"github.com/AnimusPEXUS/gologger"

	"github.com/AnimusPEXUS/dnet/common_types"
)

// This is for cases when DNet (as a Controller), whises to access module
// instance or theyr some other parts. In this case, DNet controller will be
// presented as DNET_UNIVERSAL_APPLICATION_NAME
const DNET_UNIVERSAL_APPLICATION_NAME = "localDNet"

type Controller struct {
	application_controller common_types.ApplicationControllerI
	logger                 *gologger.Logger
}

func NewController(
	application_controller common_types.ApplicationControllerI,
	logger *gologger.Logger,
) (
	*Controller,
	error,
) {
	if application_controller == nil {
		return nil, errors.New("application_controller must be specified")
	}
	ret := new(Controller)
	ret.application_controller = application_controller

	logger.Info("DNet Controller initiation completed without errors")
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

	//	ssh_conn, new_chan, request, err := ssh.NewServerConn(conn, config)
	_, _, _, err := ssh.NewServerConn(conn, config)
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
	ret := &ssh.Permissions{}
	return ret, nil
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

func (self *Controller) PossiblyNodeDiscoveredNotificationReceptor(
	module_name *common_types.ModuleName,
	address common_types.NetworkAddress,
) error {
	tracker_module_rpc_client, err := self.application_controller.GetInnodeRPC(
		common_types.ModuleNameNewF(DNET_UNIVERSAL_APPLICATION_NAME),
		common_types.ModuleNameNewF("builtin_address_tracker"),
	)
	if err != nil {
		return err
	}

	dnet_address := &common_types.Address{}

	var res bool

	t := time.Now().UTC()

	tracker_module_rpc_client.Call(
		"RPC.NoteRecord",
		&struct {
			DnetAddrStr    string
			DiscoveredDate *time.Time
			NetworkName    string
			NetworkAddress common_types.NetworkAddress
		}{
			DnetAddrStr:    dnet_address.String(),
			DiscoveredDate: &t,
			NetworkName:    module_name.Value(),
			NetworkAddress: address,
		},
		&res,
	)

	return nil
}

func (self *Controller) GetInnodeRPC(calling_app_name *common_types.ModuleName) (
	*rpc.Client,
	error,
) {
	pipe1, pipe2 := net.Pipe()
	serv := rpc.NewServer()
	serv.RegisterName("DNET", NewInnodeRPC(self, calling_app_name))
	serv.ServeConn(pipe1)
	ret := rpc.NewClient(pipe2)
	return ret, nil
}
