package pdapi

import (
	"fmt"
	"time"
)

//Entity can be either a user, schedule, or an escalation policy
type Entity struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Summary string `json:"summary"`
	Self    string `json:"self"`
	HTMLURL string `json:"html_url"`
}

//Oncall is from the JSON response from PD that we want to model, e.g.:
//{
//    "escalation_policy": {
//      "id": "xxxxxxx",
//      "type": "escalation_policy_reference",
//      "summary": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
//      "self": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
//      "html_url": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
//    },
//    "escalation_level": 1,
//    "schedule": {
//      "id": "xxxxxxx",
//      "type": "schedule_reference",
//      "summary": "xxxxxxxxxxxxxxxxxxxxxxxxxxx",
//      "self": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
//      "html_url": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
//    },
//    "user": {
//      "id": "xxxxxxx",
//      "type": "user_reference",
//      "summary": "xxxxxxxxx",
//      "self": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
//      "html_url": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
//    },
//    "start": "xxxxxxxxxxxxxxxxxxxx",
//    "end": "xxxxxxxxxxxxxxxxxxxx"
//}
type Oncall struct {
	EscalationPolicy Entity    `json:"escalation_policy"`
	EscalationLevel  uint8     `json:"escalation_level"`
	Schedule         Entity    `json:"schedule"`
	User             Entity    `json:"user"`
	Start            time.Time `json:"start"`
	End              time.Time `json:"end"`
}

//Schedules models the JSON response from call to PD endpoint /oncalls
type Schedules struct {
	Oncalls []Oncall `json:"oncalls"`
	Limit   uint16   `json:"limit"`
	More    bool     `json:"more"`
	Offset  uint16   `json:"offset"`
	Total   uint16   `json:"total"`
}

func (s *Schedules) String() string {
	var buff string
	for _, o := range s.Oncalls {
		if o.EscalationLevel == 1 {
			prefix := ""
			if o.Start.Weekday().String() == "Saturday" || o.Start.Weekday().String() == "Sunday" {
				prefix = "*"
			}
			buff += fmt.Sprintf("%s%v, %v: %v\n", prefix, o.Start.Weekday(), o.Start, o.EscalationPolicy.Summary)
		}
	}
	return buff
}
