package car

import (
	"fmt"
)

type Car struct {
	Brand string
	Plate string
	Color string
	Year  int
}

func (c *Car) Run() {
	c.Year = 2021
	fmt.Println(c.Brand, "is running", c.Year)
}
