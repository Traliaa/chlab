package devices

import (
	chlab "github.com/Traliaa/chlab/api"
	"github.com/Traliaa/chlab/channels"
	"github.com/Traliaa/chlab/logger"
	"github.com/Traliaa/chlab/model"
)

type SwitchDevice struct {
	baseDevice

	ApplyOnOff func(state bool) error

	state        *bool
	onOffChannel *channels.OnOffChannel
}

func (d *SwitchDevice) UpdateSwitchState(state bool) error {
	d.state = &state
	return d.onOffChannel.SendState(*d.state)
}

func (d *SwitchDevice) SetOnOff(state bool) error {

	if state {
		d.log.Infof("Turning On")
	} else {
		d.log.Infof("Turning Off")
	}

	return d.ApplyOnOff(state)
}

func (d *SwitchDevice) ToggleOnOff() error {
	if d.state == nil {
		d.log.Warningf("On-off channel is in an unknown state for toggling. Setting to off.")
		return d.SetOnOff(false)
	}
	return d.SetOnOff(!*d.state)
}

func CreateSwitchDevice(driver chlab.Driver, info *model.Device, conn *chlab.Connection) (*SwitchDevice, error) {

	d := &SwitchDevice{
		baseDevice: baseDevice{
			conn:   conn,
			driver: driver,
			log:    logger.GetLogger("SwitchDevice - " + *info.Name),
			info:   info,
		},
	}

	err := conn.ExportDevice(d)
	if err != nil {
		d.log.Fatalf("Failed to export device %s: %s", *info.Name, err)
	}

	d.onOffChannel = channels.NewOnOffChannel(d)
	d.conn.ExportChannel(d, d.onOffChannel, "on-off")
	if err != nil {
		d.log.Fatalf("Failed to export on-off channel: %s", err)
	}

	d.log.Infof("Created")

	return d, nil
}
