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
