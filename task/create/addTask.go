package create

import (
	"fmt"
	"gopracticeproject/resources"
	"time"
)

func AddTask() {
	var taskTitle string
	var taskDescription string
	var taskDeadlineDate time.Time
	fmt.Println("Enter the title for your new task.")
	fmt.Scanln(&taskTitle)
	fmt.Println("Enter a description for your new task.")
	fmt.Scanln(&taskDescription)
	fmt.Println("Enter the deadline for your new task.")
	fmt.Println("(Must be in the form DD-MM-YYYY.")
	taskDeadlineDate = getDate()
	// Append new task to list of pending tasks
	// Display task

	fmt.Println(taskTitle, taskDescription, taskDeadlineDate)

	exampleTasks := resources.ExampleTasks
	exampleTasks = append(exampleTasks, resources.Task{
		ID:          10, // Placeholder value, needs to be incremented from the highest value in the existing list.
		Title:       taskTitle,
		Description: taskDescription,
		Deadline:    taskDeadlineDate,
		Status:      "Pending",
	}) // This is not updating the slice in exampleTasks.go

	fmt.Println(exampleTasks)
}

func getDate() time.Time {
	var deadlineInput string
	fmt.Scanln(&deadlineInput)

	deadlineDate, err := time.Parse("02-01-2006", deadlineInput)
	if err != nil {
		fmt.Println("Error:", err)
		return time.Time{}
	}
	deadlineDate = deadlineDate.Local()
	deadlineDate = deadlineDate.Add(time.Hour*23 + time.Minute*59 + time.Second*59)
	return deadlineDate
}
