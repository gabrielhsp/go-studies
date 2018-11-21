package main

import "fmt"

func main() {
	/** This is a simple hello world using Go */
	fmt.Println("Hello World")

	/** This is the explicit way to declare a variable */
	var test int8;
	
	/** We can also set a value when we declare */
	var num int8 = 1;

	/**
	 * Using the explicit way, we can declare multiple variables
	 * Note that all variables will use the same type declaration
	 */
	var num1, num2 int8;

	/** We can also declare multiple variables and set values to then separating each value using commas */
	var num2, num3 int8 = 100, 9;

	/** Also, we can use another declaration type, using the nested way */
	var (
		testing int8 = 1
		testing2 String = "Test"
		testing boolean = false
	)

	/** 
	 * This is the implicity way to declare a variable using Go
	 * It's important to notice that the implicity way to declare variables can only be used inside a scope function
	 */
	num4 := "Testing"

	/** We can also declare multiple variables using the implicity way */
	num5, string1 := 9, "One string"

	/** Also, we can use the nested implicity declaration type */
	var (
		implicityInt = 1
		implicityString = "A simple string"
		implicityBoolean = true
	)

	/** 
	 * This is the explicity explicity way to declare constants
	 * The difference between constants and variables are that we can't set a new value to a constant
	 * That's why we need to take care using constants
	 */
	const (
		constantInt int8 = 10
		constantString String = "One more test"
		constantBoolean boolean = true
	)

	/** This is the implicity way to declare constants */
	const (
		constImpInt = 10
		constImpString = "Testing bro"
		constImpBoolean = false
	)
}