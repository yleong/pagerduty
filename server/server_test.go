package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	pdapi "github.com/yleong/pagerduty/pdapi"
)

func TestRoot(t *testing.T) {
	s := Server{
		Port: os.Getenv("PAGERDUTY_PORT"),
		PD:   pdapi.PagerDuty{Key: os.Getenv("PAGERDUTY_KEY")},
	}
	go s.Listen()
	url := fmt.Sprintf("http://localhost:%s/config", s.Port)

	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("http get %v got err of %v\n", url, err)
	}

	//Want 200
	if resp.StatusCode != http.StatusOK {
		t.Errorf("http get %v got retcode of %d, want %d\n", url, resp.StatusCode, http.StatusOK)
	}

	//Want JSON body
	defer resp.Body.Close()
	d := json.NewDecoder(resp.Body)
	d.DisallowUnknownFields()
	var s2 Server
	err = d.Decode(&s2)
	if err != nil {
		t.Fatalf("http get %v got invalid json response body err %v\n", url, err)
	}

	//Want server config echoed back
	if s != s2 {
		t.Errorf("http get %v got server %+v, want server %+v\n", url, s2, s)
	}

}
