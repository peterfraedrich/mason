package main

import (
	"github.com/thatisuday/commando"
)

func surveyHandler(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	outForm := flags["output"].Value.(string)
	res, err := doSurvey()
	if err != nil {
		panic(err)
	}
	err = writeOutput(res, outForm)
	if err != nil {
		panic(err)
	}
}
