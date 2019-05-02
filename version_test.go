package healthcheck

import (
	"os"
	"reflect"
	"testing"
)

func TestVersio_FromEnv(t *testing.T) {
	os.Clearenv()

	os.Setenv("CT_ORACLE", "https://example.com/oracle")
	os.Setenv("CT_RUNBOOK", "https://example.com/runbook")
	os.Setenv("CT_SQUAD", "blazing_squad")
	os.Setenv("CT_TIER", "99")

	v := Version{
		Name:      "testing",
		Built:     578433600,
		CircleSha: "f00",
	}

	v.FromEnv()

	expect := Version{
		Name:      "testing",
		Built:     578433600,
		CircleSha: "f00",
		Oracle:    "https://example.com/oracle",
		Runbook:   "https://example.com/runbook",
		Squad:     "blazing_squad",
		Tier:      99,
	}

	if !reflect.DeepEqual(expect, v) {
		t.Errorf("expected %+v, received %+v", expect, v)
	}
}
