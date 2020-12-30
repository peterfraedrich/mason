package main

import (
	"github.com/peterfraedrich/mason/survey"
	"github.com/thatisuday/commando"
)

func surveyHandler(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	outForm := flags["output"].Value.(string)
	res, err := survey.DoSurvey(flags["packages"].Value.(bool))
	if err != nil {
		panic(err)
	}
	err = survey.WriteOutput(res, outForm)
	if err != nil {
		panic(err)
	}
}
