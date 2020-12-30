package main

import (
	"fmt"
	"log"
	"os"
)

// This is temporary just to keep us from panicking everywhere

func handleError(err error, source string) {
	if err != nil {
		logString := fmt.Sprintf("[ %s ] %s\n", source, err.Error())
		log.Fatal(logString)
	}
	os.Exit(1)
}
