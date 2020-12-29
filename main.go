package main

import (
	"github.com/thatisuday/commando"
)

func main() {
	commando.SetExecutableName("mason").
		SetVersion("v0.0.1").
		SetDescription("Mason builds machines")

	commando.Register("survey").
		SetDescription("Surveys the machine and returns data").
		SetShortDescription("displays system information").
		AddFlag("output,o", "sets the output type (json, yaml)", commando.String, "json").
		SetAction(surveyHandler)

	commando.Parse(nil)
}
