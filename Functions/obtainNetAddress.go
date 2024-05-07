package functions

import "strings"

// ObtainNetAddress takes in ip address as string xxx.xxx.xxx.xxx returns
// netAddress xxx.xxx.xxx
func ObtainNetAddress(s1 string) string {
	// s1 := "192.168.79.78"
	// splits a string "xxx.xxx.xxx.xxx" by "."
	splitString := strings.Split(s1, ".")
	// network address is stored as first three parts of splitString joined by "."
	netAddr := strings.Join(splitString[:3], ".")

	// return an address of "xxx.xxx.xxx"
	return netAddr
}
