package main

import "fmt"

func main() {
	test(1);
	test(5);
}

func test(n int) {
	if n:=n*2;n>7 {
		fmt.Println("More than 7")
	} else {
		fmt.Println("Less than 7")
	}
}
