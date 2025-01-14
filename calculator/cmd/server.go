package cmd

import (
	"fmt"
	"calculator/pkg"

)

// RunServer simulates a calculator server
func RunServer() {
    fmt.Println("Calculator server running...")
    a, b := 10.0, 5.0
    fmt.Printf("Adding %f and %f: %f\n", a, b, pkg.Add(a, b))
}