package main

import (
	"fmt"
	"calculator/cmd"
	"calculator/internal/config"
	"calculator/pkg"

	
)


func main() {
    // Load configuration
    cfg := config.LoadConfig()
    fmt.Printf("Welcome to %s!\n", cfg.AppName)

	 // Perform operations based on user input
	 var a, b float64
	 var operator string
	 fmt.Println("Enter first number:")
	 fmt.Scanln(&a)
	 fmt.Println("Enter second number:")
	 fmt.Scanln(&b)
	 fmt.Println("Enter operator (+, -, *, /):")
	 fmt.Scanln(&operator)
 
	 switch operator {
	 case "+":
		 fmt.Printf("Result: %f\n", pkg.Add(a, b))
	 case "-":
		 fmt.Printf("Result: %f\n", pkg.Subtract(a, b))
	 case "*":
		 fmt.Printf("Result: %f\n", pkg.Multiply(a, b))
	 case "/":
			result, err := pkg.Divide(a, b)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Result: %f\n", result)
			}
	 default:
		 fmt.Println("Invalid operator. Please use +, -, *, or /.")
	 }
 
	 // Demonstrate server functionality
	 cmd.RunServer()

}