package main

import (
	"fmt"
	"slices"
)

func main() {

	// var sliceName []ElementType
	// var numbers []int
	// var numbers1 = []int{1, 2, 3}
	// numbers2 := []int{9, 8, 7}
	// slice := make([]int, 5)

	// slice from an array
	a := [5]int{1, 2, 3, 4, 5}
	slice1 := a[1:4]
	fmt.Println(slice1)
	
	// append slice
	slice1 = append(slice1, 5,6)
	fmt.Println("slice1:", slice1)
	
	// copy original slice
	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println("sliceCopy:", sliceCopy)
	
	for i, v := range slice1 {
		fmt.Println(i,v)
	}
	fmt.Println("Element at index 3 of slice1:", slice1[3])
	
	slice1[3] = 50
	fmt.Println("slice1:", slice1)
	fmt.Println("Element at index 3 of slice1:", slice1[3])

	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("slice1 is equal to sliceCopy")
	}

	// dynamic 2 dimensions
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		fmt.Println("Outer Slice:", twoD)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
			fmt.Printf("Adding value %d in outer slice at index %d, and in inner slice index of %d\n", i+j, i, j)
			fmt.Println("Inner Slice after adding value:", twoD)
		}
	}
	fmt.Println(twoD)

	// slice[low:high]
	slice2 := slice1[2:4]
	fmt.Println(slice2)
	fmt.Println("The capacity of slice2 is", cap(slice2))
	fmt.Println("The length of slice2 is", len(slice2))
}