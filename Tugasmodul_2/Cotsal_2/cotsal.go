package main

import (
	"fmt"
	"math"
)

func main() {
	var r int

	fmt.Print("Jejari = ")
	fmt.Scan(&r)
	pi := 3.1415926535

	v := (4.0 / 3.0) * pi * math.Pow(float64(r), 3)
	l := 4 * pi * math.Pow(float64(r), 2)

	fmt.Printf("Bola dengan jari-jari %d memiliki volume %.10f dan luas permukaan %.10f\n", r, v, l)
}
