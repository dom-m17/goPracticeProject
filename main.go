package main

import (
	"fmt"
	"gopracticeproject/task/display"
)

func main() {
	fmt.Println("Welcome to your task manager!")
	display.DisplayPendingTasks()
}
