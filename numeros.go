package main

import (
	"fmt"
	"math"
	"strconv"
)

func multiplosNumeros(numero, expoente, multiplo, repentindo float64) {
	x := exponent(numero, expoente)
	y := int(x)

	var c int
	for i := 0; i <= y; i++ {
		if i%int(multiplo) != 0 {
		} else {
			if c%int(repentindo) == 0 {
				fmt.Print("\n")
			}
			c++
			j := strconv.Itoa(i)
			fmt.Print(j + " ")
		}
	}
}

func exponent(x, y float64) float64 {
	return math.Pow(x, y)
}
