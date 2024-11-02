// repositories/todo_repository.go
package repositories

import (
    "todo-api/config"
    "todo-api/models"
)

func GetAllTodos() ([]models.Todo, error) {
    var todos []models.Todo
    result := config.DB.Find(&todos)
    return todos, result.Error
}

func GetTodoByID(id uint) (models.Todo, error) {
    var todo models.Todo
    result := config.DB.First(&todo, id)
    return todo, result.Error
}

func CreateTodo(todo models.Todo) error {
    result := config.DB.Create(&todo)
    return result.Error
}

func UpdateTodo(todo models.Todo) error {
    result := config.DB.Save(&todo)
    return result.Error
}

func DeleteTodoByID(id uint) error {
    result := config.DB.Delete(&models.Todo{}, id)
    return result.Error
}
