package resources

import "time"

type Task struct {
	ID          int
	Title       string
	Description string
	Deadline    time.Time
	Status      string
}

var ExampleTasks = []Task{
	{
		ID:          1,
		Title:       "Complete Project",
		Description: "Finish the task manager project in Go",
		Deadline:    time.Date(2024, time.May, 10, 23, 59, 59, 0, time.UTC),
		Status:      "Pending",
	},
	{
		ID:          2,
		Title:       "Create Project",
		Description: "Create a Go project",
		Deadline:    time.Date(2024, time.April, 5, 23, 59, 59, 0, time.UTC),
		Status:      "Completed",
	},
	{
		ID:          3,
		Title:       "Go for a run",
		Description: "Go for a 7km run",
		Deadline:    time.Date(2024, time.April, 8, 23, 59, 59, 0, time.UTC),
		Status:      "Pending",
	},
	{
		ID:          4,
		Title:       "Go to gym",
		Description: "Leg day at the gym",
		Deadline:    time.Date(2024, time.April, 6, 23, 59, 59, 0, time.UTC),
		Status:      "Completed",
	},
}
