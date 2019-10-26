package main

import "fmt"

func main() {
	a := 1
	// for statement
	for a <= 10 {
		fmt.Println(a)
		a++
	}

	// if statement with looping for
	for b := 1; b <= 10; b++ {
		if b%2 == 0 {
			fmt.Println(b, "even")
		} else {
			fmt.Println(b, "odd")
		}
	}

	// case statement
	c := 1
	for c <= 10 {
		switch c {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Undefined Number")
		}
		c++
	}
}
