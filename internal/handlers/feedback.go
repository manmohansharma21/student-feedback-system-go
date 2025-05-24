package handlers

import (
	"net/http"
	"student-feedback/internal/db"
	"student-feedback/internal/models"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/submit-feedback", SubmitFeedback)
	router.GET("/admin/feedback-summary", GetFeedbackSummary)
	router.GET("/admin/course-feedback/:id", GetCourseFeedback)
	router.DELETE("/admin/course-feedback/:id", DeleteCourseFeedback)
}

func SubmitFeedback(c *gin.Context) {
	var fb models.Feedback
	if err := c.ShouldBindJSON(&fb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.DB.Exec(`INSERT INTO feedback (student_id, course_id, rating, comments) VALUES ($1, $2, $3, $4)`,
		fb.StudentID, fb.CourseID, fb.Rating, fb.Comments)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert feedback"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Feedback submitted to DB"})
}

func GetFeedbackSummary(c *gin.Context) {
	rows, err := db.DB.Query(`SELECT student_id, course_id, rating, comments FROM feedback`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB fetch failed"})
		return
	}
	defer rows.Close()

	var result []models.Feedback
	for rows.Next() {
		var f models.Feedback
		if err := rows.Scan(&f.StudentID, &f.CourseID, &f.Rating, &f.Comments); err != nil {
			continue
		}
		result = append(result, f)
	}

	c.JSON(http.StatusOK, result)
}

func GetCourseFeedback(c *gin.Context) {
	courseID := c.Param("id")
	rows, err := db.DB.Query(`SELECT student_id, course_id, rating, comments FROM feedback WHERE course_id = $1`, courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB query failed"})
		return
	}
	defer rows.Close()

	var result []models.Feedback
	for rows.Next() {
		var f models.Feedback
		if err := rows.Scan(&f.StudentID, &f.CourseID, &f.Rating, &f.Comments); err != nil {
			continue
		}
		result = append(result, f)
	}

	c.JSON(http.StatusOK, result)
}

func DeleteCourseFeedback(c *gin.Context) {
	courseID := c.Param("id")
	_, err := db.DB.Exec(`DELETE FROM feedback WHERE course_id = $1`, courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Feedback deleted for course"})
}
