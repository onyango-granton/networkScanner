package functions

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// PingNetworkMembers sends data packets to members of a network
// This function uses KnowMyIp function (and) ObtainNetAddress function
func PingNetworkMembers() {
	// stringIPAddr is a list of obtained IP addresses without subnet mask
	stringIPAddr := []string{}
	// stringSubNet is a list of subnet masks without IP addresses
	stringSubNet := []string{}

	// looping through list of IPS
	for _, ch := range KnowMyIP() {
		// separating IP Address and Subnet Mask appending them
		stringsSplit := strings.Split(ch, "/")
		stringIPAddr = append(stringIPAddr, stringsSplit[0])
		stringSubNet = append(stringSubNet, stringsSplit[1])
	}

	// creating a file called neew.txt to store output(ip addresses)
	f, err := os.Create("Output Files/output.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// looping through list of ip addresses
	for _, ch := range stringIPAddr {
		// looping through members of network per ip address
		for i := 1; i < 100; i++ {
			cmd := exec.Command("ping", "-c", "1", ObtainNetAddress(ch)+"."+strconv.Itoa(i))
			// store command data to out new.txt file
			cmd.Stdout = f

			err := cmd.Run()
			if err != nil {
				fmt.Println(err.Error())
			}

		}
	}
}
