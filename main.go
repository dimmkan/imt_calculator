package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower float64 = 2; 
	userHeight := 1.76; 
	var userWeightKg float64 = 105; 
	IMT := userWeightKg / math.Pow(userHeight, IMTPower);
	fmt.Print(IMT); 
}