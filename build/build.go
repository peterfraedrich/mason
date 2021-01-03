package build

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

//GoBuild function
func GoBuild(blueprint []byte) error {
	//parse
	var bp Blueprint
	err := yaml.Unmarshal(blueprint, &bp)
	if err != nil {
		return err
	}

	//validate
	err = bp.Validate()
	if err != nil {
		return err
	}

	var done chan error
	if bp.Parallel {
		done = make(chan error, len(bp.Steps))
	}

	//execute
	for _, step := range bp.Steps {
		// validate
		err = step.Validate()
		if err != nil {
			return err
		}

		// render

		// run
		if bp.Parallel {
			go func() {
				err = step.Do()
				done <- err
			}()
		} else {
			err = step.Do()
			if err != nil {
				return err
			}
		}
		// results
		if !bp.Parallel {
			res, err := step.Results()
			if err != nil {
				return err
			}
			PrintResults(res)
		}

	}
	if bp.Parallel {
		for e := range done {
			if e != nil {
				fmt.Println(err)
			}
		}
		return fmt.Errorf("there were errors executing the blueprint")
	}

	//results

	return nil
}
