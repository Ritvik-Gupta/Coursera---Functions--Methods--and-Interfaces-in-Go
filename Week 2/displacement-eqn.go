package main

import "fmt"

func main() {
	var acceleration, initialVelocity, initialDisplacement float64

	fmt.Print("Enter the Acceleration :\t")
	fmt.Scanln(&acceleration)

	fmt.Print("Enter the Initial Velocity :\t")
	fmt.Scanln(&initialVelocity)

	fmt.Print("Enter the Initial Displacement :\t")
	fmt.Scanln(&initialDisplacement)

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	var time float64
	fmt.Print("Enter the Time at which to compute the Displacement :\t")
	fmt.Scanln(&time)

	fmt.Println("Displacement at Time (t) =\t", fn(time))
}

func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {
	return func(atTime float64) float64 {
		return (1/2)*acceleration*atTime*atTime + initialVelocity*atTime + initialDisplacement
	}
}
