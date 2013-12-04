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

type stats struct {
	StandardDrinks                     float64
	BodyWeightKiloGrams                float64
	DrinkingPeriodHours                float64
	Gender                             float64
	EstimatedBloodEthanolConcentration float64
}

func (s *stats) calc() float64 {
	bodyWater := BodyWaterMen

	if s.Gender != "male" {
		bodyWater = BodyWaterWomen
		s.Gender = "female"
	}

	return ((BodyWaterInTheBlood * s.StandardDrinks * GramsToSwedishStandards) / (bodyWater * s.BodyWeightKiloGrams)) - (Metabolism * s.DrinkingPeriodHours)
}

func (s *stats) String() string {
	fmt.Println(s.EstimatedBloodEthanolConcentration, "g/dL", s.Gender)
}
