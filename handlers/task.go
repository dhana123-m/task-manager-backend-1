package handlers

import (
	"context"
	"net/http"

	"task-manager/models"

	"firebase.google.com/go/v4/db"
	"github.com/gin-gonic/gin"
)

var DB *db.Client

// GET
func GetTasks(c *gin.Context) {
	ctx := context.Background()

	var tasks map[string]models.Task
	err := DB.NewRef("tasks").Get(ctx, &tasks)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// CREATE
func CreateTask(c *gin.Context) {
	var task models.Task
	c.BindJSON(&task)

	ctx := context.Background()

	ref, err := DB.NewRef("tasks").Push(ctx, task)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	task.ID = ref.Key

	DB.NewRef("tasks/"+task.ID).Set(ctx, task)

	c.JSON(http.StatusOK, task)
}

// UPDATE
func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task
	c.BindJSON(&task)

	ctx := context.Background()

	updateData := map[string]interface{}{
		"title":    task.Title,
		"assignee": task.Assignee,
		"priority": task.Priority,
		"status":   task.Status,
		"dueDate":  task.DueDate,
	}

	err := DB.NewRef("tasks/"+id).Update(ctx, updateData)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}

// DELETE
func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	ctx := context.Background()

	err := DB.NewRef("tasks/" + id).Delete(ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
