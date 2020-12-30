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
		AddFlag("output,o", "sets the output type (json, yaml, xml, toml, hcl)", commando.String, "json").
		AddFlag("packages,p", "whether to collect package info", commando.Bool, true).
		SetAction(surveyHandler)

	commando.Register("build").
		SetDescription("Builds the given blueprint").
		SetShortDescription("Builds a blueprint").
		AddFlag("blueprint,b", "which blueprint to build", commando.String, "blueprint.yaml").
		SetAction(buildHandler)

	commando.Parse(nil)
}
