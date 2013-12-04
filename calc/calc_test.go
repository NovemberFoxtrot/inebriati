package inebriati

import (
	"testing"
)

func TestNew(t *testing.T) {
	i := New(0.0, 0.0, 0.0, "male")

	if i.String() != "0e+00" {
		t.Errorf("%s", i.String())
	}
}

func TestCalc(t *testing.T) {
	test := []struct {
		in  stats
		out string
	}{
		{stats{0.0, 0.0, 0.0, "male", 0.0}, "NaN"},
		{stats{0.0, 0.0, 0.0, "female", 0.0}, "NaN"},
		{stats{3.0, 80.0, 2.0, "male", 0.0}, "3.2534484e-02"},
		{stats{2.5, 70.0, 2.0, "female", 0.0}, "3.6495626e-02"},
	}

	for _, person := range test {
		i := New(person.in.StandardDrinks, person.in.BodyWeightKiloGrams, person.in.DrinkingPeriodHours, person.in.Gender)

		i.Calc()

		if i.String() != person.out {
			t.Errorf("in: %s out: %s for %v", person.out, i.String(), person.in)
		}
	}
}
