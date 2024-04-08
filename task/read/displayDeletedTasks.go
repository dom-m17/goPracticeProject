package read

import (
	"fmt"
	"gopracticeproject/resources"
)

func DisplayDeletedTasks() {
	exampleTasks := resources.ExampleTasks
	fmt.Println("Here are your deleted tasks:")
	foundDeletedTasks := false
	for _, task := range exampleTasks {
		if task.Status == "Deleted" {
			foundDeletedTasks = true
			fmt.Printf("ID: %d\n", task.ID)
			fmt.Printf("Title: %s\n", task.Title)
			fmt.Printf("Description: %s\n", task.Description)
			fmt.Printf("Deadline: %s\n", task.Deadline)
			fmt.Printf("Status: %s\n", task.Status)
			fmt.Println()
		}
	}

	if !foundDeletedTasks {
		fmt.Println("No deleted tasks found.")
	}
}
