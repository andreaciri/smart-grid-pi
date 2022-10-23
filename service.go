package main

import (
	"log"
	"time"

	"github.com/andreaciri/smart-grid-pi/solaredge"
)

const (
	minimumPowerSurplusWatt = 2000
)

type SolarEdgeClient interface {
	GetCurrentPower() (power *solaredge.Power, err error)
}

type Relay interface {
	SwitchOn()
	SwitchOff()
}

type Led interface {
	SwitchRed()
	SwitchYellow()
	SwitchGreen()
	SwitchOff()
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

	defer func() {
		s.relay.SwitchOff()
		s.led.SwitchOff()
	}()

	for {

		// sleep during night
		if s.nightTime() {
			s.relay.SwitchOff()
			s.led.SwitchOff()
			time.Sleep(s.refresh)
			continue
		}

		power, err := s.solarEdgeClient.GetCurrentPower()
		if err != nil {
			return err
		}

		switch {
		case power.ProductionSurplus(minimumPowerSurplusWatt):
			s.relay.SwitchOn()
			s.led.SwitchGreen()
			s.log(*power, "ON")
			break

		case power.Production():
			s.relay.SwitchOff()
			s.led.SwitchYellow()
			s.log(*power, "OFF")
			break

		default:
			s.relay.SwitchOff()
			s.led.SwitchRed()
			s.log(*power, "OFF")
		}

		time.Sleep(s.refresh)
	}

}

func (s Service) nightTime() bool {
	now := time.Now()
	return now.Hour() <= 8 && now.Hour() >= 18
}

func (s Service) log(power solaredge.Power, enabled string) {
	log.Printf(
		"Production %f kW, Consumption %f kW, smart grid %s\n",
		power.SiteCurrentPowerFlow.Pv.CurrentPower,
		power.SiteCurrentPowerFlow.Load.CurrentPower,
		enabled,
	)
}
