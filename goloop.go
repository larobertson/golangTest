package main

import "fmt"

func main() {
	// a couple of different loop formats
	// Go only has the for loop

	for i:=0; i<10; i++ {
		fmt.Println(i)
	}

	// more commonly loops will look like the one below in Go
	j:=0
	for j < 10 {
		fmt.Println(j)
		j++
	}
	
}