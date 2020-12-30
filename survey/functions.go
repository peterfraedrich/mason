package survey

import (
	"encoding/json"
	"fmt"
	"net"
	"runtime"
	"strings"

	"github.com/chyeh/pubip"
	"github.com/clbanning/anyxml"
	"github.com/pelletier/go-toml"
	"github.com/rodaine/hclencoder"
	"github.com/shirou/gopsutil/cpu"
	"gopkg.in/yaml.v3"
)

//WriteOutput function
func WriteOutput(data interface{}, format string) error {
	d := []byte{}
	var err error
	switch strings.ToLower(format) {
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
	case "xml":
		d, err = anyxml.Xml(data)
		if err != nil {
			return err
		}
	case "toml":
		d, err = toml.Marshal(data)
		if err != nil {
			return err
		}
	case "hcl":
		d, err = hclencoder.Encode(data)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Format %s is not a valid format", format)
	}
	fmt.Printf("%s\n", d)
	return nil
}

func getPreferredIP() (string, string) {
	conn, _ := net.Dial("udp", "1.1.1.1:1")
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	sz, _ := localAddr.IP.DefaultMask().Size()
	return localAddr.IP.String(), fmt.Sprintf("/%v", sz)
}

func getPublicIP() string {
	ip, err := pubip.Get()
	if err != nil {
		return ""
	}
	return ip.String()
}

func getCPUCores(i []cpu.InfoStat) int {
	switch runtime.GOOS {
	case "linux":
		return len(i)
	case "windows":
		return int(i[0].Cores)
	default:
		return int(i[0].Cores)
	}
}
