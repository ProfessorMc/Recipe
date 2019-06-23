package appliance

import (
	"errors"
	"fmt"
)

type Appliance interface {
	TurnOn() error
	TurnOff() error
	IsOn() bool
	SetPower(bool)
	Label
}

type Label interface {
	GetName() string
	GetBrand() string
}

func Plugin(appliance Appliance) {
	appliance.SetPower(true)
}

func Unplug(appliance Appliance) {
	appliance.SetPower(false)
}

func BuildApplianceError(errorMsg string, appliance Appliance) error {
	returnMsg := fmt.Sprintf("[Appliance: %s] Error:%s", appliance.GetName(), errorMsg)
	return errors.New(returnMsg)
}
