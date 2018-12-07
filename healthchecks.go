package healthcheck

import (
	"reflect"
	"time"
)

var (
	// Tick is the time to wait between healthcheck tests
	Tick = 10 * time.Second
)

// Healthchecks represent the state of a run set of healthchecks
// and
type Healthchecks struct {
	ReportTime   time.Time     `json:"report_as_of"`
	Healthchecks []Healthcheck `json:"healthchecks"`

	timer   time.Duration
	version Version
}

// New returns a Healthchecks object which exposes an API and contains
// timers and logic for running healthchecks
func New(v Version) Healthchecks {
	return Healthchecks{
		Healthchecks: make([]Healthcheck, 0),

		timer:   Tick,
		version: v,
	}
}

// Add takes a Healthcheck and enrolls it into the Healthchecks thing
func (h *Healthchecks) Add(hc Healthcheck) {
	h.Healthchecks = append(h.Healthchecks, hc)
}

// Start will run a timer doing healthchecks and stuff
func (h *Healthchecks) Start() {
	for _ = range time.Tick(h.timer) {
		go h.recheck()
	}
}

func (h *Healthchecks) recheck() {
	size := len(h.Healthchecks)
	c := make(chan Healthcheck)

	for _, hc := range h.Healthchecks {
		go func(hc1 Healthcheck) {
			// Ensure correct healthcheck is scoped
			hc1 = hc1

			hc1.Run()
			c <- hc1
		}(hc)
	}

	rechecked := make([]Healthcheck, 0)
	for hc := range c {
		rechecked = append(rechecked, hc)

		if len(rechecked) == size {
			close(c)
		}
	}

	h.ReportTime = time.Now()
	h.Healthchecks = rechecked
}

func (h Healthchecks) byType(t string) (hList []Healthcheck) {
	hList = make([]Healthcheck, 0)

	for _, hc := range h.Healthchecks {
		v := reflect.ValueOf(hc).FieldByName(t)

		if !v.IsValid() {
			continue
		}

		if _, ok := v.Interface().(bool); !ok {
			return
		}

		if v.Bool() {
			hList = append(hList, hc)
		}
	}

	return
}
