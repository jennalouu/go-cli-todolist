package constants

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
