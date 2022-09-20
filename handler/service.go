package handler

import (
	"context"

	"github.com/ashtkn/go_todo_app/entity"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . ListTaskService AddTaskService

type ListTaskService interface {
	ListTask(ctx context.Context) (entity.Task, error)
}

type AddTaskService interface {
	AddTask(ctx context.Context, task entity.Task) (*entity.Task, error)
}
