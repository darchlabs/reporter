package commander

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/darchlabs/reporter"
	"github.com/darchlabs/synchronizer-v2/pkg/api/event"
)

func synchronizers(url string, s Storage) error {
	// fetch to synchronizers
	r, err := http.Get(fmt.Sprintf("%s/api/v1/events", url))
	if err != nil {
		return err
	}

	// parse body for synchronizers
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	// unmarshall body response
	rr := &event.ListEventResponse{}
	err = json.Unmarshal(body, rr)
	if err != nil {
		return err
	}

	// define base vars
	now := time.Now()
	reports := make([]*reporter.Report, 0)

	// iterate over response and prepare populate reports
	for _, v := range rr.Data {
		report := &reporter.Report{
			ID: v.ID,
			Status: rr.Meta.Cronjob.Status,
			CreatedAt: now,
		}
		
		reports = append(reports, report)
	}

	// prepare group report
	groupReport := &reporter.GroupReport{
		Type: reporter.ServiceTypeSynchronizers,
		Reports: reports,
		CreatedAt: now,
	}

	// insert group report inside database
	err = s.InsertGroupReport(groupReport)
	if err != nil {
		return err
	}
	
	return nil
}