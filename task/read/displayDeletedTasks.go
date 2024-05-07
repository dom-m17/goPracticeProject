package read

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DisplayDeletedTasks() {
	fmt.Println()
	fmt.Println("Here are your deleted tasks:")
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	rows, err := dbpool.Query(context.Background(), "select * from tasks where status = 'deleted';")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	for rows.Next() {
		var id int
		var title string
		var description sql.NullString
		var deadline time.Time
		var status string

		err := rows.Scan(&id, &title, &description, &deadline, &status)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Scan failed: %v\n", err)
			os.Exit(1)
		}

		var descString string
		if description.Valid {
			descString = description.String
		} else {
			descString = "NULL"
		}

		fmt.Printf("Title: %s\n", title)
		fmt.Printf("Description: %s\n", descString)
		fmt.Printf("Deadline: %s\n", deadline.Format("02-01-2006"))
		fmt.Println()
	}

	fmt.Println("End of deleted tasks")
}
