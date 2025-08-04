package main

import (
    "net/http"
    "blog/v2/models"
    "blog/v2/controllers"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()  // router with default middleware installed
    // index route
    models.ConnectDatabase()
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello World",
        })
    })
    router.POST("/posts", controllers.CreatePost)
    router.GET("/posts", controllers.FindPosts)
    router.GET("/posts/:id", controllers.FindPost)
    router.PATCH("/posts/:id", controllers.UpdatePost)
    router.DELETE("/posts/:id", controllers.DeletePost)
    // run the server
    router.Run()
}
