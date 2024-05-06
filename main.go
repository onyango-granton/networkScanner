package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
	// "time"
	// "strings"
)

func KnowMyIP() {
	cmd := exec.Command("ip", "addr")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	fmt.Println(string(out))
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



func PingMembersOfNet(n int) {
		cmd := exec.Command("ping", "192.168.89."+strconv.Itoa(n))

		// cmd.Stdout = customOutput{}
		cmd.Stdout = os.Stdout
		fmt.Println(cmd.Stdout)

		var d = 1000 * time.Microsecond
   		var t = time.Now().Add(d)

		for {
			for time.Now().Before(t){
				fmt.Println("Hello")
				cmd.Run()
			}
			break
		}
		// cmd.Run()

		fmt.Println("Here")

		// if err != nil {
		// 	fmt.Println(err.Error())
		// }

}

func main() {
    PingMembersOfNet(173)
}
