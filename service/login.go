package service

import (
	"context"

	"github.com/ashtkn/go_todo_app/store"
)

type Login struct {
	DB             store.Queryer
	Repo           UserGetter
	TokenGenerator TokenGenerator
}

func (l *Login) Login(ctx context.Context, name, pw string) (string, error) {
	u, err := l.Repo.GetUser(ctx, l.DB, name)
	if err != nil {
		return "", err
	}

	if err := u.ComparePassword(pw); err != nil {
		return "", err
	}

	token, err := l.TokenGenerator.GenerateToken(ctx, *u)
	if err != nil {
		return "", err
	}

	return string(token), nil
}
