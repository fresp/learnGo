package main

import "fmt"

func main() {
	/*
		//	Shortvariable
		name := "Max"
		fmt.Println("My Dogs name is", name)*/

	/*
		//	Constant variable
		const x string = "Hello World"
		x = "Some other string" // Not working change value, because const cannot be change later
		fmt.Println(x) */

	//	Multiple Variable with matematika function
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}
