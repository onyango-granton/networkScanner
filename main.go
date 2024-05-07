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

func PingNetworkMembers() {
	stringIPAddr := []string{}
	stringSubNet := []string{}
	for _, ch := range KnowMyIP() {
		stringsSplit := strings.Split(ch, "/")
		// fmt.Println(stringsSplit[0], "IP Addr")
		stringIPAddr = append(stringIPAddr, stringsSplit[0])
		// fmt.Println(stringsSplit[1], "SubnetMask")
		stringSubNet = append(stringSubNet, stringsSplit[1])
	}
	fmt.Println(stringIPAddr)
	fmt.Println(stringSubNet)

	f, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, ch := range stringIPAddr {
		for i := 1; i < 100; i++ {
			cmd := exec.Command("ping", "-c", "1", ObtainNetAddress(ch)+"."+strconv.Itoa(i))
			cmd.Stdout = f

			err := cmd.Run()
			if err != nil {
				fmt.Println(err.Error())
			}

		}
	}
}

func ObtainNetAddress(s1 string) string {
	// s1 := "192.168.79.78"
	splitString := strings.Split(s1, ".")
	netAddr := strings.Join(splitString[:3], ".")
	// fmt.Println(netAddr)
	return netAddr
}

func PingMembersOfNet(n int) {
	cmd := exec.Command("ping", "-c", "2", "192.168.89."+strconv.Itoa(n))

	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("END")
}

func main() {
	// ctx := context.Background()
	// // The context now times out after 1 second
	// // alternately, we can call `cancel()` to terminate immediately
	// ctx, _ = context.WithTimeout(ctx, 7*time.Second)

	// cmd := exec.CommandContext(ctx, "ping", "192.268.0.33")

	// out, err := cmd.Output()
	// if err != nil {
	//   fmt.Println("could not run command: ", err)
	// }
	// fmt.Println("Output: ", string(out))
	// for i := 1; i < 255; i++{
	// 	PingMembersOfNet(i)
	// }
	// PingMembersOfNet(173)
	PingNetworkMembers()
}
