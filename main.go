// main.go
package main

import (
    "todo-api/config"
    "todo-api/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // データベース設定
    config.SetupDatabase()

    // ルート設定
    routes.SetupRoutes(r)

    r.Run(":8080") // ポート8080で起動
}
