package main

import "fmt"

// two types of of methods
	// value receivers
		// just receives values (doesn't need to modify or change anything)
		// essentially makes a copy, more useful for small structs
	// pointer receivers
		// can modify values/attributes in the struct type

const usixteenbitmax float64 = 65535
const kmhMultiple float64 = 1.60934

	type car struct {
		gasPedal uint16 // anything from 0-65k (65535)
		brakePedal uint16
		steeringWheel int16 // -32k to +32k
		topSpeedKmh float64
	}

	func (c *car) kmh() float64 {
		// saying type car associates it with the car struct
		// value receiver

		// c.topSpeedKmh = 500 // <---- see how it changes the end result
		return float64(c.gasPedal) * (c.topSpeedKmh/usixteenbitmax)
	}

	func (c car) mph() float64 {
		// value receiver
		// c.topSpeedKmh = 500 < ---- see how it changes the end result
		return float64(c.gasPedal) * (c.topSpeedKmh/usixteenbitmax/kmhMultiple)
	}

	func (c *car) newTopSpeed(newspeed float64) {
		// pointer receiver
		c.topSpeedKmh = newspeed
	}

	func newerTopSpeed (c car, speed float64) car {
		// another way to write a method without using value/pointer
		c.topSpeedKmh = speed
		return c
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
		fmt.Println(aCar.mph())
		aCar.newTopSpeed(500)
		fmt.Println("new speed kmh", aCar.kmh())
		fmt.Println("new speed mph", aCar.mph())
		// the newer top speed method
		aCar = newerTopSpeed(aCar, 500)
		fmt.Println("new speed diff method kmh", aCar.kmh())
		fmt.Println("new speed diff method mph", aCar.mph())

	}
