package newrelic

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	V2_ROOT = "https://api.newrelic.com/v2/"
)

type GetAllAlertsResponse struct {
	Policies []Policy `json:"policies"`
}

type Policy struct {
	Id                 int    `json:"id"`
	IncidentPreference string `json:"incident_preference"`
	CreatedAt          int64  `json:"created_at"`
	UpdatedAt          int64  `json:"updated_at"`
}

func NewNewRelicService(apiKey string, version string) *NewRelic {
	nr := &NewRelic{
		apiKey:  apiKey,
		version: version,
	}

	return nr
}

func (nr *NewRelic) GetAllAlerts() (GetAllAlertsResponse, error) {
	req, makeReqError := http.NewRequest("GET", "https://api.newrelic.com/v2/alerts_policies.json", nil)
	if makeReqError != nil {
		log.Fatalf("Error: failed to create request \n%s\n", makeReqError.Error())
		os.Exit(0)
	}

	req.Header.Set("X-Api-Key", nr.apiKey)
	client := &http.Client{Timeout: time.Second * 10}
	resp, reqError := client.Do(req)
	if reqError != nil {
		log.Fatalf("ERROR: request failed. \n%s\n  ", reqError.Error())
		os.Exit(0)
	}

	defer resp.Body.Close()
	body, readAllError := ioutil.ReadAll(resp.Body)
	if readAllError != nil {
		log.Fatalf("ERROR: failed to retrieve all alerts. \n%s\n", readAllError.Error())
		os.Exit(0)
	}

	var allAlerts GetAllAlertsResponse
	unmarshallError := json.Unmarshal(body, &allAlerts)
	if unmarshallError != nil {
		log.Fatalf("ERROR: failed to unmarshal response. \n%s\n", unmarshallError.Error())
		os.Exit(0)
	}

	return allAlerts, nil
}
