package main

import (
	"fmt"
	"math"
)

func main() {

	x, y := 3, 4

	var f float64 = math.Sqrt(float64(x*x + y*y))

	var z uint = uint(f)

	fmt.Println(x, y, z)

}
