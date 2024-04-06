package main

import (
	"fmt"
	"gopracticeproject/task/display"
	"gopracticeproject/task/input"
	"os"
)

func main() {

	fmt.Println("\nWelcome to your task manager!")

	for {

		userInput := input.MainMenuInput()

		switch userInput {
		case "q":
			fmt.Println("Exiting the task manager...")
			os.Exit(0)
		case "menu":
			display.DisplayOptions()
		case "1":
			display.DisplayAllTasks()
		case "2":
			display.DisplayPendingTasks()
		case "3":
			display.DisplayCompletedTasks()
		case "4":
			display.DisplayDeletedTasks()
		case "5":
			fmt.Println("This feature is not yet completed")
		case "6":
			fmt.Println("This feature is not yet completed")
		case "7":
			fmt.Println("This feature is not yet completed")
		default:
			fmt.Println("Invalid input. Please try again.")
		}
	}
}
