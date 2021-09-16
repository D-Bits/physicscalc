package main

import (
	"fmt"
	"math"
)

func main() {

	value := gravForce(5.98*(math.Pow(10, 24)), 7.34*(math.Pow(10, 22)), 385000000)

	fmt.Println(value)
}
