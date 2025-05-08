package basics

import "fmt"

// 1. PascalCase
// Commonly used for naming Structs, Interfaces, Enums
// Eg. CalculateArea, UserInfo, NewHTTPRequest
type EmployeeGoogle struct {
	FirstName string
	LastName string
	Age int
}

func main() {
	// 2. snake_case
	// Commonly used for naming variables, constants, and file names
	// Eg. user_id, first_name, last_name, http_request

	// 3. UPPERCASE
	// UPPERCASE is used 	exclusively for naming constants in go
	// Constants are immutable unlike variables which are mutable
	const MAXRETRIES = 5

	// 4. mixedCase
	// mixedCase is similiar to camel case, start with a lowercase letter
	// Eg. javaScript, htmlDocument, isValid
	var employeeID = 10001
	fmt.Println("Employee ID: ", employeeID)

}