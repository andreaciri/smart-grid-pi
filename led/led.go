package led

import (
	"github.com/stianeikeland/go-rpio"
)

type Led struct {
	red    rpio.Pin
	yellow rpio.Pin
	green  rpio.Pin
}

func NewLed() (*Led, error) {
	if err := rpio.Open(); err != nil {
		return nil, err
	}

	red := rpio.Pin(23)    // GPIO 23, physical pin 16
	yellow := rpio.Pin(24) // GPIO 24, physical pin 18
	green := rpio.Pin(25)  // GPIO 25, physical pin 22

	red.Output()
	yellow.Output()
	green.Output()

	return &Led{
		red:    red,
		yellow: yellow,
		green:  green,
	}, nil
}

func (l Led) SwitchRed() {
	l.yellow.Low()
	l.green.Low()
	l.red.High()
}

func (l Led) SwitchYellow() {
	l.red.Low()
	l.green.Low()
	l.yellow.High()
}

func (l Led) SwitchGreen() {
	l.red.Low()
	l.yellow.Low()
	l.green.High()
}
