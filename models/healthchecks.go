package models

import (
	"time"
)

// Healthchecks represent the state of a run set of healthchecks
// and stuff
type Healthchecks struct {
	ReportTime   time.Time     `json:"report_as_of"`
	Healthchecks []Healthcheck `json:"healthchecks"`

	Version Version       `json:"-"`
	Timer   time.Duration `json:"-"`
}
