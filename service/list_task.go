package service

import (
	"context"
	"fmt"

	"github.com/ashtkn/go_todo_app/auth"
	"github.com/ashtkn/go_todo_app/entity"
	"github.com/ashtkn/go_todo_app/store"
)

type ListTask struct {
	DB   store.Queryer
	Repo TaskLister
}

func (l *ListTask) ListTask(ctx context.Context) (entity.Tasks, error) {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found in context")
	}

	tasks, err := l.Repo.ListTasks(ctx, l.DB, id)
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}
	return tasks, nil
}
