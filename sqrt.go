package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
	z := x
	for i:=0;i<10;i++ {
		z=z-( ((z*z)-x) / 2*z )
		fmt.Println("i:",i,"z:",z)
	}
	return z
}

func main() {
	var n float64= 2
    fmt.Println("Mine: ",Sqrt(n))
    fmt.Println("Math: ",math.Sqrt(n))
}