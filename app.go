package main

import (
	"context"
	"fmt"
	"slices"
)

// App struct
type App struct {
	ctx context.Context
}

var allTodos []string

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// AddTodo add a new todo to the list
func (a *App) AddTodo(todo string) []string {
	fmt.Printf("Adding todo %s to index %v\n", todo, len(allTodos))

	allTodos = append(allTodos, todo)

	return allTodos
}

// RemoveTodo removes todo by index
func (a *App) RemoveTodo(index int) []string {
	fmt.Printf("Removing todo with index %v\n", index)

	if len(allTodos) == 0 {
		fmt.Println("Todos are empty")
		return []string{}
	}

	if len(allTodos) == 1 && index == 0 {
		allTodos = []string{}
	} else {
		allTodos = slices.Delete(allTodos, index, index+1)
	}

	return allTodos
}
