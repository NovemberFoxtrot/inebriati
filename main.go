package main

import (
	"flag"
	"fmt"
	"github.com/NovemberFoxtrot/inebriati/calc"
)

var (
	StandardDrinks      = flag.Float64("drinks", 2.5, "How many drinks consumed")
	BodyWeightKiloGrams = flag.Float64("weight", 70.0, "Body weight in kilograms")
	DrinkingPeriodHours = flag.Float64("hours", 2.0, "How many hours it took to drink")
	Gender              = flag.String("gender", "male", "Specify male or female")
)

func main() {
	flag.Parse()

	i := inebriati.New(*StandardDrinks,*BodyWightKiloGrams,*DrinkingPeriod,*Gender)
	i.Calc()

	fmt.Println(i)
}
