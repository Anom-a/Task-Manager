package controllers

import (
	"Task-Manager/data"
	"Task-Manager/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	err := c.ShouldBindBodyWithJSON(&task)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	// handles id is empty or the name is empty
	if task.ID == "" || task.Name == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID or name are required"})
		return
	}
	if task.DueDate.IsZero() {
		task.DueDate = time.Now().Add(24 * time.Hour)
	}
	err = data.CreateTask(task)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func GetAllTasks(c *gin.Context) {
	tasks, err := data.GetAllTasks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tasks from mongodb local server"})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task, err := data.GetTaskById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updated models.Task
	if err := c.ShouldBindBodyWithJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	err := data.UpdateTask(id, updated)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found or update failed"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := data.DeleteTask(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
