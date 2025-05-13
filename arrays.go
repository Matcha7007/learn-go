package main

import "fmt"

func main() {
	// var arrayName [size]elementType

	var numbers [5]int
	fmt.Println(numbers)
	
	numbers[4] = 25
	fmt.Println(numbers)

	numbers[0] = 9
	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Banana", "Orang", "Grape"}
	fmt.Println("Fruits array:", fruits)
	fmt.Println("Third element:", fruits[2])

	// when updated the copied array it did not affect to the original array
	oroginalArray := [3]int{1,2,3}
	copiedArray := oroginalArray
	copiedArray[0] = 100
	fmt.Println("Original array:", oroginalArray)
	fmt.Println("Copied array:", copiedArray)

	// iterate by length array
	for i:=0; i<len(numbers); i++ {
		fmt.Println("Element at index,", i, ":", numbers[i])
	}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// use _ to store value when we cannot use it
	// _ underscore is blank identifier, used to store unused values
	for _, v:= range numbers {
		fmt.Printf("Value: %d\n", v)
	}

	a, _ := someFunc()
	fmt.Println(a)

	b := 2
	_ = b

	fmt.Println("The length of numbers array is", len(numbers))

	// comparing arrays
	array1 := [3]int{1,2,3}
	array2 := [3]int{1,2,3}

	fmt.Println("Array1 is equal to Array2:", array1 == array2)

	// multi dimensional array
	var matrix [3][3]int = [3][3]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	fmt.Println(matrix)
	
	// pointers * and addressing &
	newOriginalArray := [3]int{1,2,3}
	var newCopiedArray *[3]int
	newCopiedArray = &newOriginalArray // addresses to newOriginalArray
	newCopiedArray[0] = 100 // so when we updated value at newCopiedArray, will be stored to the newOriginalArraygo 

	fmt.Println("Original Array:", newOriginalArray)
	fmt.Println("Copied Array:", *newCopiedArray)
}

func someFunc() (int, int){
	return 1, 2
}