// config/config.go
package config

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "todo-api/models"
    "log"
)

var DB *gorm.DB

func SetupDatabase() {
    var err error
    DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("データベース接続失敗:", err)
    }

    // マイグレーション（Todoモデル）
    DB.AutoMigrate(&models.Todo{})
}
