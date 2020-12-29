package main

import (
	"github.com/StackExchange/wmi"
)

func getNameservers() []string {
	var res []Win32_NetworkAdapterConfiguration
	q := wmi.CreateQuery(&res, "")
	err := wmi.QueryNamespace(q, &res, `root\cimv2`)
	if err != nil {
		panic(err)
	}
	for _, n := range res {
		if len(n.DNSServerSearchOrder) != 0 {
			return n.DNSServerSearchOrder
		}
	}
	return []string{}
}
