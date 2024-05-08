package functions

import (
	"fmt"
	"os"
	"strings"
)

func CleanOutput() {
	PingNetworkMembers()

	testFile := "Output Files/output.txt"

	res, _ := os.ReadFile(testFile)

	splitString := strings.Split(string(res), "\n")

	resArr := []string{}

	clients := []string{}

	for _, ch := range splitString {
		if ch == "" {
			continue
		}
		resArr = append(resArr, ch)
		// fmt.Println(ch, i)
	}

	count := 0

	for i := 0; i < len(resArr); i++ {
		if strings.Contains(resArr[i], "---") {
			trimLeft := strings.TrimLeft(resArr[i], "--- ")
			trimRight := strings.TrimRight(trimLeft, " ping statistics ---")
			if i+3 < len(resArr) && strings.Contains(resArr[i+3], "Destination Host Unreachable") {
				count++
				// fmt.Println("Offline	", trimRight)
				clients = append(clients, "Offline	"+trimRight)
			} else {
				count++
				// fmt.Println("Online	", trimRight)
				clients = append(clients, "Online	"+trimRight)
			}
		}
	}

	for j := 0; j < len(clients); j++ {
		for i := 0; i < len(clients); i++ {
			if i+1 < len(clients) && clients[i] > clients[i+1] {
				x := clients[i]
				clients[i] = clients[i+1]
				clients[i+1] = x
			}
		}
	}

	onlineClients := 0
	offlineClients := 0

	for i := 0; i < len(clients); i++ {
		if strings.Contains(clients[i], "Online") {
			onlineClients++
		}
		if strings.Contains(clients[i], "Offline") {
			offlineClients++
		}
		// fmt.Println(clients[i])
	}

	total := onlineClients + offlineClients

	percentOnline := float64(onlineClients) / float64(total) * 100
	// percentOffline := offlineClients/total * 100

	fmt.Printf("Scanned Address: %v\nOnline Clients: %v\nFree Addresses: %v\nPercent Online: %.2f\n", total, onlineClients, offlineClients, percentOnline)
}
