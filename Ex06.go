package main

import (
	// "fmt"
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		// fmt.Println("-------------------------------------------")
		// fmt.Printf("b[%v] = %s", i, b[i])
		// fmt.Println("-------------------------------------------")
		b[i] = 65
		// fmt.Println("///////////////////////////////////////////")
		// fmt.Printf("b[%v] = %s", i, b[i])
		// fmt.Println("///////////////////////////////////////////")

	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
