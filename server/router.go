package server

import (
    "github.com/gin-gonic/gin"

    "github.com/gabrielgmusskopf/crud_go/controllers"
)

func NewRouter() *gin.Engine {
    
    router := gin.Default()

    user := new(controllers.UserController)
    router.GET("/users", user.GetAll)
    router.POST("/users", user.Create)
    router.DELETE("/users/:id", user.Delete)
    router.GET("/users/:id", user.GetById)
    router.PUT("/users/:id", user.Update)

    return router
}
