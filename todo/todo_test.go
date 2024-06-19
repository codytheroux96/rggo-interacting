package todo_test

import (
	//"os"
	"testing"

	"github.com/codytheroux96/rggo-interacting"
)

func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != "New Task" {
		t.Errorf("expected %q; got %q instead", taskName, l[0].Task)
	}
}