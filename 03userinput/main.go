package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to usser input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the reting for our Pizza: ")

	// comma ok || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)
}