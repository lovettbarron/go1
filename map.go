package main

import "fmt"

type Pin struct {
	lat, lon float64
}

var m map[string]Pin

var t = map[string]Pin{
	"TheOne" : { 123.2, 333.2 },
	"TheTwo" : { 232.2, 444.2 }, // Why do I not need a comma here?
}

func main() {
	m = make(map[string]Pin)
	m["place"] = Pin{
		1.01, 2.02, // Why do I need a comma here?
	}
	m["anotherPlace"] = Pin{ 4.01, -122.02 } // Comma as indicator of "Consider next line?"
	fmt.Println(m["place"])
	fmt.Println(m)
}