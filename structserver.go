package main

import "fmt"

type car struct {
	gasPedal uint16 // anything from 0-65k (65535)
	brakePedal uint16
	steeringWheel int16 // -32k to +32k
	topSpeedKmh float64
}

func main(){
	// best way to do this:
	aCar := car{gasPedal: 22341, 
							brakePedal: 0, 
							steeringWheel: 12561, 
							topSpeedKmh: 225.0}
	// or do this:
	// aCar := car{22341,0,12562,225.0}

	fmt.Println(aCar.gasPedal)
}