package main

import (
	"fmt"
	"reflect"
)

type ControllerCommunicatorForAppLogger struct {
	p *ControllerCommunicatorForApp
}

func (self *ControllerCommunicatorForAppLogger) Text(txt string) {
	self.p.controller.logger.Text(
		fmt.Sprintf("module '%s': %s", self.p.name.Value(), txt),
	)
}

func (self *ControllerCommunicatorForAppLogger) Info(txt string) {
	self.p.controller.logger.Info(
		fmt.Sprintf("module '%s': %s", self.p.name.Value(), txt),
	)
}

func (self *ControllerCommunicatorForAppLogger) Warning(txt string) {
	self.p.controller.logger.Warning(
		fmt.Sprintf("module '%s': %s", self.p.name.Value(), txt),
	)
}

func (self *ControllerCommunicatorForAppLogger) Error(value interface{}) {
	var value_str string
	switch value.(type) {
	case string:
		value_str = value.(string)
	case error:
		value_str = value.(error).Error()
	default:
		value_str = reflect.ValueOf(value).String()
	}

	self.p.controller.logger.Error(
		fmt.Sprintf("module '%s': %s", self.p.name.Value(), value_str),
	)
}
