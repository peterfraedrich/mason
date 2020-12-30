package survey

import (
	"fmt"
	"strings"
)

//Win32_NetworkAdapterConfiguration type
type Win32_NetworkAdapterConfiguration struct {
	DNSServerSearchOrder []string
}

//Package type
type Package struct {
	Name      string `json:"name" msgpack:"name"`
	Version   string `json:"version" msgpack:"version"`
	ManagedBy string `json:"managed_by" msgpack:"managed_by"`
}

//Data type
type Data struct {
	Packages        PackageList `json:"packages" msgpack:"packages"`
	PackageManagers []string    `json:"package_managers" msgpack:"package_managers"`
}

//PackageList type
type PackageList []Package

//NameVersionEvalForm method
func (pl PackageList) NameVersionEvalForm() string {
	out := make([]string, len(pl))
	for i, v := range pl {
		if v.Version == "" {
			out[i] = fmt.Sprintf(`''%s''`, v.Name)
		} else {
			out[i] = fmt.Sprintf(`''%s-%s''`, v.Name, v.Version)
		}
	}
	return fmt.Sprintf("(%s)", strings.Join(out, ", "))
}

//NameEvalForm method
func (pl PackageList) NameEvalForm() string {
	out := make([]string, len(pl))
	for i, v := range pl {
		out[i] = fmt.Sprintf(`''%s''`, v.Name)
	}
	return fmt.Sprintf("(%s)", strings.Join(out, ", "))
}
