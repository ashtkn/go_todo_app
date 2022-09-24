package main

import (
	"context"
	"net/http"

	"github.com/ashtkn/go_todo_app/auth"
	"github.com/ashtkn/go_todo_app/clock"
	"github.com/ashtkn/go_todo_app/config"
	"github.com/ashtkn/go_todo_app/handler"
	"github.com/ashtkn/go_todo_app/service"
	"github.com/ashtkn/go_todo_app/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()

	// GET /health
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	// Database and Repository
	db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	clocker := clock.RealClocker{}
	r := store.Repository{Clocker: clocker}

	kvsDb, err := store.NewKVS(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}

	// JWTer
	jwter, err := auth.NewJWTer(kvsDb, clocker)

	// Validator
	v := validator.New()

	// GET/POST Tasks
	at := &handler.AddTask{Service: &service.AddTask{DB: db, Repo: &r}, Validator: v}
	lt := &handler.ListTask{Service: &service.ListTask{DB: db, Repo: &r}}
	mux.Route("/tasks", func(r chi.Router) {
		r.Use(handler.AuthMiddleware(jwter))
		mux.Post("/tasks", at.ServeHTTP)
		mux.Get("/tasks", lt.ServeHTTP)
	})

	// POST /register
	ru := &handler.RegisterUser{Service: &service.RegisterUser{DB: db, Repo: &r}, Validator: v}
	mux.Post("/register", ru.ServeHTTP)

	// POST /login

	if err != nil {
		return nil, cleanup, err
	}
	lu := &handler.Login{Service: &service.Login{DB: db, Repo: &r, TokenGenerator: jwter}, Validator: v}
	mux.Post("/login", lu.ServeHTTP)

	return mux, cleanup, nil
}
