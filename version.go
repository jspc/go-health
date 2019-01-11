package healthcheck

import (
	"os"

	"github.com/jspc/go-health/models"
)

const (
	OracleEnvVar  = "CT_ORACLE"
	RunbookEnvVar = "CT_RUNBOOK"
	SquadEnvVar   = "CT_SQUAD"
)

// Version wraps the version model with convenience functions
type Version models.Version

// FromEnv will update a Version with the data from the environment.
// It is aimed at kubernetes deploys- there are some magic vars inserted
// into the environment for deploys based on base-service 1.7.0 and above
func (v *Version) FromEnv() {
	v.Oracle = os.Getenv(OracleEnvVar)
	v.Runbook = os.Getenv(RunbookEnvVar)
	v.Squad = os.Getenv(SquadEnvVar)
}
