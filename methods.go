package main

import "fmt"

// two types of of methods
	// value receivers
		// just receives values (doesn't need to modify or change anything)
	// pointer receivers

const usixteenbitmax float64 = 65535
const kmhMultiple float64 = 1.60934

	type car struct {
		gasPedal uint16 // anything from 0-65k (65535)
		brakePedal uint16
		steeringWheel int16 // -32k to +32k
		topSpeedKmh float64
	}

	func (c car) kmh() float64 {
		// saying type car associates it with the car struct
		return float64(c.gasPedal) * (c.topSpeedKmh/usixteenbitmax)
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
		fmt.Println(aCar.kmh())
	}
