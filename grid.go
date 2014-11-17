package main
import "fmt"

func main() {
	x:=16
	y:=8
	for i:=0;i<(x*y);i++; {
		row(i,x,y)
	}
}

func row(i,x,y int) {
	if(i%x!=0) fmt.Print(i)
	else fmt.Println(i);
}