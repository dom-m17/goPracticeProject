package display

import (
	"fmt"
	"gopracticeproject/resources"
)

func DisplayAllTasks() {
	exampleTasks := resources.ExampleTasks
	fmt.Println("Here are all of your tasks:")
	for _, task := range exampleTasks {
		fmt.Printf("ID: %d\n", task.ID)
		fmt.Printf("Title: %s\n", task.Title)
		fmt.Printf("Description: %s\n", task.Description)
		fmt.Printf("Deadline: %s\n", task.Deadline)
		fmt.Printf("Status: %s\n", task.Status)
		fmt.Println()
	}
}
