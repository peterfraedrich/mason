package survey

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var supportedManagers = map[string]func() ([]Package, error){
	"Server Manager": getRoles,
	"Windows":        getApps,
	"chocolatey":     getChoco,
}

func GetPackages() (Data, error) {
	var data Data
	for pm, fn := range supportedManagers {
		p, err := fn()
		if err != nil {
			return data, err
		}
		if len(p) > 0 {
			data.PackageManagers = append(data.PackageManagers, pm)
		}
		for _, pkg := range p {
			data.Packages = append(data.Packages, pkg)
		}
	}
	fmt.Printf("%+v\n", data)
	return data, nil
}

func getRoles() ([]Package, error) {
	var pkgs []Package
	roles, err := doWMIQuery("$ProgressPreference = \"SilentlyContinue\"; get-windowsfeature | Where-Object{$_.installed -eq $true -and $_.featuretype -eq 'Role'} | ConvertTo-Json")
	if err != nil {
		return pkgs, err
	}
	var wmiRoles []WMIServerRole
	err = json.Unmarshal(roles, &wmiRoles)
	if err != nil {
		return pkgs, err
	}
	for _, r := range wmiRoles {
		p := Package{
			Name:      "(Server Role) " + r.Name,
			Version:   fmt.Sprintf("%v.%v", r.AdditionalInfo.MajorVersion, r.AdditionalInfo.MinorVersion),
			ManagedBy: "Windows Server Manager",
		}
		pkgs = append(pkgs, p)
	}
	return pkgs, nil
}

func getApps() ([]Package, error) {
	var pkgs []Package
	products, err := doWMIQuery("$ProgressPreference = \"SilentlyContinue\"; get-wmiobject -class win32_product | ConvertTo-Json")
	if err != nil {
		return pkgs, err
	}
	var prods []WMIProduct
	err = json.Unmarshal(products, &prods)
	if err != nil {
		return pkgs, err
	}
	for _, r := range prods {
		p := Package{
			Name:      r.Name,
			Version:   r.Version,
			ManagedBy: "Windows",
		}
		pkgs = append(pkgs, p)
	}
	return pkgs, nil
}

func getChoco() ([]Package, error) {
	var pkgs []Package
	_, err := os.Stat(`C:\ProgramData\chocolatey\bin\choco.exe`)
	if os.IsNotExist(err) {
		return pkgs, nil
	}
	if err != nil {
		return pkgs, err
	}
	var stdout bytes.Buffer
	proc := exec.Command("powershell.exe", "choco list -lr | ConvertTo-Json")
	proc.Stdout = &stdout
	err = proc.Run()
	if err != nil {
		return pkgs, err
	}
	sout, err := ioutil.ReadAll(&stdout)
	if err != nil {
		return pkgs, err
	}
	var lines []string
	err = json.Unmarshal(sout, &lines)
	if err != nil {
		return pkgs, err
	}
	for idx, i := range lines {
		if idx == 0 {
			continue
		}
		line := strings.Split(i, "|")
		if len(line) >= 2 {
			pkgs = append(pkgs, Package{
				Name:      line[0],
				Version:   line[1],
				ManagedBy: "chocolatey",
			})
		}
	}
	return pkgs, nil
}

func doWMIQuery(query string) ([]byte, error) {
	var stdout bytes.Buffer
	proc := exec.Command("powershell.exe", query)
	proc.Stdout = &stdout
	err := proc.Run()
	if err != nil {
		return []byte{}, err
	}
	sout, err := ioutil.ReadAll(&stdout)
	if err != nil {
		return []byte{}, err
	}
	return sout, nil
}

type WMIServerRole struct {
	Name                      string
	DisplayName               string
	Description               string
	Installed                 bool
	InstallState              int
	FeatureType               string
	Path                      string
	Depth                     int
	DependsOn                 []string
	Parent                    string
	ServerComponentDescriptor map[string]interface{}
	SubFeatures               []string
	SystemService             []interface{}
	Notification              []interface{}
	BestPracticesModelId      string
	EventQuery                string
	PostConfigurationNeeded   bool
	AdditionalInfo            struct {
		MajorVersion int
		MinorVersion int
		NumericId    int
		InstallName  string
	}
}

type WMIProduct struct {
	Name    string
	Version string
}
