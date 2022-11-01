package solaredge

const (
	PhotovoltaicStatusActive = "Active"
	PhotovoltaicStatusIdle   = "Idle"
)

type Power struct {
	SiteCurrentPowerFlow SiteCurrentPowerFlow `json:"siteCurrentPowerFlow"`
}
type Connections struct {
	From string `json:"from"`
	To   string `json:"to"`
}
type Grid struct {
	Status       string  `json:"status"`
	CurrentPower float64 `json:"currentPower"`
}
type Load struct {
	Status       string  `json:"status"`
	CurrentPower float64 `json:"currentPower"`
}
type Pv struct {
	Status       string  `json:"status"`
	CurrentPower float64 `json:"currentPower"`
}
type SiteCurrentPowerFlow struct {
	UpdateRefreshRate int           `json:"updateRefreshRate"`
	Unit              string        `json:"unit"`
	Connections       []Connections `json:"connections"`
	Grid              Grid          `json:"GRID"`
	Load              Load          `json:"LOAD"`
	Pv                Pv            `json:"PV"`
}

// Production returns true if the system is producing electricity
func (p Power) Production() bool {
	return p.SiteCurrentPowerFlow.Pv.Status == PhotovoltaicStatusActive &&
		p.SiteCurrentPowerFlow.Pv.CurrentPower > 0
}

// ProductionSurplus returns true if the system is producing electricity,
// and the current unused power is greather or equal to thresholdWatt.
func (p Power) ProductionSurplus(thresholdWatt int) bool {
	return p.Production() &&
		(p.SiteCurrentPowerFlow.Pv.CurrentPower-p.SiteCurrentPowerFlow.Load.CurrentPower) >=
			float64(thresholdWatt)/1000
}
