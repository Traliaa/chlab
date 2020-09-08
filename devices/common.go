package devices

import (
	chlab "github.com/Traliaa/chlab/api"
	"github.com/Traliaa/chlab/logger"
	"github.com/Traliaa/chlab/model"
)

type baseDevice struct {
	log       *logger.Logger
	driver    chlab.Driver
	info      *model.Device
	conn      *chlab.Connection
	sendEvent func(event string, payload interface{}) error
}

func (d *baseDevice) GetDeviceInfo() *model.Device {
	return d.info
}

func (d *baseDevice) GetDriver() chlab.Driver {
	return d.driver
}

func (d *baseDevice) SetEventHandler(sendEvent func(event string, payload interface{}) error) {
	d.sendEvent = sendEvent
}

func (d *baseDevice) Log() *logger.Logger {
	return d.log
}

// Re-expressed with public members so that drivers that don't want to use the Device structs devices.*
// can do so without re-implementing all the Device methods themselves.

type BaseDevice struct {
	Driver chlab.Driver
	Info   *model.Device

	Conn *chlab.Connection
	Log_ *logger.Logger

	SendEvent func(event string, payload interface{}) error
}

func (d *BaseDevice) GetDeviceInfo() *model.Device {
	return d.Info
}

func (d *BaseDevice) GetDriver() chlab.Driver {
	return d.Driver
}

func (d *BaseDevice) SetEventHandler(sendEvent func(event string, payload interface{}) error) {
	d.SendEvent = sendEvent
}

func (d *BaseDevice) Log() *logger.Logger {
	return d.Log_
}
