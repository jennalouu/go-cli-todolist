package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"charm.land/lipgloss/v2"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

const (
	TASK_DONE_PRETTY        = "● done"
	TASK_IN_PROGRESS_PRETTY = "○ in Progress"
	TASK_NOT_STARTED_PRETTY = "○ not Started"

	TASK_DONE        = "done"
	TASK_IN_PROGRESS = "in Progress"
	TASK_NOT_STARTED = "not Started"

	DEFAULT_FILE_NAME = "todo.json"

	ERROR_FILE_NOT_FOUND_OR_EMPTY = "You don't have any ToDo's yet, you can add some if you want to with the 'add' command :)"
)

type Task struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "shows all of the tasks in the todo list",
	Run: func(cmd *cobra.Command, args []string) {
		todos := getTodos()

		fmt.Println("Your To-Do's for today")
		showPrettyTodoList(todos)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}

func showPrettyTodoList(todos []Task) {
	pinkHeader := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFB6C1")).
		Bold(true)

	purpleId := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#E6E6FA"))

	mintActive := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#98FB98")).
		Bold(true)

	rosePause := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFC0CB"))

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{
		pinkHeader.Render("#ID"),
		pinkHeader.Render("Name"),
		pinkHeader.Render("Status"),
	})

	for _, todo := range todos {
		var statusStyle lipgloss.Style
		var prettyStatus string

		switch todo.Status {
		case TASK_DONE:
			statusStyle = mintActive
			prettyStatus = TASK_DONE_PRETTY
		case TASK_IN_PROGRESS:
			statusStyle = mintActive
			prettyStatus = TASK_IN_PROGRESS_PRETTY
		case TASK_NOT_STARTED:
			statusStyle = rosePause
			prettyStatus = TASK_NOT_STARTED_PRETTY
		}

		t.AppendRows([]table.Row{
			{purpleId.Render(todo.ID), todo.Name, statusStyle.Render(prettyStatus)},
		})
	}

	cuteStyle := table.StyleRounded

	frameStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#D8BFD8"))

	cuteStyle.Box.Left = frameStyle.Render("│")
	cuteStyle.Box.Right = frameStyle.Render("│")
	cuteStyle.Box.TopLeft = frameStyle.Render("─")
	cuteStyle.Box.TopRight = frameStyle.Render("─")
	cuteStyle.Box.BottomLeft = frameStyle.Render("─")
	cuteStyle.Box.BottomRight = frameStyle.Render("─")
	cuteStyle.Box.MiddleHorizontal = frameStyle.Render("─")
	cuteStyle.Box.MiddleVertical = frameStyle.Render("│")
	cuteStyle.Box.TopLeft = frameStyle.Render("╭")
	cuteStyle.Box.TopRight = frameStyle.Render("╮")
	cuteStyle.Box.BottomLeft = frameStyle.Render("╰")
	cuteStyle.Box.BottomRight = frameStyle.Render("╯")
	cuteStyle.Box.TopSeparator = frameStyle.Render("┬")
	cuteStyle.Box.BottomSeparator = frameStyle.Render("┴")
	cuteStyle.Box.LeftSeparator = frameStyle.Render("├")
	cuteStyle.Box.RightSeparator = frameStyle.Render("┤")
	cuteStyle.Box.PageSeparator = frameStyle.Render("┼")

	t.SetStyle(cuteStyle)

	t.Render()
}

func getTodos() []Task {
	_, err := os.Stat(DEFAULT_FILE_NAME)
	if err != nil {
		fmt.Println(ERROR_FILE_NOT_FOUND_OR_EMPTY)
		os.Exit(0)
	}

	rawData, err := os.ReadFile(DEFAULT_FILE_NAME)
	if err != nil {
		fmt.Println("Error reading the todo file:", err)
		os.Exit(0)
	} else if len(rawData) == 0 || string(rawData) == "[]" {
		fmt.Println(ERROR_FILE_NOT_FOUND_OR_EMPTY)
		os.Exit(0)
	}

	var tasks []Task
	err = json.Unmarshal(rawData, &tasks)
	if err != nil {
		fmt.Println("Error parsing the todo file:", err)
		os.Exit(0)
	}

	return tasks
}
