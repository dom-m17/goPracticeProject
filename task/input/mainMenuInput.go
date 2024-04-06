package input

import "fmt"

func MainMenuInput() string {
	var userSelection string

	// fmt.Println("You are at the main menu.")
	fmt.Println("\nEnter your selection: ")
	fmt.Println("(Type 'menu' to view options)")
	// fmt.Println("Type 1 to view all tasks")
	// fmt.Println("Type 2 to view pending tasks")
	// fmt.Println("Type 3 to view completed tasks")
	// fmt.Println("Type 4 to view deleted tasks")
	// fmt.Println("Type 5 to add new task")
	// fmt.Println("Type 6 to amend task")
	// fmt.Println("Type 7 to complete task")

	fmt.Scanln(&userSelection)
	return userSelection
}
