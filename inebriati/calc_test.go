package inebriati

import (
	"testing"
)

func TestCalc(t *testing.T) {
	test := []struct {
		in  stats
		out string
	}{
		{stats{0.0, 0.0, 0.0, "male", 0.0}, "NaN g/dL"},
		{stats{0.0, 0.0, 0.0, "female", 0.0}, "NaN g/dL"},
		{stats{3.0, 80.0, 2.0, "male", 0.0}, "0.032534484 g/dL"},
		{stats{2.5, 70.0, 2.0, "female", 0.0}, "0.036495626 g/dL"},
	}

	for _, person := range test {
		i := New(person.in.StandardDrinks, person.in.BodyWeightKiloGrams, person.in.DrinkingPeriodHours, person.in.Gender)

		i.Calc()

		if i.String() != person.out {
			t.Errorf("in: %s out: %s for %v", person.out, i.String(), person.in)
		}
	}
}
