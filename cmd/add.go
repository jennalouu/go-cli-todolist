package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"go-cli-todolist/constants"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task to the todo list",
	Run: func(cmd *cobra.Command, args []string) {
		status := constants.TASK_NOT_STARTED
		if cmd.Flags().Changed("task") == false {
			fmt.Println("How about you give me a task to add, you know, like a name or something?")
			return
		}
		task, err := cmd.Flags().GetString("task")
		if err != nil {
			fmt.Println("Could not get the task through the flag, sorry :(")
		} else if task == "" {
			fmt.Println("You can't add an empty task, Idiot <3")
			return
		}

		inProgress, _ := cmd.Flags().GetBool("in-progress")
		if inProgress {
			status = constants.TASK_IN_PROGRESS
		}

		if addTask(task, status) {
			fmt.Printf("Yeah! Added task '%s' with status '%s' successfully!\n", task, status)
		} else {
			fmt.Println("Could not add the task, sorry :(")
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("task", "t", "", "Name of the task to add")
	addCmd.Flags().BoolP("in-progress", "p", false, "sets the status to in Progress")
}

func addTask(taskName string, status string) bool {
	if _, err := os.Stat(constants.DEFAULT_FILE_NAME); errors.Is(err, os.ErrNotExist) {
		os.WriteFile(constants.DEFAULT_FILE_NAME, []byte("[]"), 0644)
	}

	rawData, err := os.ReadFile(constants.DEFAULT_FILE_NAME)
	if err != nil {
		fmt.Println("Error reading the todo file:", err)
		os.Exit(0)
	}

	var tasks []constants.Task
	err = json.Unmarshal(rawData, &tasks)
	if err != nil {
		fmt.Println("Error parsing the todo file:", err)
		os.Exit(0)
	}

	var prettyId string
	id := len(tasks) + 1
	if id < 10 {
		prettyId = fmt.Sprintf("0%d", id)
	} else {
		prettyId = fmt.Sprintf("%d", id)
	}

	newTask := constants.Task{
		ID:     prettyId,
		Name:   taskName,
		Status: status,
	}

	tasks = append(tasks, newTask)
	jsonData, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Error converting tasks to JSON:", err)
		os.Exit(0)
	}

	err = os.WriteFile(constants.DEFAULT_FILE_NAME, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to the todo file:", err)
		os.Exit(0)
	}

	return true
}
