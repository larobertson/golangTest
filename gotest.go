package main

import ("fmt"
				"math"
				"math/rand")
// fmt is format

func foo() {
	fmt.Println("The square root of 4 is", math.Sqrt(4))
}

func main() {
	foo()
	fmt.Println("A number from 1-100", rand.Intn(100))
}