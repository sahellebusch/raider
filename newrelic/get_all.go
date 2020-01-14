package newrelic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Policy is a NewRelic alert policy definition
type Policy struct {
	ID                 int    `json:"id"`
	IncidentPreference string `json:"incident_preference"`
	CreatedAt          int64  `json:"created_at"`
	UpdatedAt          int64  `json:"updated_at"`
}

// GetAllAlerts will fetch the entire list of alerts policies.
func (nr *NewRelic) GetAllAlerts() ([]Policy, error) {

	type getAllAlertsResponse struct {
		Policies []Policy `json:"policies"`
	}

	var error error
	req, error := http.NewRequest("GET", fmt.Sprintf("%s/alerts_policies.json", v2RootURL), nil)
	if error != nil {
		return make([]Policy, 0), error
	}

	req.Header.Set("X-Api-Key", nr.apiKey)
	client := &http.Client{Timeout: time.Second * 10}
	resp, error := client.Do(req)
	if error != nil {
		return make([]Policy, 0), error
	}

	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return make([]Policy, 0), error
	}

	var allAlerts getAllAlertsResponse
	error = json.Unmarshal(body, &allAlerts)
	if error != nil {
		return make([]Policy, 0), error
	}

	return allAlerts.Policies, nil
}
