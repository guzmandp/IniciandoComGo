package car

import (
	"fmt"
)

type Car struct {
	Brand string `json: "brand`
	Plate string `Json: "plate"`
	Color string `Json: "color"`
	Year  int    `Json: "year"`
}

func (c *Car) Run() {
	c.Year = 2021
	fmt.Println(c.Brand, "is running", c.Year)
}
