package main

import "fmt"

// declare global variables we should use var, gofer symbol cannot use to  declar global variable outside scope
// global scope is limited only to the package, outside package cannot access the variables
// variables are used for storing data that can change during program execution
// and variables are mutable
var middleName = "Cane"

func main() {
	// declare local variables
	// declare variable using var
	// var age int
	// var name string = "John"
	// var name1 = "Jane"

	// declare variables using gofer symbol :=, data type include in assign value
	// count := 10
	// lastName := "Smith"

	middleName := "Mine"
	fmt.Println(middleName)
	// output middleName will be "Mine"

	// ---- Default Values
	// 1. Numeric Types: 0
	// 2. Boolean Types: false
	// 3. String Types: ""
	// 4. Pointers, Slices, Maps, Functions, and Structs: nil

	// ---- Scope
	printName()
}

func printName() {
	firstName := "Michael"
	fmt.Println(firstName)
}