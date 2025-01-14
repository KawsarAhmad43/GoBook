package main

import "fmt"

func PrintDataTypes() {
	// basic types: int, float64, string, bool
	// compex types: array, slice, map, struct, pointer, function
	

    // defining int
	var age int = 25
	fmt.Println("Age: ", age)

	// defining float64
	var salary float64 = 300.00
	fmt.Println("Salary: ", salary)


	// defining string
	var name string = "Kawsar"
	fmt.Println("Name: ", name)

	// defining bool
	var isMarried bool = true
	fmt.Println("Is Married: ", isMarried)

	// defining array
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array: ", arr)

	// Defining slice
	slice := []int{1, 2, 3, 4, 5} 
	fmt.Println("Slice:", slice)

	// defining map
	user:= map[string]string{"name": "Kawsar", "role": "software engineer"}
	fmt.Println("Map: ", user)

}