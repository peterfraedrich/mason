package main

import (
	"io/ioutil"

	"github.com/peterfraedrich/mason/build"
	"github.com/peterfraedrich/mason/survey"
	"github.com/thatisuday/commando"
)

func surveyHandler(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	outForm := flags["output"].Value.(string)
	res, err := survey.DoSurvey(flags["packages"].Value.(bool))
	if err != nil {
		handleError(err, "surveyHandler.DoSurvey")
		return
	}
	err = survey.WriteOutput(res, outForm)
	if err != nil {
		handleError(err, "surveyHandler.WriteOutput")
		return
	}
}

func buildHandler(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	bpFile := flags["blueprint"].Value.(string)
	bp, err := ioutil.ReadFile(bpFile)
	if err != nil {
		handleError(err, "buildHandler.ReadFile")
		return
	}
	err = build.GoBuild(bp)
	if err != nil {
		handleError(err, "buildHandler.GoBuild")
		return
	}
}
