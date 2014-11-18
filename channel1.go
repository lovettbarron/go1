package main
import "fmt"

func main() {
	ch := make(chan int,10)
	ch <- 11
	ch <- 12
	ch <- 13
	close(ch)
	// v := ch

	for i := range ch {
		defer fmt.Println("Sweet done",i) 
		fmt.Println(i)
	}
	
}

