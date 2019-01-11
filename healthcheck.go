package healthcheck

import (
	"time"

	"github.com/jspc/go-health/models"
)

// Healthcheck wraps a healthcheck model to be a top level export
// of this package, for backwards compatability
type Healthcheck models.Healthcheck

// Run will.... run the healthcheck
func (h *Healthcheck) Run() {
	h.State = "running"

	h.LastRun = time.Now()
	h.Success, h.Output = h.F()
	h.Duration = float64(time.Since(h.LastRun)) / 1000000.0

	h.State = "run"
}
