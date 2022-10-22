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

	red := rpio.Pin(13)    // GPIO 13, physical pin 33
	yellow := rpio.Pin(19) // GPIO 19, physical pin 35
	green := rpio.Pin(26)  // GPIO 26, physical pin 37

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
