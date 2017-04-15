package dnet

import (
	"crypto/rsa"
	"crypto/tls"
	"errors"
	"net"

	"bitbucket.org/AnimusPEXUS/dnet/common_types"
)

type ControllerApplicationModuleWrapper struct {
	Name string
	Mod  common_types.ApplicationModule
}

func ControllerApplicationModuleWrapperNew(
	mod common_types.ApplicationModule,
) *ControllerApplicationModuleWrapper {
	ret := new(ControllerApplicationModuleWrapper)
	ret.Name = mod.Name()
	ret.Mod = mod
	return ret
}

type ControllerKeyPairRequestCB func() *rsa.PrivateKey
type ControllerTLSCertificateCB func() *tls.Certificate
type ControllerConnectionPassToServiceAsync func() error

type Controller struct {
	ApplicationModules []*ControllerApplicationModuleWrapper
}

func NewController() (*Controller, error) {
	ret := new(Controller)

	return ret, nil
}

func (self *Controller) AppendApplicationModule(
	mod common_types.ApplicationModule,
) error {
	mod_name := mod.Name()

	for _, i := range self.ApplicationModules {
		if i.Name == mod_name {
			return errors.New("module with this name already added")
		}
	}

	self.ApplicationModules = append(
		self.ApplicationModules,
		ControllerApplicationModuleWrapperNew(mod),
	)
	return nil
}

func (self *Controller) Start() error {
	return nil
}

func (self *Controller) Stop() {
}

func (self *Controller) AcceptConn(in_conn *net.Conn) {

}

func (self *Controller) GetApplication(name string) (
	common_types.ApplicationModule,
	error,
) {
	for _, i := range self.ApplicationModules {
		if i.Name == name {
			return i.Mod, nil
		}
	}

	return nil, errors.New("application with such name not found")
}
