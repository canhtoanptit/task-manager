package handlers

import (
	"net/http"
	"task-management-api/internal/database"
	"task-management-api/internal/models"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	DueDate     time.Time `json:"due_date" binding:"required"`
	AssigneeID  uint      `json:"assignee_id" binding:"required"`
}

type TaskUpdateRequest struct {
	Status string `json:"status" binding:"required"`
}

func GetTasks(c *gin.Context) {
	assigneeID := c.Query("assignee_id")
	status := c.Query("status")
	sortBy := c.Query("sort_by")
	if sortBy == "" {
		sortBy = "created_at"
	}

	query := database.DB.Model(&models.Task{})

	if assigneeID != "" {
		query = query.Where("assignee_id = ?", assigneeID)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	var tasks []models.Task
	if err := query.Order(sortBy).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func GetMyTasks(c *gin.Context) {
	userID := c.GetUint("userID")
	var tasks []models.Task
	if err := database.DB.Where("assignee_id = ?", userID).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var req TaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		AssigneeID:  req.AssigneeID,
		CreatedBy:   c.GetUint("userID"),
		CreatedAt:   time.Now(),
	}

	if err := database.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	taskID := c.Param("id")
	userID := c.GetUint("userID")

	var req TaskUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var task models.Task
	if err := database.DB.Where("id = ? AND assignee_id = ?", taskID, userID).
		First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	task.Status = req.Status
	if err := database.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func GetEmployeeSummary(c *gin.Context) {
	type Summary struct {
		EmployeeID     uint   `json:"employee_id"`
		Email          string `json:"email"`
		TotalTasks     int64  `json:"total_tasks"`
		CompletedTasks int64  `json:"completed_tasks"`
	}

	var summaries []Summary
	if err := database.DB.Raw(`
		SELECT 
			u.id as employee_id,
			u.email,
			COUNT(t.id) as total_tasks,
			SUM(CASE WHEN t.status = 'Completed' THEN 1 ELSE 0 END) as completed_tasks
		FROM users u
		LEFT JOIN tasks t ON u.id = t.assignee_id
		WHERE u.role = 'employee'
		GROUP BY u.id, u.email
	`).Scan(&summaries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, summaries)
}
