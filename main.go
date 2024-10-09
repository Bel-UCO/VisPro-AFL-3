package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	DateTime    time.Time `json:"datetime"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
}

const (
	StatusCreated   = "Created"
	StatusPending   = "Pending"
	StatusCompleted = "Completed"
)

var tasks = []Task{
	{
		DateTime:    time.Date(2024, time.October, 8, 14, 0, 0, 0, time.UTC),
		Description: "Example task 1",
		Status:      StatusCreated,
	},
	{
		DateTime:    time.Date(2024, time.October, 9, 12, 30, 0, 0, time.UTC),
		Description: "Example task 2",
		Status:      StatusPending,
	},
	{
		DateTime:    time.Date(2024, time.October, 9, 16, 30, 0, 0, time.UTC),
		Description: "Example task 3",
		Status:      StatusCompleted,
	},
}

// Function to Get All Tasks
func getTasks(c *gin.Context) {
	response := make([]struct {
		ID          int    `json:"id"`
		DateTime    string `json:"datetime"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}, len(tasks))

	for i, task := range tasks {
		response[i] = struct {
			ID          int    `json:"id"`
			DateTime    string `json:"datetime"`
			Description string `json:"description"`
			Status      string `json:"status"`
		}{
			ID:          i + 1, // ID based on index
			DateTime:    task.DateTime.Format("02-01-2006 15:04"),
			Description: task.Description,
			Status:      task.Status,
		}
	}
	c.JSON(http.StatusOK, response)
}

// Function to Get a Task by ID
func getTask(c *gin.Context) {
	id := c.Param("id")
	var taskID int
	if _, err := fmt.Sscanf(id, "%d", &taskID); err != nil || taskID < 1 || taskID > len(tasks) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	index := taskID - 1 // Convert ID to index
	if index < 0 || index >= len(tasks) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Return the found task
	response := struct {
		ID          int    `json:"id"`
		DateTime    string `json:"datetime"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}{
		ID:          taskID,
		DateTime:    tasks[index].DateTime.Format("02-01-2006 15:04"),
		Description: tasks[index].Description,
		Status:      tasks[index].Status,
	}

	c.JSON(http.StatusOK, response)
}

// Function to Create a Task
func createTask(c *gin.Context) {
	var newTask Task

	// Bind JSON input to the newTask variable
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate status
	if newTask.Status != StatusCreated && newTask.Status != StatusPending && newTask.Status != StatusCompleted {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status. Allowed values are 'Created', 'Pending', or 'Completed'."})
		return
	}

	// Set the current time for DateTime if not provided
	if newTask.DateTime.IsZero() {
		newTask.DateTime = time.Now().UTC() // Gunakan waktu saat ini jika tidak ada
	}

	// Add the new task to the global slice
	tasks = append(tasks, newTask)

	// Respond with the created task
	c.JSON(http.StatusCreated, struct {
		ID          int    `json:"id"`
		DateTime    string `json:"datetime"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}{
		ID:          len(tasks), // ID based on current length
		DateTime:    newTask.DateTime.Format("02-01-2006 15:04"),
		Description: newTask.Description,
		Status:      newTask.Status,
	})
}

// Function to Update a Task
func updateTask(c *gin.Context) {
	id := c.Param("id")
	var taskID int
	if _, err := fmt.Sscanf(id, "%d", &taskID); err != nil || taskID < 1 || taskID > len(tasks) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	index := taskID - 1 // Convert ID to index
	if index < 0 || index >= len(tasks) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Update status based on current status
	switch tasks[index].Status {
	case StatusCreated:
		tasks[index].Status = StatusPending
	case StatusPending:
		tasks[index].Status = StatusCompleted
	case StatusCompleted:
		c.JSON(http.StatusForbidden, gin.H{"error": "Task is already completed and cannot be updated"})
		return
	}

	// Format the DateTime for the response
	response := struct {
		ID          int    `json:"id"`
		DateTime    string `json:"datetime"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}{
		ID:          taskID,
		DateTime:    tasks[index].DateTime.Format("02-01-2006 15:04"),
		Description: tasks[index].Description,
		Status:      tasks[index].Status,
	}

	c.JSON(http.StatusOK, response) // Return the updated task
}

// Function to Delete a Task
func deleteTask(c *gin.Context) {
	id := c.Param("id")
	var taskID int
	if _, err := fmt.Sscanf(id, "%d", &taskID); err != nil || taskID < 1 || taskID > len(tasks) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	index := taskID - 1 // Convert ID to index
	if index < 0 || index >= len(tasks) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Remove the task
	tasks = append(tasks[:index], tasks[index+1:]...)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func main() {
	router := gin.New()

	router.GET("/tasks", getTasks)         // Route to get all tasks
	router.GET("/task/:id", getTask)       // Route to get task by ID
	router.POST("/task", createTask)       // Route to create new task
	router.PUT("/task/:id", updateTask)    // Route to update task status
	router.DELETE("/task/:id", deleteTask) // Route to delete task

	router.Run(":8080")
}