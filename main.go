package main

import (
	"context"
	"log"

	"task-manager/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	InitFirebase()

	ctx := context.Background()

	dbURL := "https://fir-dcc97-default-rtdb.asia-southeast1.firebasedatabase.app"
	client, err := App.DatabaseWithURL(ctx, dbURL)
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}

	handlers.DB = client // ✅ correct

	r := gin.Default()

	// CORS FIX
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	r.GET("/tasks", handlers.GetTasks)
	r.POST("/tasks", handlers.CreateTask)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.DELETE("/tasks/:id", handlers.DeleteTask)

	r.Run(":8080")
}
