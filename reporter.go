package reporter

import "time"

type ServiceType string

const (
	ServiceTypeSynchronizers ServiceType = "synchronizers"
	ServiceTypeJobs          ServiceType = "jobs"
	ServiceTypeNodes         ServiceType = "nodes"
)

type Report struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type GroupReport struct {
	Type      ServiceType `json:"type"`
	Reports   []*Report   `json:"services"`
	CreatedAt time.Time   `json:"createdAt"`
}
