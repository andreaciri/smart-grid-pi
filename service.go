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

type Service struct {
	solarEdgeClient SolarEdgeClient
	relay           Relay
	refresh         time.Duration
}

func NewService(solarEdgeClient SolarEdgeClient, relay Relay, refresh time.Duration) Service {
	return Service{
		solarEdgeClient: solarEdgeClient,
		relay:           relay,
		refresh:         refresh,
	}
}

func (s Service) Run() error {

	for {
		// power, err := s.solarEdgeClient.GetCurrentPower()
		// if err != nil {
		// 	return err
		// }
		// log.Println("Power! ", power)

		s.relay.Toggle()
		time.Sleep(s.refresh)
	}

}

func energySurplus(power solaredge.Power) bool {
	return power.SiteCurrentPowerFlow.Pv.Status == solaredge.PhotovoltaicStatusActive &&
		power.SiteCurrentPowerFlow.Pv.CurrentPower > 1.5
}
