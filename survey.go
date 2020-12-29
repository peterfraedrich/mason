package main

import (
	"fmt"

	"github.com/shirou/gopsutil/net"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
)

func doSurvey() (SurveyResults, error) {
	done := make(chan bool, 3)
	res := &SurveyResults{}
	var err error
	go func() {
		res.Host, err = host.Info()
		if err != nil {
			fmt.Println(err)
		}
		done <- true
	}()
	go func() {
		c, err := cpu.Info()
		if err != nil {
			fmt.Println(err)
		}
		c[0].Flags = nil
		res.CPU = c[0]
		done <- true
	}()
	go func() {
		res.Net, err = net.Interfaces()
		if err != nil {
			fmt.Println(err)
		}
		done <- true
	}()
	for i := 0; i < 3; i++ {
		<-done
	}
	return *res, nil
}
