package router

import (
	"github.com/Anom-a/Task-Manager/data"
	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	router := gin.Default()
	router.GET("/tasks", data.GetTasks)
	router.GET("/tasks/:id", data.GetTaskById)
	router.POST("/tasks", data.AddTask)
	router.PUT("/tasks/:id", data.UpdateTask)
	router.DELETE("/tasks/:id", data.DeleteTask)
	router.Run()
}
