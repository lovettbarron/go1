package main

import "fmt"

type Vertex struct {
	x int
	y int
	z int
}

func main() {
	v:=Vertex{1,2,3}
	vp:=&v
	vt:=&v.z
	
	fmt.Println(v)
	fmt.Println(v.z)
	fmt.Println(vp)
	fmt.Println(*vp)
	fmt.Println(vt)
	fmt.Println(*vt)
}