package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (address IPAddr) String() string {
	// output := ""
	// for i := 0; i < 4; i++ {
	// 	output += fmt.Sprintf("%d.", address[i])
	// }
	// fmt.Println(output)
	// return output
	return fmt.Sprintf("%d.%d.%d.%d", address[0], address[1], address[2], address[3])
}
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
