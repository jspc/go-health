package healthcheck

import (
	"os"
)

const (
	OracleEnvVar  = "CT_ORACLE"
	RunbookEnvVar = "CT_RUNBOOK"
	SquadEnvVar   = "CT_SQUAD"
)

// Version exposes the version, version config, and basic
// links to bits and bobs it needs
type Version struct {
	Name      string `json:"release_name"`
	Built     int64  `json:"built"`
	CircleSha string `json:"version"`
	Oracle    string `json:"oracle"`
	Runbook   string `json:"runbook"`
	Squad     string `json:"squad"`
}

// FromEnv will update a Version with the data from the environment.
// It is aimed at kubernetes deploys- there are some magic vars inserted
// into the environment for deploys based on base-service 1.7.0 and above
func (v *Version) FromEnv() {
	v.Oracle = os.Getenv(OracleEnvVar)
	v.Runbook = os.Getenv(RunbookEnvVar)
	v.Squad = os.Getenv(SquadEnvVar)
}
