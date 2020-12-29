package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/shirou/gopsutil/net"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
)

func doSurvey() (map[string]interface{}, error) {
	done := make(chan bool, 3)
	res := map[string]interface{}{}
	go func() {
		host, err := host.Info()
		if err != nil {
			fmt.Println(err)
		}
		res["host"] = map[string]interface{}{
			"hostname":       host.Hostname,
			"uptime_seconds": host.Uptime,
			"boot_time_unix": host.BootTime,
			"boot_time_iso":  time.Unix(int64(host.BootTime), 0).Format(time.RFC3339),
			"os_type":        host.OS,
			"os_name":        host.Platform,
			"os_family":      host.PlatformFamily,
			"os_version":     host.PlatformVersion,
			"kernel_version": host.KernelVersion,
			"architecture":   host.KernelArch,
			"host_id":        host.HostID,
		}
		done <- true
	}()
	go func() {
		c, err := cpu.Info()
		if err != nil {
			fmt.Println(err)
		}
		res["cpu"] = map[string]interface{}{
			"vendor":     c[0].VendorID,
			"model":      strings.TrimSpace(c[0].ModelName),
			"mhz":        c[0].Mhz,
			"cache_size": c[0].CacheSize,
			"cores":      getCPUCores(c),
		}
		done <- true
	}()
	go func() {
		_, err := net.Interfaces()
		if err != nil {
			fmt.Println(err)
		}
		ip, mask := getPreferredIP()
		res["net"] = map[string]interface{}{
			"ipv4_private":         ip,
			"ipv4_private_netmask": mask,
			"ipv4_public":          getPublicIP(),
			"nameservers":          getNameservers(),
		}
		done <- true
	}()
	for i := 0; i < 3; i++ {
		<-done
	}
	return res, nil
}
