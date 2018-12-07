package healthcheck

import (
	"testing"
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