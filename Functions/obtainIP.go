package functions

import (
	"os"
	"strings"
)

// ObtainIP takes in a file as string argument and return a list of IP Addresses
func ObtainIP(s string) []string {
	// Read contents of file passed as argument
	out, _ := os.ReadFile(s)
	splitS := strings.Split(string(out), "\n")
	res := []string{}
	ipAddr := []string{}

	for _, ch := range splitS {
		ch = strings.Trim(ch, " ")
		res = append(res, ch)
	}

	for _, ch := range res {
		if strings.Contains(ch, "inet") {
			// Filters out ipv6 addresses
			if strings.Contains(ch, "inet6") {
				continue
			}
			// Filters out loopback address
			if strings.Contains(ch, "127.0.0.1") {
				continue
			}

			// Cleans up to obtain only the ip address and broadcast address
			ch = strings.TrimLeft(ch, "inet ")
			ch = strings.TrimRight(ch, " scope global dynamic noprefixroute wlp4s0")

			stringsSplit := strings.Split(ch, " ")
			// Appending only the ip address with subnet mask
			ipAddr = append(ipAddr, stringsSplit[0])

		}
	}

	return ipAddr
}
