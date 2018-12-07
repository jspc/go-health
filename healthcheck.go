package healthcheck

import (
	"time"
)

// Healthcheck represents an individual healthcheck
type Healthcheck struct {
	// Name is a simple string to help explain the point of the healthcheck
	// It doesn't have to even be unique- it's just to make output more
	// readable
	Name string `json:"name"`

	// Readiness and Liveness determines what check endpoints are affected by
	// what healthchecks- it's also possible to set these as both false- this
	// will, essentially, make a healthcheck advisory only
	Readiness bool `json:"readiness"`
	Liveness  bool `json:"liveness"`

	// RunbookItem should point to the url and anchor, where possible, of
	// the healthcheck/ dependency. It may be in the runbook of the service
	// or it may be another runbook
	RunbookItem string `json:"runbook_item"`

	// F is the function that a healthcheck runs. It returns true for a
	// successful healthcheck, and false for a failing healthcheck.
	// It also returns optional output.
	F func() (bool, interface{}) `json:"-"`

	// The following are overwritten when a healthcheck is run
	State    string      `json:"state"` // enum: not_run, running, run
	LastRun  time.Time   `json:"last_run"`
	Duration float64     `json:"duration_ms"`
	Success  bool        `json:"success"`
	Output   interface{} `json:"output"`
}

// Run will.... run the healthcheck
func (h *Healthcheck) Run() {
	h.State = "running"

	h.LastRun = time.Now()
	h.Success, h.Output = h.F()
	h.Duration = float64(time.Since(h.LastRun)) / 1000000.0

	h.State = "run"
}
