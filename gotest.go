package main

// unused imports will cause an error
import ("fmt"
				"math"
				"math/rand")
// fmt is format
// to import packages inside of packages we use / notation (opposed to dot)
// to use math & rand we have to import both separately

// use go doc to learn more about a package
	// example, run in terminal: go doc fmt
			// or go doc math/rand Intn

func add(x,y float64) float64 {
	return x+y
}

func foo() {
	fmt.Println("The square root of 4 is", math.Sqrt(4))
}

func getRandom() {
	// random is funny in go, unless we redefine our range, it always gives us the same random variable
	fmt.Println("A number from 1-100", rand.Intn(100))
}

func multiple(a,b string) (string, string) {
	// we have to specify every return type, even if they are the same return type
	return a,b
}

func tutThree() {
		// define a variable:
		// previously:
		// var num1,num2 float64 = 5.6, 9.5
		num1,num2 := 5.6, 9.5
		// when inside a function we don't have to necessarily give it a type
		// but the type can't change once it compiles!
		// variables stay ONE type
		// however, to change later on we need to make sure we declare type
	
		fmt.Println("this is what our add function does:", add(num1,num2))
		// program will not work with unused variables
	
		w1, w2 := "Hey", "there"
	
		fmt.Println(multiple(w1,w2))
	
		// to convert a type:
		var a int = 62
		var b float64 = float64(a)
	
		x := a // x will be type int
		fmt.Println("just to use some variables, both are = to a but diff data types!", b, x)
}

func tutFour() {
	// learning about pointers

	x := 15
	a := &x // memory address <--- point to something use &
	// to read through use * (read what is at that memory address)
	fmt.Println(a)
	fmt.Println(*a)
	*a = 5
	fmt.Println("reading through the memory address at a",x)
	*a = *a**a
	fmt.Println("now what is a?", x)
	fmt.Println("and a?", *a)
}

func main() {
	// foo()
	// getRandom()
	// tutThree()
	tutFour()
}