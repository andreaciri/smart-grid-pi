package solaredge

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	getPowerPath = "/site/%s/currentPowerFlow?api_key=%s"
)

type Client struct {
	baseURL string
	siteID  string
	apiKey  string
}

func NewClient(baseURL, siteID, apiKey string) Client {
	return Client{baseURL: baseURL, siteID: siteID, apiKey: apiKey}
}

func (c Client) GetCurrentPower() (power *Power, err error) {

	resp, err := http.Get(fmt.Sprintf(c.baseURL+getPowerPath, c.siteID, c.apiKey))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &power); err != nil {
		return nil, err
	}
	return power, nil
}
