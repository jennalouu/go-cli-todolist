package cmd

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"go-cli-todolist/constants"

	"github.com/spf13/pflag"
)

func TestAddTask_IDPaddingAndPersistence(t *testing.T) {
	os.Remove(constants.DEFAULT_FILE_NAME)
	t.Cleanup(func() {
		os.Remove(constants.DEFAULT_FILE_NAME)
	})

	success1 := addTask("Buy groceries", constants.TASK_NOT_STARTED)
	if !success1 {
		t.Fatalf("expected addTask to return true for first task")
	}

	success2 := addTask("Write Go tests", constants.TASK_IN_PROGRESS)
	if !success2 {
		t.Fatalf("expected addTask to return true for second task")
	}

	rawData, err := os.ReadFile(constants.DEFAULT_FILE_NAME)
	if err != nil {
		t.Fatalf("failed to read created test file: %v", err)
	}

	var tasks []constants.Task
	if err := json.Unmarshal(rawData, &tasks); err != nil {
		t.Fatalf("failed to parse JSON from test file: %v", err)
	}

	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks in file, got %d", len(tasks))
	}

	if tasks[0].ID != "01" {
		t.Errorf("expected first task ID to be %q, got %q", "01", tasks[0].ID)
	}
	if tasks[1].ID != "02" {
		t.Errorf("expected second task ID to be %q, got %q", "02", tasks[1].ID)
	}
}

func TestAddCmd_CLIExecution(t *testing.T) {
	os.Remove(constants.DEFAULT_FILE_NAME)
	t.Cleanup(func() {
		os.Remove(constants.DEFAULT_FILE_NAME)
	})

	tests := []struct {
		name           string
		args           []string
		expectedOutput string
	}{
		{
			name:           "Missing task flag",
			args:           []string{"add"},
			expectedOutput: "How about you give me a task to add",
		},
		{
			name:           "Empty task string",
			args:           []string{"add", "--task", ""},
			expectedOutput: "You can't add an empty task, Idiot <3",
		},
		{
			name:           "Valid task addition",
			args:           []string{"add", "--task", "Learn Go Testing"},
			expectedOutput: "Yeah! Added task 'Learn Go Testing'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				addCmd.Flags().VisitAll(func(f *pflag.Flag) {
					f.Changed = false
					f.Value.Set(f.DefValue)
				})

				rootCmd.SetArgs(tt.args)
				rootCmd.Execute()
			})

			if !strings.Contains(output, tt.expectedOutput) {
				t.Errorf("expected output to contain %q, but got:\n%s", tt.expectedOutput, output)
			}
		})
	}

}
