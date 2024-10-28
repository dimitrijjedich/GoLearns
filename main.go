package main

import (
	"fmt"
	"golang.org/x/tour/pic"
)

func sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z = z - ((z*z - x) / (2 * z))
	}
	return z
}

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for y := range result {
		result[y] = make([]uint8, dx)
		for x := range result[y] {
			//result[y][x] = uint8((x + y) / 2)
			//result[y][x] = uint8(x * y)
			result[y][x] = uint8(x ^ y)
		}
	}
	return result
}

func main() {
	fmt.Println("Hello World")
	pic.Show(Pic)
}
