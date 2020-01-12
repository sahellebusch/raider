package alert

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	table "github.com/olekukonko/tablewriter"
	"github.com/spf13/viper"
)

func GetAll() {
	type Policy struct {
		Id                 int    `json:"id"`
		IncidentPreference string `json:"incident_preference"`
		CreatedAt          int64  `json:"created_at"`
		UpdatedAt          int64  `json:"updated_at"`
	}

	type GetAlertsResponse struct {
		Policies []Policy `json:"policies"`
	}

	req, makeReqError := http.NewRequest("GET", "https://api.newrelic.com/v2/alerts_policies.json", nil)
	if makeReqError != nil {
		log.Fatalf("Error: failed to create request \n%s\n", makeReqError.Error())
		os.Exit(0)
	}

	req.Header.Set("X-Api-Key", viper.Get("API_KEY").(string))
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

	var allAlerts GetAlertsResponse
	unmarshallError := json.Unmarshal(body, &allAlerts)
	if unmarshallError != nil {
		log.Fatalf("ERROR: failed to unmarshal response. \n%s\n", unmarshallError.Error())
		os.Exit(0)
	}

	table := table.NewWriter(os.Stdout)
	for _, policy := range allAlerts.Policies {
		createdAt := time.Unix((policy.CreatedAt / int64(time.Microsecond)), 0)
		updatedAt := time.Unix((policy.UpdatedAt / int64(time.Microsecond)), 0)

		table.Append([]string{strconv.Itoa(policy.Id), policy.IncidentPreference, createdAt.String(), updatedAt.String()})
	}
	table.SetHeader([]string{"Id", "Incident Preference", "Created", "Updated"})
	table.Render()

	fmt.Println("\nFor more information about alert policy types, visit the link below.\nhttps://docs.newrelic.com/docs/alerts/new-relic-alerts/configuring-alert-policies/specify-when-new-relic-creates-incidents")
}
