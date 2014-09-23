package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/ninjasphere/go-ninja/api"
	"github.com/ninjasphere/go-ninja/channels"
	"github.com/ninjasphere/go-ninja/model"
	"github.com/ninjasphere/go-ninja/rpc"
)

type FakeLight struct {
	driver       ninja.Driver
	info         *model.Device
	sendEvent    func(event string, payload interface{}) error
	onOffChannel *channels.OnOffChannel
}

func NewFakeLight(driver ninja.Driver, id int) *FakeLight {
	light := &FakeLight{
		driver: driver,
		info: &model.Device{
			ID:     fmt.Sprintf("light%d", id),
			IDType: "fake",
		},
	}

	light.onOffChannel = channels.NewOnOffChannel(light)

	return light
}

func (l *FakeLight) GetDeviceInfo() *model.Device {
	return l.info
}

func (l *FakeLight) GetDriver() ninja.Driver {
	return l.driver
}

func (l *FakeLight) SetOnOff(state bool) error {
	log.Printf("Turning %t", state)
	return nil
}

func (l *FakeLight) ToggleOnOff() error {
	log.Println("Toggling")
	return nil
}

func (l *FakeLight) SetEventHandler(sendEvent func(event string, payload interface{}) error) {
	l.sendEvent = sendEvent
}

var reg, _ = regexp.Compile("[^a-z0-9]")

// Exported by service/device schema
func (l *FakeLight) SetName(message *rpc.Message, name *string, result *string) error {
	log.Printf("Setting device name to %s", *name)

	safe := reg.ReplaceAllString(*name, "")
	result = &safe

	log.Printf("Pretending we can only set 5 lowercase alphanum. Name now: %s", safe)

	return nil
}
