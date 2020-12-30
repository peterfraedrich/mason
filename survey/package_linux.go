package survey

import (
	"regexp"
	"strings"

	"github.com/go-cmd/cmd"
)

// supported managers maps string package manager name -> function to discover
var supportedManagers = map[string]func() ([]Package, error){
	"apk":    getAPK,
	"dnf":    getDNF,
	"dpkg":   getDpkg,
	"pacman": getPacman,
	"yum":    getYum,
	"zypper": getZypper,
}

//GetPackages func
func GetPackages() (Data, error) {
	var data Data
	for pm, fn := range supportedManagers {
		if !exists(pm) {
			continue
		}
		data.PackageManagers = append(data.PackageManagers, pm)
		p, err := fn()
		if err != nil {
			return data, err
		}
		for _, pkg := range p {
			data.Packages = append(data.Packages, pkg)
		}
	}
	return data, nil
}

func exists(manager string) bool {
	c := cmd.NewCmd("which", manager)
	status := <-c.Start()
	if status.Exit != 0 {
		return false
	}
	return true
}

func getDpkg() ([]Package, error) {
	var lines []Package
	c := cmd.NewCmd("dpkg", "-l")
	status := <-c.Start()
	for idx, i := range status.Stdout {
		if idx > 4 {
			line := regexp.MustCompile(`\W{2,}`).Split(i, -1)
			if len(line) >= 3 {
				lines = append(lines, Package{
					Name:      line[1],
					Version:   line[2],
					ManagedBy: "dpkg",
				})
			}
		}
	}
	return lines, nil
}

func getYum() ([]Package, error) {
	var lines []Package
	c := cmd.NewCmd("yum", "list", "installed", "-q")
	status := <-c.Start()
	for idx, i := range status.Stdout {
		if idx != 0 {
			line := regexp.MustCompile(`\W{2,}`).Split(i, -1)
			if len(line) >= 2 {
				lines = append(lines, Package{
					Name:      line[0],
					Version:   line[1],
					ManagedBy: "yum",
				})
			}
		}
	}
	return lines, nil
}

func getAPK() ([]Package, error) {
	var lines []Package
	c := cmd.NewCmd("apk", "list", "-q")
	status := <-c.Start()
	for _, i := range status.Stdout {
		if !strings.HasPrefix(i, "WARNING") {
			spl := strings.Split(i, "-")
			if len(spl) >= 2 {
				lines = append(lines, Package{
					Name:      spl[0],
					Version:   spl[1],
					ManagedBy: "apk",
				})
			}
		}
	}
	return lines, nil
}

func getZypper() ([]Package, error) {
	var lines []Package
	c := cmd.NewCmd("rpm", "-q -a --queryformat \"%{NAME}==%{VERSION}\n\"")
	status := <-c.Start()
	for _, i := range status.Stdout {
		spl := strings.Split(i, "==")
		if len(spl) >= 2 {
			lines = append(lines, Package{
				Name:      spl[0],
				Version:   spl[1],
				ManagedBy: "zypper",
			})
		}
	}
	return lines, nil
}

func getDNF() ([]Package, error) {
	var lines []Package
	c := cmd.NewCmd("rpm", "-q -a --queryformat \"%{NAME}==%{VERSION}\n\"")
	status := <-c.Start()
	for _, i := range status.Stdout {
		spl := strings.Split(i, "==")
		if len(spl) >= 2 {
			lines = append(lines, Package{
				Name:      spl[0],
				Version:   spl[1],
				ManagedBy: "dnf",
			})
		}
	}
	return lines, nil
}

func getPacman() ([]Package, error) {
	var lines []Package
	c := cmd.NewCmd("pacman", "-Qe")
	status := <-c.Start()
	for _, i := range status.Stdout {
		spl := strings.Split(i, " ")
		if len(spl) >= 2 {
			lines = append(lines, Package{
				Name:      spl[0],
				Version:   spl[1],
				ManagedBy: "pacman",
			})
		}
	}
	return lines, nil
}

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
