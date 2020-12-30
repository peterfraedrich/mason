package survey

import (
	"fmt"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
)

//DoSurvey function
func DoSurvey(getPackages bool) (map[string]interface{}, error) {
	numWorkers := 5
	done := make(chan bool, numWorkers)
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
		ip, mask := getPreferredIP()
		res["net"] = map[string]interface{}{
			"ipv4_private":         ip,
			"ipv4_private_netmask": mask,
			"ipv4_public":          getPublicIP(),
			"nameservers":          getNameservers(),
		}
		done <- true
	}()
	go func() {
		if getPackages {
			d, _ := GetPackages()
			res["packages"] = map[string]interface{}{
				"list":             d.Packages,
				"package_managers": d.PackageManagers,
			}
		}
		done <- true
	}()
	go func() {
		a, _ := GetAWS()
		res["cloud"] = map[string]interface{}{
			"aws": a,
		}
		done <- true
	}()
	for i := 0; i < numWorkers; i++ {
		<-done
	}
	return res, nil
}
