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

func add(x float64,y float64) float64 {
	return x+y
}

func foo() {
	fmt.Println("The square root of 4 is", math.Sqrt(4))
}

func getRandom() {
	//random is funny in go, unless we redefine our range, it always gives us the same random variable
	fmt.Println("A number from 1-100", rand.Intn(100))
}

func main() {
	foo()
	getRandom()
	//define a variable:
	var num1 float64 = 5.6
	var num2 float64 = 9.5

	fmt.Println("this is what our add function does:", add(num1,num2))
}