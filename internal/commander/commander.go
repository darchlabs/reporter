package commander

import (
	"fmt"

	"github.com/darchlabs/reporter"
)

type Storage interface {
	InsertGroupReport(g *reporter.GroupReport) error
}

func Reporter(t reporter.ServiceType, url string, s Storage) error {
	if t == reporter.ServiceTypeSynchronizers {
		return synchronizers(url, s)
	}

	if t == reporter.ServiceTypeJobs {
		return jobs(url, s)
	}

	if t == reporter.ServiceTypeNodes {
		return nodes()
	}

	return fmt.Errorf("invalid service_type=%s in Reporter(t)", string(t))
}
