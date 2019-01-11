package healthcheck

import (
	"testing"

	"github.com/jspc/go-health/models"
)

type dummyCheck struct {
	s bool
	r interface{}
}

func (d dummyCheck) work() (bool, interface{}) {
	return d.s, d.r
}

func TestHealthchecks_Recheck(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("unexpected error: %+v", err)
		}
	}()

	v := Version{}
	h := New(v)

	h.Add(Healthcheck{
		Name: "test-one",
		F:    dummyCheck{true, 123}.work,
	})

	h.Add(Healthcheck{
		Name: "test-two",
		F:    dummyCheck{false, "abc"}.work,
	})

	h.recheck()

	t1 := false
	t2 := false

	for _, hc := range h.Healthchecks {
		if hc.Name == "test-one" {
			t1 = true
			if !hc.Success {
				t.Errorf("expected true, received %+v", hc.Success)
			}
		}

		if hc.Name == "test-two" {
			t2 = true
			if hc.Success {
				t.Errorf("expected false, received %+v", hc.Success)
			}
		}
	}

	if !t1 {
		t.Errorf("never received result for t1")
	}

	if !t2 {
		t.Errorf("never received result for t2")
	}
}

func TestHealthchecks_ByType(t *testing.T) {
	for _, test := range []struct {
		name         string
		healthchecks []models.Healthcheck
		t            string
		expectLen    int
	}{
		{"Two of two, liveness", []models.Healthcheck{models.Healthcheck{Name: "one", Liveness: true}, models.Healthcheck{Name: "two", Liveness: true}}, "Liveness", 2},
		{"One of two, liveness", []models.Healthcheck{models.Healthcheck{Name: "one", Liveness: true}, models.Healthcheck{Name: "two", Liveness: false}}, "Liveness", 1},
		{"None of two, liveness", []models.Healthcheck{models.Healthcheck{Name: "one", Liveness: false}, models.Healthcheck{Name: "two", Liveness: false}}, "Liveness", 0},

		{"Two of two, readiness", []models.Healthcheck{models.Healthcheck{Name: "one", Readiness: true}, models.Healthcheck{Name: "two", Readiness: true}}, "Readiness", 2},
		{"One of two, readiness", []models.Healthcheck{models.Healthcheck{Name: "one", Readiness: true}, models.Healthcheck{Name: "two", Readiness: false}}, "Readiness", 1},
		{"None of two, readiness", []models.Healthcheck{models.Healthcheck{Name: "one", Readiness: false}, models.Healthcheck{Name: "two", Readiness: false}}, "Readiness", 0},

		{"Dodgy name", []models.Healthcheck{models.Healthcheck{Name: "one", Readiness: true}, models.Healthcheck{Name: "two", Readiness: true}}, "Nonsuch", 0},
		{"Not a bool", []models.Healthcheck{models.Healthcheck{Name: "one", Readiness: true}, models.Healthcheck{Name: "two", Readiness: true}}, "RunbookItem", 0},
	} {
		t.Run(test.name, func(t *testing.T) {
			h := Healthchecks{
				Healthchecks: test.healthchecks,
			}

			received := len(h.byType(test.t))
			if test.expectLen != received {
				t.Errorf("expected %d, received %d", test.expectLen, received)
			}
		})
	}
}
