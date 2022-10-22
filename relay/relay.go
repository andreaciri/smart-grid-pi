package relay

import (
	"github.com/stianeikeland/go-rpio"
)

type Relay struct {
	pin rpio.Pin
}

func NewRelay() (*Relay, error) {
	if err := rpio.Open(); err != nil {
		return nil, err
	}

	pin := rpio.Pin(3) // GPIO 3, physical pin 05
	pin.Output()
	pin.Low()

	return &Relay{
		pin: pin,
	}, nil
}

func (r Relay) SwitchOn() {
	r.pin.High()
}

func (r Relay) SwitchOff() {
	r.pin.Low()
}

func (r Relay) Toggle() {
	r.pin.Toggle()
}
