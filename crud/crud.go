package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Task represents a task model in our database
type Task struct {
	ID          uint   `gorm:"primaryKey"`  // ID is the primary key
	Title       string `json:"title"`       // Title of the task
	Description string `json:"description"` // Description of the task
	Status      string `json:"status"`      // Status of the task (pending, completed, etc.)
}

// DB variable to hold our database connection
var DB *gorm.DB

// Initialize MySQL Database connection
func InitDB() {
	dsn := "user:password@tcp(127.0.0.1:3306)/taskdb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Auto-migrate our Task struct into a table in the database
	DB.AutoMigrate(&Task{})
}

// CreateTask creates a new task
func CreateTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&task)
	c.JSON(http.StatusCreated, task)
}

// GetTasks retrieves all tasks
func GetTasks(c *gin.Context) {
	var tasks []Task
	DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

// GetTaskByID retrieves a task by its ID
func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	var task Task
	if err := DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// UpdateTask updates an existing task by ID
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task Task
	if err := DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Save(&task)
	c.JSON(http.StatusOK, task)
}

// DeleteTask deletes a task by ID
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := DB.Delete(&Task{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func main() {
	// Initialize the database connection
	InitDB()

	// Create a Gin router instance
	r := gin.Default()

	// Define routes for CRUD operations
	r.POST("/tasks", CreateTask)
	r.GET("/tasks", GetTasks)
	r.GET("/tasks/:id", GetTaskByID)
	r.PUT("/tasks/:id", UpdateTask)
	r.DELETE("/tasks/:id", DeleteTask)

	// Start the server on port 8080
	r.Run(":8080")
}

// *****************************************Explanation:

// This is a gorm model, representing the structure of a task. It will be translated into a database table when gorm's AutoMigrate function is called.
// Fields: ID, Title, Description, Status.
// Database Setup (InitDB):

// We set up the MySQL connection string using the gorm.Open() function.
// AutoMigrate is used to ensure that the Task struct is synced with the database.
// CRUD Handlers:

// CreateTask: Handles POST /tasks, creates a new task in the database.
// GetTasks: Handles GET /tasks, retrieves all tasks from the database.
// GetTaskByID: Handles GET /tasks/:id, retrieves a specific task by ID.
// UpdateTask: Handles PUT /tasks/:id, updates a task with the given ID.
// DeleteTask: Handles DELETE /tasks/:id, deletes a task by ID.

// ***************************************Summary
// Gin helps build web services with routing and middleware support.
// gorm simplifies database interactions, especially CRUD operations, by mapping Go structs to database tables.
// By combining Gin and gorm, you can quickly build robust and scalable REST APIs in Go.

// In Go, Gin is a high-performance HTTP web framework, and gorm is an ORM (Object Relational Mapping) library that simplifies interactions with databases. When working together, Gin handles routing and HTTP requests, while gorm helps manage database interactions, allowing developers to perform CRUD (Create, Read, Update, Delete) operations easily.

// ****************************************1. Gin Framework
// Gin is a lightweight framework for building web applications and REST APIs. It provides a simple way to handle routing, middleware, and HTTP requests.

// Router: Gin’s router is the core component for handling HTTP requests. It matches the HTTP method (e.g., GET, POST, PUT, DELETE) with specific routes.
// Middleware: Gin allows easy integration of middleware to handle cross-cutting concerns such as logging, authentication, or rate-limiting.
// Handlers: These are functions that are triggered when a specific route is hit.

// ****************************************2. gorm ORM
// gorm is a Go library that simplifies working with databases like MySQL, PostgreSQL, etc. It allows you to map Go structs to database tables and provides methods for performing CRUD operations.
