package calc

import (
	"testing"
)

func TestNew(t *testing.T) {
	c := New(0.0, 0.0, 0.0, "male")

	if c.String() != "0e+00" {
		t.Errorf("%s", c.String())
	}
}
