package models

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
