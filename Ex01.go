// Exercise: Slices
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	//a slice of length dy
	p := make([][]uint8, dy)

	//each element of which is a slice of dx 8-bit unsigned integers
	for i := 0; i < dy; i++ {
		p[i] = make([]uint8, dx)

		for j := 0; j < dx; j++ {
			p[i][j] = uint8(i*i + j*j)
			// p[i][j] = uint8(i ^ j)
			// p[i][j] = uint8((i + j*i) / 2)
			// p[i][j] = uint8((i + j*i) / 2)
			// p[i][j] = uint8((i + j) / 2)
		}
	}
	return p
}

func main() {
	pic.Show(Pic)
}
