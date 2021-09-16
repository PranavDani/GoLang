package main

import "fmt"

const LoginToken string = "akshdaushgd" //Public

func main() {
	var username string = "Pranav"
	fmt.Println(username)
	fmt.Printf("Variable is of the type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of the type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of the type: %T \n", smallVal)
	
	var smallFloat float64 = 255.455544546354654
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of the type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of the type: %T \n", anotherVariable)

	// implicit type

	var website = "google.com"
	fmt.Println(website)

	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of the type: %T \n", LoginToken)
}
