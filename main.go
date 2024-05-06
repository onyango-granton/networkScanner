package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ping", "google.com")

	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
	}
}