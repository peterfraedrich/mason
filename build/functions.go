package build

import "fmt"

func runCommand(shell string, command string, arguments []string) (stdout string, stderr string, err error) {

	return "", "", nil
}

//PrintResults prints results
func PrintResults(res StepResults) error {
	fmt.Printf("%+v\n", res)
	return nil
}
