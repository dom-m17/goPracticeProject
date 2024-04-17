// package main

// import (
// 	"fmt"
// 	"gopracticeproject/task/create"
// 	"gopracticeproject/task/input"
// 	"gopracticeproject/task/read"
// 	"os"
// )

// func main() {

// 	fmt.Println("\nWelcome to your task manager!")

// 	for {

// 		userInput := input.MainMenuInput()

// 		switch userInput {
// 		case "q":
// 			fmt.Println("Exiting the task manager...")
// 			os.Exit(0)
// 		case "menu":
// 			read.DisplayOptions()
// 		case "1":
// 			read.DisplayAllTasks()
// 		case "2":
// 			read.DisplayPendingTasks()
// 		case "3":
// 			read.DisplayCompletedTasks()
// 		case "4":
// 			read.DisplayDeletedTasks()
// 		case "5":
// 			create.AddTask()
// 		case "6":
// 			fmt.Println("This feature is not yet completed")
// 		case "7":
// 			fmt.Println("This feature is not yet completed")
// 		case "8":
// 			fmt.Println("This feature is not yet completed")
// 		default:
// 			fmt.Println("Invalid input. Please try again.")
// 		}
// 	}
// }

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var first_name string
	var last_name string
	err = dbpool.QueryRow(context.Background(), "select * from gopracticeproject where first_name='Dominic';").Scan(&first_name, &last_name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(first_name, last_name)
}
