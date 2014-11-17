package main
import (
	"fmt"
)

type Rock struct {
	x,y,z int
	radius float64
}



func main() {
	var t [5]*Rock;

	for i:=0;i<5;i++ {
	t[i] = new(Rock)
	t[i].radius=float64(i*2)
	t[i].x,t[i].y,t[i].z=i,i,i

	}

	fmt.Println(t)
	for i:=0;i<5;i++ {
		fmt.Println(*t[i])
	}
}