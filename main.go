package main

import (
	// "context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	// "time"
	// "strings"
)

func KnowMyIP() []string{
	cmd := exec.Command("ip", "addr")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	// fmt.Println(string(out))

	err1 := os.WriteFile("output.txt", out, 0644)
	if err1 != nil{
		fmt.Println(err1.Error())
	}

	// for _,ch := range obtainIP("output.txt"){
	// 	fmt.Println(ch)
	// }
	return obtainIP("output.txt")
}

func obtainIP(s string) []string {
	out, _ := os.ReadFile(s)
	splitS := strings.Split(string(out), "\n")
	res := []string{}
	// inet := []string{}
	ipAddr := []string{}

	for _, ch := range splitS{
		ch = strings.Trim(ch," ")
		res = append(res, ch)
	}

	for _, ch := range res{
		if strings.Contains(ch, "inet"){
			if strings.Contains(ch, "inet6"){
				continue
			}
			if strings.Contains(ch, "127.0.0.1"){
				continue
			}
			ch = strings.TrimLeft(ch, "inet ")
			ch = strings.TrimRight(ch, " scope global dynamic noprefixroute wlp4s0")
			// inet = append(inet, ch)

			stringsSplit := strings.Split(ch, " ")
			ipAddr = append(ipAddr, stringsSplit[0])

			// for k, num := range stringsSplit[0]{

			// }
		}
	}

	// for _, ch := range ipAddr{
	// 	re ,_ := strconv.Atoi(ch)
	// 	fmt.Println(re)
	// }

	return ipAddr

}

func PingAddress() {
	cmd := exec.Command("ping", "google.com")

	// startTime := time.Now()
	

	cmd.Stdout = os.Stdout
	startTime := time.Now()

	fmt.Println("Here")

	err := cmd.Run()

	duration := time.Since(startTime).Seconds()

	if int(duration) == 2{
		fmt.Println("Here")
	}

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

type customOutput struct{}

func (c customOutput) Write(p []byte) (int, error) {
	fmt.Println("received output: ", string(p))
	return len(p), nil
}


func PingNetworkMembers() {
	stringIPAddr := []string{}
	stringSubNet := []string{}
	for _, ch := range KnowMyIP(){
		stringsSplit := strings.Split(ch, "/")
		// fmt.Println(stringsSplit[0], "IP Addr")
		stringIPAddr = append(stringIPAddr, stringsSplit[0])
		// fmt.Println(stringsSplit[1], "SubnetMask")
		stringSubNet = append(stringSubNet, stringsSplit[1])
	}
}



func PingMembersOfNet(n int) {
		cmd := exec.Command("ping", "-c","2","192.168.89."+strconv.Itoa(n))

		cmd.Stdout = os.Stdout

		err := cmd.Run()

		if err != nil{
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