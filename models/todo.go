// models/todo.go
package models

import "gorm.io/gorm"

type Todo struct {
    gorm.Model
    Title       string `json:"title" binding:"required"`
    Description string `json:"description"`
    Status      string `json:"status" binding:"required"`
}