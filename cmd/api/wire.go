package main

/*
Wire.go - Manual Dependency Injection

This file documents the dependency injection wiring for the application.
Currently using manual DI for simplicity. For larger applications,
consider using Google Wire for automated DI generation.

Dependency Graph:
-----------------
configs.Config
    └── database.NewPostgresDB
            └── postgres.NewTodoRepository (implements contract.TodoRepository)
                    └── todo.NewService
                            └── handler.NewHandler
                                    └── handler.RegisterRoutes

How to use Google Wire (optional):
----------------------------------
1. Install wire: go install github.com/google/wire/cmd/wire@latest
2. Add wire build tags to this file
3. Define provider functions
4. Run: wire ./cmd/api

Example wire.go with build tags:

	//go:build wireinject
	// +build wireinject

	package main

	import (
		"github.com/google/wire"
		...
	)

	func InitializeApp() (*App, error) {
		wire.Build(
			configs.LoadConfig,
			database.NewPostgresDB,
			postgres.NewTodoRepository,
			todo.NewService,
			handler.NewHandler,
		)
		return &App{}, nil
	}
*/
