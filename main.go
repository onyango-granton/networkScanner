package main

import (
	// "context"
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
	for i := 1; i < 255; i++{
		PingMembersOfNet(i)
	}
	// PingMembersOfNet(173)
}