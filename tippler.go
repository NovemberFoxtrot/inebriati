package main

import (
	"flag"
	"fmt"
	"github.com/NovemberFoxtrot/inebriati/inebriati"
)

func main() {
	drinks := flag.Float64("drinks", 2.5, "How many drinks consumed")
	weight := flag.Float64("weight", 70.0, "Body weight in kilograms")
	hours := flag.Float64("hours", 2.0, "How many hours it took to drink")
	gender := flag.String("gender", "male", "Specify male or female")

	flag.Usage = func() {
		fmt.Println(`
  NAME:

    inebriati

  DESCRIPTION:

    Calculates estimated blood ethanol concentration (EBAC) 

  COMMANDS:

    -drinks              Number of standard drinks comsumed.
    -weight              Weight in kiligrams.
    -hours               Drinking period in hours.
    -gender              Select male or female.
`[1:])
	}

	flag.Parse()

	i := inebriati.New(*drinks, *weight, *hours, *gender)
	i.Calc()

	fmt.Println(i)
}
