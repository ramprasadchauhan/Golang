package main

import "fmt"

const LoginToken string = "rtargtfsgvsnb" // public

func main() {
	// var username string = "ramprasad"
	// fmt.Println(username)
	// fmt.Printf("Variable is type of : %T \n ", username)

	// var isLoggedIn bool = true
	// fmt.Println(isLoggedIn)
	// fmt.Printf("Variable is type of : %T \n ", isLoggedIn)

	// var smallVal uint8 = 255
	// fmt.Println(smallVal)
	// fmt.Printf("Variable is type of : %T \n ", smallVal)

	var smallFloat float32 = 255.4545487123456789
	fmt.Println(smallFloat)
	fmt.Printf("Variable is type of : %T \n ", smallFloat)

	// default values and aliases
	 var anotherVariable  int
	 fmt.Println(anotherVariable)
	fmt.Printf("Variable is type of : %T \n ", anotherVariable)

	// implicit type

	var website = "google.com"
	fmt.Println(website)
	
	// no var style

	numberOfUser := 20000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is type of : %T \n ", LoginToken)

}

// := walarus operator  is a shorthand for declaring a variable and assigning a value to it at the same time, it is use d insite method