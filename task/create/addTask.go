package create

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
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
	fmt.Println("(Must be in the form YYYY-MM-DD.")
	taskDeadlineDate = getDate()

	fmt.Println(taskTitle, taskDescription, taskDeadlineDate)

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	_, err = dbpool.Exec(context.Background(), "INSERT INTO tasks (title, description, deadline) VALUES ($1, $2, $3)", taskTitle, taskDescription, taskDeadlineDate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert task: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Task added successfully!")
}

func getDate() time.Time {
	var deadlineInput string
	fmt.Scanln(&deadlineInput)

	deadlineDate, err := time.Parse("2006-01-02", deadlineInput)
	if err != nil {
		fmt.Println("Error:", err)
		return time.Time{}
	}
	deadlineDate = deadlineDate.Local()
	deadlineDate = deadlineDate.Add(time.Hour*23 + time.Minute*59 + time.Second*59)
	return deadlineDate
}
