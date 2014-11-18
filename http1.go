package main
import (
	"fmt"
	"net/http"
)

type Obj struct {}

func (o Obj) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"Returning the thing!")
}

func main() {
	var o Obj
	http.ListenAndServe("localhost:4000",o);
}
