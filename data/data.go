package data

import (
	"net/http"
	"time"

	"github.com/Anom-a/Task-Manager/models"
	"github.com/gin-gonic/gin"
)

var tasks = []models.Task{
	{ID: "1", Title: "Task 1", Description: "This is task 1 ", DueDate: time.Now(), Status: "ongoing"},
	{ID: "2", Title: "Task 2", Description: "This is task 2 ", DueDate: time.Now(), Status: "not-started"},
	{ID: "3", Title: "Task 3", Description: "This is task 3 ", DueDate: time.Now(), Status: "completed"},
}

var allowedStatus = map[string]bool{
	"ongoing":     true,
	"not-started": true,
	"completed":   true,
}

// GetTasks returns all tasks
func GetTasks(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, tasks)
}

// GetTaskById returns a specific task by an id
func GetTaskById(ctx *gin.Context) {
	// retrieve a URL parameter
	id := ctx.Param("id")
	for _, task := range tasks {
		if id == task.ID {
			ctx.IndentedJSON(http.StatusOK, task)
			return
		}
	}
	// returning an error message if the task is not found
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task Not Found"})
}

// UpdateTask updates a task
func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var updated models.Task
	// checking the error if there is incosistency between task data fields
	if err := ctx.ShouldBindJSON(&updated); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !allowedStatus[updated.Status] {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid status"})
		return
	}
	for i, task := range tasks {
		if id == task.ID {
			tasks[i] = updated
			ctx.IndentedJSON(http.StatusAccepted, gin.H{"tasks": updated})
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "The task with the specified id doesn't exist"})
}

// DeleteTask deletes a task
func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...) // deletes tasks[i]
			ctx.IndentedJSON(http.StatusAccepted, gin.H{"message": "Deletion Completed"})
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "The task with the specified id doesn't exist"})
}

// AddTask adds a new task
func AddTask(ctx *gin.Context) {
	var newTask models.Task
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// check if there's a duplicate id
	for _, task := range tasks {
		if task.ID == newTask.ID {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "There is already a task associated with this id"})
			return
		}
	}
	// vallidating status
	if !allowedStatus[newTask.Status] {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Task status is invalid"})
		return
	}
	tasks = append(tasks, newTask)
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Task successfuly added", "task": newTask})
}
