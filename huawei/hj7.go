package main

import (
	"fmt"
	"math"
)

func main() {
	var num float32
	fmt.Scanf("%f", &num)
	num += 0.5
	fmt.Println(math.Floor(float64(num)))
}
