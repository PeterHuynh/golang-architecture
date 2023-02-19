package main

import "fmt"

type person struct {
	first string
}

func main() {
	p1 := person{
		first: "James",
	}

	fmt.Printf("%T\n", p1)
}
