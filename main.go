package main

import (
	"fmt"
	"os"
	"os/exec"
)

func KnowMyIP() {
	cmd := exec.Command("ip", "addr")

	out, err :=cmd.Output()

	if err != nil{
		fmt.Println(err.Error())
		os.Exit(0)
	}

	fmt.Println(string(out))

}

func PingAddress() {
	cmd := exec.Command("ping", "google.com")

	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil{
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func main() {
	// cmd := exec.Command("ping", "google.com")

	// cmd.Stdout = os.Stdout

	// err := cmd.Run()

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	KnowMyIP()
	PingAddress()
}