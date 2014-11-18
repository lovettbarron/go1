package main

import (
	"fmt"
	"time"
)

func main() {
	if err:=run();err!=nil {
		fmt.Println(err)
	}
}

type Error struct {
	When time.Time
	What string
}

func (e *Error) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &Error {
		time.Now(),
		"ERROR MUTHAFUCKA",
	}

}