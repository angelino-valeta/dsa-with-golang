package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Process and Manipulate array")

	arr := []int{3, 5, 4, 7, 8, 10}

	// Access elements
	fmt.Printf("Accessing third element: %v \n", arr[2])

	// Modify value
	arr[2] = 20
	fmt.Println("Modify the value of the third element: ", arr[2])

	// Array size
	fmt.Println("Array size: ", len(arr))

	// Iterate the array
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Index: %v - Value: %v\n", i, arr[i])
	}

	// Search element 20
	elementSearch := 20
	for i := 0; i < len(arr); i++ {
		if elementSearch == arr[i] {
			fmt.Printf("Found the value %v of the index %v\n", elementSearch, i)
			break
		}
	}

	// Array sort ASC
	sort.Ints(arr)
	fmt.Println("Array sort ASC")
	for _, v := range arr {
		fmt.Printf("%d \n", v)
	}

	// Insert and Remove elements in the array
	arr = append(arr, 34)
	fmt.Println("Insert elemet 34 in the array with index ", len(arr)-1)
	fmt.Println("Remove last element")
	arr = arr[:len(arr)-1]
	for _, v := range arr {
		fmt.Println(v)
	}

}
