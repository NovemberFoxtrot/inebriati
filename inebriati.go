package main

import (
	"fmt"
)

const (
	BodyWaterInTheBlood     = 0.806
	BodyWaterMen            = 0.58
	BodyWaterWomen          = 0.49
	Metabolism              = 0.017
	GramsToSwedishStandards = 1.2
)

func main() {
	var StandardDrinks float64
	var BodyWeightKiloGrams float64
	var DrinkingPeriodHours float64

	StandardDrinks = 2.5
	BodyWeightKiloGrams = 70.0
	DrinkingPeriodHours = 2.0

	EstimatedBloodEthanolConcentration := ((BodyWaterInTheBlood * StandardDrinks * GramsToSwedishStandards) / (BodyWaterWomen * BodyWeightKiloGrams)) - (Metabolism * DrinkingPeriodHours)

	fmt.Println(EstimatedBloodEthanolConcentration, "g/dL")
}
