package main

import (
	"log"
	"os"
	"regexp"

	"github.com/yleong/pagerduty/pdapi"
	"github.com/yleong/pagerduty/server"
)

func main() {
	user := os.Getenv("PAGERDUTY_USER")
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailPattern)

	if re.MatchString(user) {
		log.Printf("PAGERDUTY_USER (%s) looks like an email. It should be the user ID instead (e.g. X8Y4Z0W which matches [A-Z0-9]{7} )\n", user)
		return
	} 

	s := server.Server{
		Port: os.Getenv("PAGERDUTY_PORT"),
		PD: pdapi.PagerDuty{
			Key:  os.Getenv("PAGERDUTY_KEY"),
			User: user,
			URL:  "https://api.pagerduty.com",
		},
	}
	s.Listen()
}
