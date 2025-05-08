package main

import (
	"fmt"
	"math"
)

func main() {
	// variable declaration
	var a, b int = 10, 3
	var result int

	// addition
	result = a + b
	fmt.Println("Addition: ", result)

	// substraction
	result = a - b
	fmt.Println("Subtraction: ", result)

	// multiplication
	result = a * b
	fmt.Println("Multiplication: ", result)

	// devision
	result = a / b
	fmt.Println("Devision: ", result)
	// the result devided a by b = 3 is not accurate. It because a and b are integer.
	// if we want to accurate result, we should make operand a or b is floating
	// example
	const p float64 = 10 / 3.0
	fmt.Println(p)

	// remainder
	result = a % b
	fmt.Println("Remainder: ", result)

	// overflow with signed integers
	var maxInt int64 = 9223372036854775807 // max value that int64 can hold
	fmt.Println(maxInt)
	
	maxInt = maxInt + 1
	fmt.Println(maxInt)
	// when maxInt add 1, it can go to negative value -9223372036854775808 as minimum value

	// overflow with unsigned integers
	var uMaxInt uint64 = 18446744073709551615 // max value that uint64 can hold
	fmt.Println(uMaxInt)
	
	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)
	// when uMaxInt add 1, it cannot go to negative value but go to 0 as minimum value

	// underflow with floating point number
	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)
	
	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)

}