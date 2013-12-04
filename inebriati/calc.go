package inebriati

import (
	"strconv"
)

const (
	BodyWaterInTheBlood     = 0.806
	BodyWaterMen            = 0.58
	BodyWaterWomen          = 0.49
	Metabolism              = 0.015
	GramsToSwedishStandards = 1.2
)

type stats struct {
	StandardDrinks                     float64
	BodyWeightKiloGrams                float64
	DrinkingPeriodHours                float64
	Gender                             string
	EstimatedBloodEthanolConcentration float64
}

func New(sd, bwkg, dph float64, g string) *stats {
	return &stats{sd, bwkg, dph, g, 0.0}
}

func (s *stats) Calc() float64 {
	bodyWater := BodyWaterMen
	metabolism := Metabolism

	if s.Gender != "male" {
		bodyWater = BodyWaterWomen
		metabolism = 0.017
		s.Gender = "female"
	}

	s.EstimatedBloodEthanolConcentration = (BodyWaterInTheBlood*s.StandardDrinks*GramsToSwedishStandards)/(bodyWater*s.BodyWeightKiloGrams) - (metabolism * s.DrinkingPeriodHours)

	return s.EstimatedBloodEthanolConcentration
}

func (s *stats) String() string {
	return strconv.FormatFloat(s.EstimatedBloodEthanolConcentration, 'f', -1, 32) + " g/dL"
}
