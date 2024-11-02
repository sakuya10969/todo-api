// controllers/todo_controller.go
package controllers

import (
    "todo-api/models"
    "todo-api/services"
    "todo-api/utils"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

func GetTodos(c *gin.Context) {
    todos, err := services.GetAllTodos()
    if err != nil {
        utils.RespondError(c, http.StatusInternalServerError, err.Error())
        return
    }
    utils.RespondJSON(c, http.StatusOK, todos)
}

func GetTodoByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    todo, err := services.GetTodoByID(uint(id))
    if err != nil {
        utils.RespondError(c, http.StatusNotFound, "Todo not found")
        return
    }
    utils.RespondJSON(c, http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
    var todo models.Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        utils.RespondError(c, http.StatusBadRequest, err.Error())
        return
    }

    if err := services.CreateTodo(todo); err != nil {
        utils.RespondError(c, http.StatusInternalServerError, err.Error())
        return
    }
    utils.RespondJSON(c, http.StatusCreated, todo)
}

func UpdateTodo(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var todo models.Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        utils.RespondError(c, http.StatusBadRequest, err.Error())
        return
    }

    todo.ID = uint(id)
    if err := services.UpdateTodo(todo); err != nil {
        utils.RespondError(c, http.StatusInternalServerError, err.Error())
        return
    }
    utils.RespondJSON(c, http.StatusOK, todo)
}

func DeleteTodoByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := services.DeleteTodoByID(uint(id)); err != nil {
        utils.RespondError(c, http.StatusInternalServerError, err.Error())
        return
    }
    c.Status(http.StatusNoContent)
}
