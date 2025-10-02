package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "go-prcatice2/internal/middleware"
    "go-prcatice2/internal/user/repository"
    "go-prcatice2/internal/user/service"
    userHttp "go-prcatice2/internal/user/transport/http"
    "go-prcatice2/pkg/db"
)

func main() {
    database, err := db.InitDB()
    if err != nil {
        log.Fatal(err)
    }

    repo := repository.NewUserRepository(database)
    svc := service.NewUserService(repo)
    handler := userHttp.NewUserHandler(svc)

    r := gin.Default()
    r.Use(middleware.APIKeyMiddleware())

    handler.Register(r)

    r.Run(":8080")
}
