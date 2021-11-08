package main

import (
	"fmt"
)

func fibonacci() func() int {
	f0 := 0
	f1 := 1
	return func() int {
		fn := f0
		f0 = f1
		f1 = f1 + fn
		return fn
	}
}

func main() {

	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
