package main

import (
	// "context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	// "time"
	// "strings"
)

// KnowMyIp returns IP address of host Machine in lias with ObtainIP function
// This function uses ObtainIP function
func KnowMyIP() []string {
	// cmd executes ip addr command
	// in linux this shows the network address across interface(s)
	cmd := exec.Command("ip", "addr")

	// output from the exec.Command is stored
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	// output is then written to an output text file
	// output.txt has been manualy edited to hide sensitive info
	err1 := os.WriteFile("output.txt", out, 0o644)
	if err1 != nil {
		fmt.Println(err1.Error())
	}

	return ObtainIP("output.txt")
}

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

// PingAddress pings a website using ping command in linux
func PingAddress() {
	cmd := exec.Command("ping", "google.com")

	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

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
	f, err := os.Create("new.txt")
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

// ObtainNetAddress takes in ip address as string xxx.xxx.xxx.xxx returns
// netAddress xxx.xxx.xxx
func ObtainNetAddress(s1 string) string {
	// s1 := "192.168.79.78"
	// splits a string "xxx.xxx.xxx.xxx" by "."
	splitString := strings.Split(s1, ".")
	// network address is stored as first three parts of splitString joined by "."
	netAddr := strings.Join(splitString[:3], ".")
	
	// return an address of "xxx.xxx.xxx"
	return netAddr
}


func main() {
	PingNetworkMembers()
}
