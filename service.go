package main

import (
	"time"

	"github.com/andreaciri/smart-grid-pi/solaredge"
)

type SolarEdgeClient interface {
	GetCurrentPower() (power *solaredge.Power, err error)
}

type Relay interface {
	SwitchOn()
	SwitchOff()
	Toggle()
}

type Led interface {
	SwitchRed()
	SwitchYellow()
	SwitchGreen()
}

type Service struct {
	solarEdgeClient SolarEdgeClient
	relay           Relay
	led             Led
	refresh         time.Duration
}

func NewService(solarEdgeClient SolarEdgeClient, relay Relay, led Led, refresh time.Duration) Service {
	return Service{
		solarEdgeClient: solarEdgeClient,
		relay:           relay,
		led:             led,
		refresh:         refresh,
	}
}

func (s Service) Run() error {

	lights := []func(){s.led.SwitchRed, s.led.SwitchYellow, s.led.SwitchGreen}
	i := 0

	for {
		// power, err := s.solarEdgeClient.GetCurrentPower()
		// if err != nil {
		// 	return err
		// }
		// log.Println("Power! ", power)

		s.relay.Toggle()

		lights[i]()
		i = i + 1
		if i == 3 {
			i = 0
		}

		time.Sleep(s.refresh)
	}

}

func energySurplus(power solaredge.Power) bool {
	return power.SiteCurrentPowerFlow.Pv.Status == solaredge.PhotovoltaicStatusActive &&
		power.SiteCurrentPowerFlow.Pv.CurrentPower > 1.5
}
