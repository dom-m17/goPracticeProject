package display

import (
	"fmt"
	"gopracticeproject/resources"
)

func DisplayCompletedTasks() {
	exampleTasks := resources.ExampleTasks
	fmt.Println("Here are your completed tasks:")
	completedTaskFound := false
	for _, task := range exampleTasks {
		if task.Status == "Completed" {
			completedTaskFound = true
			fmt.Printf("ID: %d\n", task.ID)
			fmt.Printf("Title: %s\n", task.Title)
			fmt.Printf("Description: %s\n", task.Description)
			fmt.Printf("Deadline: %s\n", task.Deadline)
			fmt.Printf("Status: %s\n", task.Status)
			fmt.Println()
		}
	}

	if !completedTaskFound {
		fmt.Println("No completed tasks found")
	}
}
