package main

import (
	"log"
	"time"

	"github.com/andreaciri/smart-grid-pi/solaredge"
)

const (
	stateOFF = "OFF"
	stateON  = "ON"
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
	solarEdgeClient         SolarEdgeClient
	relay                   Relay
	led                     Led
	refresh                 time.Duration
	minimumPowerSurplusWatt int
}

func NewService(solarEdgeClient SolarEdgeClient, relay Relay, led Led, refresh time.Duration, minimumPowerSurplusWatt int) Service {
	return Service{
		solarEdgeClient:         solarEdgeClient,
		relay:                   relay,
		led:                     led,
		refresh:                 refresh,
		minimumPowerSurplusWatt: minimumPowerSurplusWatt,
	}
}

func (s Service) Run() error {

	state := stateOFF

	defer func() {
		s.relay.SwitchOff()
		s.led.SwitchOff()
	}()

	for {

		if s.nightTime() {
			log.Println("going to sleep...")
			state = stateOFF
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
		case state == stateOFF && power.ProductionSurplus(s.minimumPowerSurplusWatt):
			// enough surplus, turn smart grid ON
			state = stateON
			s.relay.SwitchOn()
			s.led.SwitchGreen()
			s.log(*power, state)

		case state == stateON && power.ProductionSurplus(0):
			// keep smart grid ON when powered by self-consumption
			s.log(*power, state)

		case power.Production():
			// production not enough, turn smart grid OFF
			state = stateOFF
			s.relay.SwitchOff()
			s.led.SwitchYellow()
			s.log(*power, state)

		default:
			// no production, turn smart grid OFF
			state = stateOFF
			s.relay.SwitchOff()
			s.led.SwitchRed()
			s.log(*power, state)
		}

		time.Sleep(s.refresh)
	}

}

func (s Service) nightTime() bool {
	now := time.Now()
	return now.Hour() <= 8 || now.Hour() >= 18
}

func (s Service) log(power solaredge.Power, state string) {
	log.Printf(
		"Production %f kW, Consumption %f kW, smart grid %s\n",
		power.SiteCurrentPowerFlow.Pv.CurrentPower,
		power.SiteCurrentPowerFlow.Load.CurrentPower,
		state,
	)
}
