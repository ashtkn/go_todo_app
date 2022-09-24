package service

import (
	"context"
	"fmt"

	"github.com/ashtkn/go_todo_app/entity"
	"github.com/ashtkn/go_todo_app/store"
)

type ListTask struct {
	DB   store.Queryer
	Repo TaskLister
}

func (l *ListTask) ListTask(ctx context.Context) (entity.Tasks, error) {
	tasks, err := l.Repo.ListTasks(ctx, l.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}
	return tasks, nil
}
