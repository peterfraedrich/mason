package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

func writeOutput(data interface{}, format string) error {
	d := []byte{}
	var err error
	switch format {
	case "json":
		d, err = json.Marshal(data)
		if err != nil {
			return err
		}
	case "yaml":
		d, err = yaml.Marshal(data)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Format %s is not a valid format", format)
	}
	fmt.Printf("%s\n", d)
	return nil
}
