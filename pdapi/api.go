package pdapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//PagerDuty represents the remote API
type PagerDuty struct {
	Key  string
	User string
	URL  string
}

//GetKey returns the pagerduty API token
func (pd *PagerDuty) GetKey() string {
	return pd.Key
}

//GetSchedules returns a Schedules with 3 month's worth of oncall schedules
func (pd *PagerDuty) GetSchedules() (*Schedules, error) {
	//issue a http get to pagerduty APi using your token
	date := getDate(3)
	path := fmt.Sprintf("/oncalls?limit=100&user_ids%%5B%%5D=%s&until=%s", pd.User, date)
	token := fmt.Sprintf("Token token=%s", pd.Key)

	req, err := http.NewRequest("GET", pd.URL+path, nil)
	req.Header.Add("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	schedules := &Schedules{}
	d := json.NewDecoder(resp.Body)
	err = d.Decode(schedules)
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

func getDate(monthsAhead int) string {
	now := time.Now()
	return now.AddDate(0, monthsAhead, 0).Format("2006-01-02")
}
