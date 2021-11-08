// Exercise: Maps

package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

//Returns a MAP of ints (count)
//Count == every word in string s
func WordCount(s string) map[string]int {

	count := make(map[string]int)
	words := strings.Fields(s)

	fmt.Println("----------------START----------------")
	fmt.Println("INPUT: ", words)
	fmt.Println("Counting words...")

	for index, value := range words {
		count[value] += 1
		fmt.Println("Index: ", index, " Value: ", value)
	}
	fmt.Println("Done!")
	fmt.Println("OUTPUT: ", count)
	fmt.Println("----------------END-----------------")
	return count
}
func main() {
	wc.Test(WordCount)
}
