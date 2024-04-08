package read

import (
	"fmt"
	"gopracticeproject/resources"
)

func DisplayPendingTasks() {
	exampleTasks := resources.ExampleTasks
	fmt.Println("Here are your pending tasks:")
	pendingTaskFound := false
	for _, task := range exampleTasks {
		if task.Status == "Pending" {
			pendingTaskFound = true
			fmt.Printf("ID: %d\n", task.ID)
			fmt.Printf("Title: %s\n", task.Title)
			fmt.Printf("Description: %s\n", task.Description)
			fmt.Printf("Deadline: %s\n", task.Deadline)
			fmt.Printf("Status: %s\n", task.Status)
			fmt.Println()
		}
	}

	if !pendingTaskFound {
		fmt.Println("No pending tasks found")
	}
}
