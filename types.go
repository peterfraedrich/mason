package main

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/net"
)

// SurveyResults holds data for the survey
type SurveyResults struct {
	Host *host.InfoStat      `json:"host"`
	CPU  cpu.InfoStat        `json:"cpu"`
	Net  []net.InterfaceStat `json:"net"`
}
