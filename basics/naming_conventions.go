package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase
	// Structs, interfaces, enums

	// snake_case
	// variables, constants, filenames

	// UPPERCASE
	// Constants

	// mixedCase
	// variables

	const MAXRETRIES = 5
	var employeeID = 1001
	fmt.Println("EmployeeId:", employeeID)
}
