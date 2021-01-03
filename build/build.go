package build

import "gopkg.in/yaml.v3"

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

	//execute

	//// render

	//// run

	//// results

	//results

	return nil
}
