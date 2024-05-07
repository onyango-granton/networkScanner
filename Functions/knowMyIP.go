package functions

import (
	"fmt"
	"os"
	"os/exec"
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
