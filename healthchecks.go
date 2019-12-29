package healthcheck

import (
	"reflect"
	"time"

	"github.com/jspc/go-health/models"
)

var (
	// Tick is the time to wait between healthcheck tests
	Tick = 10 * time.Second
)

// Healthchecks wraps the Healthchecks model and provides functions and
// things
type Healthchecks models.Healthchecks

// New returns a Healthchecks object which exposes an API and contains
// timers and logic for running healthchecks
func New(v Version) *Healthchecks {
	return &Healthchecks{
		Healthchecks: make([]models.Healthcheck, 0),

		Timer:   Tick,
		Version: models.Version(v),
	}
}

// Add takes a Healthcheck and enrolls it into the Healthchecks thing
func (h *Healthchecks) Add(hc Healthcheck) {
	h.Healthchecks = append(h.Healthchecks, models.Healthcheck(hc))
}

// Start will run a timer doing healthchecks and stuff
func (h *Healthchecks) Start() {
	for _ = range time.Tick(h.Timer) {
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
		}(Healthcheck(hc))
	}

	rechecked := make([]models.Healthcheck, 0)
	for hc := range c {
		rechecked = append(rechecked, models.Healthcheck(hc))

		if len(rechecked) == size {
			close(c)
		}
	}

	h.ReportTime = time.Now()
	h.Healthchecks = rechecked
}

func (h Healthchecks) byType(t string) (hList []models.Healthcheck) {
	hList = make([]models.Healthcheck, 0)

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
