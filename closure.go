package main

import "fmt"


// Closure test
func q() func(int) int { // <-- Don't get this syntax
	s := 0
	return func(x int) int {
		s += x
		fmt.Println("s is ",s)
		return s
	}
}

// Fibonacci
func fib() func(int) int {
	f:=1
	return func(x int) int {
		f += f;
		// fmt.Println("Fib is ", f)
		return f
	}
}

func main() {

	// Function as Variable
	p := func() float64 {
		return 3.41
	}

	// Closure test
	fmt.Println("Func in Var", p())
	thing := q()
	for i:=0;i<10;i++ {
		fmt.Println("Closure",i,thing(i))	
	}

	// Fibonacci
	f := fib()
	for i:=0;i<10;i++ {
		fmt.Println("Fib",i,f(i))	
	}
}