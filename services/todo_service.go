// services/todo_service.go
package services

import (
    "todo-api/models"
    "todo-api/repositories"
)

func GetAllTodos() ([]models.Todo, error) {
    return repositories.GetAllTodos()
}

func GetTodoByID(id uint) (models.Todo, error) {
    return repositories.GetTodoByID(id)
}

func CreateTodo(todo models.Todo) error {
    return repositories.CreateTodo(todo)
}

func UpdateTodo(todo models.Todo) error {
    return repositories.UpdateTodo(todo)
}

func DeleteTodoByID(id uint) error {
    return repositories.DeleteTodoByID(id)
}
