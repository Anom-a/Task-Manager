package router

import (
	"Task-Manager/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter()*gin.Engine{
	r:= gin.Default()
	taskRoutes := r.Group("/tasks")
	{
		taskRoutes.POST("/", controllers.CreateTask)
		taskRoutes.GET("/", controllers.GetAllTasks)
		taskRoutes.GET("/:id", controllers.GetTaskById)
		taskRoutes.PUT("/:id", controllers.UpdateTask)
		taskRoutes.DELETE("/:id", controllers.DeleteTask)
	}
	return r
}