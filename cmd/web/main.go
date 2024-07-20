package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

const taskFile = "task.json"

func main() {
	for {
		fmt.Println("Task List")
		fmt.Println("1 Add Task")
		fmt.Println("2 View Task")
		fmt.Println("3 Delete Task")
		fmt.Println("4 Mark Task as Complete")
		fmt.Println("5 Exit")
		fmt.Println("Enter your choice: ")

		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')

		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			addTask()
		case "2":
			viewTask()
		case "3":
			deleteTask()
		case "4":
			markTaskAsComplete()
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid Choice. Please try again. Please enter a number between 1 and 5.")
		}
	}

}

func addTask() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter task description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)
}

