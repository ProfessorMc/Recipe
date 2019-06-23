package appliance

type Appliance interface {
	TurnOn() error
	TurnOff() error
	IsOn() bool
	SetPower(bool)

}

type Label interface {
	GetName() string
	GetBrand() string
}

func Plugin(appliance Appliance) {
	panic("implement Plugin")
}

func Unplug(appliance Appliance) {
	panic("implement Unplug")
}

func BuildApplianceError(errorMsg string, appliance Appliance) error {
	panic("implement BuildApplianceError")

}