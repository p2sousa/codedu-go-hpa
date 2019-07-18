package main

import (
	"fmt"
	"math"
)

func loop() float64 {
	x := 0.0001
	for index := 0; index < 100000000; index++ {
		x = x + math.Sqrt(x)
	}
	return x
}

func main() {
	loop()
	fmt.Println("Code.education Rocks!")
}