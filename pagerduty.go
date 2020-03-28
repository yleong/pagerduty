package main

import (
	"os"

	"github.com/yleong/pagerduty/pdapi"
	"github.com/yleong/pagerduty/server"
)

func main() {
	s := server.Server{
		Port: os.Getenv("PAGERDUTY_PORT"),
		PD: pdapi.PagerDuty{
			Key:  os.Getenv("PAGERDUTY_KEY"),
			User: os.Getenv("PAGERDUTY_USER"),
			URL:  "https://api.pagerduty.com",
		},
	}
	s.Listen()
}
