package main

import (
	"fmt"
	"time"

	"github.com/ccie57654/hellogopher/stringutils"

	"github.com/jboursiquot/go-proverbs"
)

const location = "Remote"

var name string = "lawrence lucas"

func main() {

	//name = 'L'
	from := `Idaho`
	var n time.Time = time.Now()

	var proverb = "Undefined"
	if p, err := proverbs.Nth(-4); err == nil {
		proverb = p.Saying
	} else {
		proverb = err.Error()
	}
	//n = time.Now()

	fmt.Printf("Hello, fellow %s Gophers!\n", location)
	fmt.Printf("My name is %s and I'm from %s.\n", stringutils.Upper(name), from)
	fmt.Printf("By the time %v o'clock EST comes around, we'll know how to code in Go!\n", n)
	fmt.Printf("Here's a Go proverb: %s\n", proverb)
	fmt.Println("Let's get started!")
}
