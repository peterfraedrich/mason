package survey

import (
	"bytes"
	"os/exec"
	"strings"
)

//Get func
func GetPackages() (Data, error) {
	var data Data
	lines, err := getBrew()
	if err != nil {
		return data, err
	}
	data.Packages = lines
	data.PackageManagers = []string{"homebrew"}
	return data, nil
}

func getBrew() ([]Package, error) {
	var lines []Package
	cmd := exec.Command("brew", "list", "--versions")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return lines, err
	}
	l := out.String()
	lns := strings.Split(l, "\n")
	for _, line := range lns {
		if len(line) < 1 {
			continue
		}
		pkg := strings.Split(line, " ")
		lines = append(lines, Package{
			Name:    pkg[0],
			Version: pkg[len(pkg)-1],
		})
	}
	return lines, nil
}
