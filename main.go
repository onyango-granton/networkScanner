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

func PingMembersOfNet() {
	startTime := time.Now()
	duration := int(time.Since(startTime).Seconds())
	for i := 170; i < 173; i++ {
		cmd := exec.Command("ping", "192.168.89."+strconv.Itoa(i))

		cmd.Stdout = os.Stdout

		duration = int(time.Since(startTime).Seconds())

		if duration == 2{
			fmt.Println("Here")
		}

		err := cmd.Run()

		fmt.Println("Here")

		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func main() {
	// cmd := exec.Command("ping", "google.com")

	// cmd.Stdout = os.Stdout

	// err := cmd.Run()

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// KnowMyIP()
	// PingAddress()
	PingMembersOfNet()
}
