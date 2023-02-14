package commander

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/darchlabs/nodes/pkg/client"
	"github.com/darchlabs/reporter"
)

func nodes(url string, s Storage) error {
	// fetch to jobs
	r, err := http.Get(fmt.Sprintf("%s/api/v1/nodes/status", url))
	if err != nil {
		return err
	}

	// parse body for jobs
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	// unmarshall body response
	rr := &client.GetStatusHandlerResponse{}
	err = json.Unmarshal(body, rr)
	if err != nil {
		return err
	}

	// define base vars
	now := time.Now()
	reports := make([]*reporter.Report, 0)

	// iterate over response and prepare populate reports
	for _, v := range rr.Nodes {
		report := &reporter.Report{
			ID:        v.ID,
			Status:    string(v.Status),
			CreatedAt: now,
		}

		reports = append(reports, report)
	}

	// prepare group report
	groupReport := &reporter.GroupReport{
		Type:      reporter.ServiceTypeNodes,
		Reports:   reports,
		CreatedAt: now,
	}

	// insert group report inside database
	err = s.InsertGroupReport(groupReport)
	if err != nil {
		return err
	}

	return nil
}
