package inebriati

import (
	"flag"
	"fmt"
)

const (
	BodyWaterInTheBlood     = 0.806
	BodyWaterMen            = 0.58
	BodyWaterWomen          = 0.49
	Metabolism              = 0.017
	GramsToSwedishStandards = 1.2
)

var (
	StandardDrinks      = flag.Float64("drinks", 2.5, "How many drinks consumed")
	BodyWeightKiloGrams = flag.Float64("weight", 70.0, "Body weight in kilograms")
	DrinkingPeriodHours = flag.Float64("hours", 2.0, "How many hours it took to drink")
	Gender              = flag.String("gender", "male", "Specify male or female")
)

func calc() float64 {
	bodyWater := BodyWaterMen

	if *Gender != "male" {
		bodyWater = BodyWaterWomen
		*Gender = "female"
	}

	return ((BodyWaterInTheBlood * *StandardDrinks * GramsToSwedishStandards) / (bodyWater * *BodyWeightKiloGrams)) - (Metabolism * *DrinkingPeriodHours)
}

func main() {
	flag.Parse()

	EstimatedBloodEthanolConcentration := calc()

	fmt.Println(EstimatedBloodEthanolConcentration, "g/dL", *Gender)
}
