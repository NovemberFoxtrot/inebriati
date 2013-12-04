package inebriati

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	i := New(0.0, 0.0, 0.0, "male")

	if i.String() != "0e+00" {
		t.Errorf("%s", i.String())
	}
}

func TestCalc(t *testing.T) {
	i := New(0.0, 0.0, 0.0, "male")

	test := []struct {
		in  stats
		out string
	}{
		{stats{0.0, 0.0, 0.0, "male", 0.0}, "0e+00"},
		{stats{0.0, 0.0, 0.0, "female", 0.0}, "0e+00"},
	}

	for _, person := range test {
		fmt.Println(person.in, i)
		i := New(person.in.StandardDrinks, person.in.BodyWeightKiloGrams, person.in.DrinkingPeriodHours, person.in.Gender)
		if i.String() != person.out {
			t.Errorf("%s", person)
		}
	}
}
