package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"go-cli-todolist/constants"
)

func captureOutput(f func()) string {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestShowPrettyTodoList(t *testing.T) {
	tasks := []constants.Task{
		{
			ID:     "01",
			Name:   "task 1",
			Status: constants.TASK_NOT_STARTED,
		},
		{
			ID:     "02",
			Name:   "task 2",
			Status: constants.TASK_DONE,
		},
	}

	output := captureOutput(func() {
		showPrettyTodoList(tasks)
	})

	expectedContents := []string{
		"#ID",
		"NAME",
		"STATUS",
		"01",
		"task 1",
		"02",
		"task 2",
	}

	for _, expected := range expectedContents {
		if !strings.Contains(output, expected) {
			t.Errorf("Achtung: Die Ausgabe sollte %q enthalten, tut sie aber nicht.\nAusgabe war:\n%s", expected, output)
		}
	}
}

func TestGetTodos(t *testing.T) {
	testContent := `[{"id":"01","name":"Test-Task","status":"done"}]`

	err := os.WriteFile(constants.DEFAULT_FILE_NAME, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Test-Datei konnte nicht erstellt werden: %v", err)
	}

	t.Cleanup(func() {
		os.Remove(constants.DEFAULT_FILE_NAME)
	})

	todos := getTodos()

	if len(todos) != 1 {
		t.Fatalf("expected: 1 Task, actual: %d Tasks", len(todos))
	}

	if todos[0].Name != "Test-Task" {
		t.Errorf("expected Task-Name: %q, actual: %q", "Test-Task", todos[0].Name)
	}
}
