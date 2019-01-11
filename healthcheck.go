package healthcheck

import (
	"github.com/jspc/go-health/models"
)

// Healthcheck wraps a healthcheck model to be a top level export
// of this package, for backwards compatability
type Healthcheck models.Healthcheck
