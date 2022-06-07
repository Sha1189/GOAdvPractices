package main

import (
	"errors"
	"fmt"
	"math"
)

func calcCircleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("Area calculation failed, radius is less than zero")
	}
	return math.Pi * radius * radius, nil
}
func main() {
	radius := -10.00
	area, err := calcCircleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}
