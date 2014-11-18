package main

import "fmt"

type Pin struct {
	lat, lon float64
}

var m map[string]Pin

func main() {
	m = make(map[string]Pin)
	m["place"] = Pin{
		1.01, 2.02,
	}
	m["anotherPlace"] = Pin{
		4.01, -122.02,
	}
	fmt.Println(m["place"])
	fmt.Println(m)
}