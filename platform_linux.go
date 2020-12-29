package main

import (
	"bytes"
	"io/ioutil"
	"strings"
)

func getNameservers() []string {
	f, err := ioutil.ReadFile("/etc/resolv.conf")
	if err != nil {
		return []string{}
	}
	nameservers := []string{}
	lines := bytes.Split(f, []byte("\n"))
	for _, l := range lines {
		line := string(l[:])
		if strings.HasPrefix(line, "nameserver") {
			elems := strings.Split(line, " ")
			nameservers = append(nameservers, elems[1])
		}
	}
	return nameservers
}
