package functions

import (
	"fmt"
	"os"
	"os/exec"
)

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
