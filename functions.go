package main

import (
	"encoding/json"
	"fmt"
	"net"
	"runtime"

	"github.com/chyeh/pubip"
	"github.com/shirou/gopsutil/cpu"
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
