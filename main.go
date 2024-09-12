package main

import (
	"fmt"
	"gostart/car"
)

func main() {
	ferrari := car.Car{
		Brand: "Ferrari",
		Plate: "488 GTB",
		Color: "red",
		Year:  2021,
	}
	ferrari.Run()
	fmt.Printf(ferrari.Brand)
}
