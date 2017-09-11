package main

import "fmt"

func main() {
	// var declares one or more vars
	var a string = "initial"
	fmt.Println(a)

	// Declare multiple vars
	var b, c int = 1, 2
	fmt.Println(b, c)

	// infer type
	var d = true
	fmt.Println(d)

	// default to Zero for type if not declared
	var e int
	fmt.Println(e)

	// shorthand
	f := "short"
	fmt.Println(f)
}
