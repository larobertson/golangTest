package main

import ("fmt"
				"math"
				"math/rand")
// fmt is format
// to import packages inside of packages we use / notation (opposed to dot)
// to use math & rand we have to import both separately

// use go doc to learn more about a package
	// example, run in terminal: go doc fmt
			//or go doc math/rand Intn

func foo() {
	fmt.Println("The square root of 4 is", math.Sqrt(4))
}

func main() {
	foo()
	fmt.Println("A number from 1-100", rand.Intn(100))
}