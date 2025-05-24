package handlers

// import (
// 	"net/http"
// 	"strconv"
// 	"student-feedback/internal/models"

// 	"github.com/gin-gonic/gin"
// )

// var feedbacks []models.Feedback

// func RegisterRoutes(router *gin.Engine) {
// 	router.POST("/submit-feedback", SubmitFeedback)
// 	router.GET("/admin/feedback-summary", GetFeedbackSummary)
// 	router.GET("/admin/course-feedback/:id", GetCourseFeedback)
// }

// func SubmitFeedback(c *gin.Context) {
// 	var fb models.Feedback
// 	if err := c.ShouldBindJSON(&fb); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	feedbacks = append(feedbacks, fb)
// 	c.JSON(http.StatusOK, gin.H{"message": "Feedback submitted successfully"})
// }

// func GetFeedbackSummary(c *gin.Context) {
// 	c.JSON(http.StatusOK, feedbacks)
// }

// func GetCourseFeedback(c *gin.Context) {
// 	courseID := c.Param("id")
// 	var result []models.Feedback
// 	for _, f := range feedbacks {
// 		if f.CourseID == atoi(courseID) {
// 			result = append(result, f)
// 		}
// 	}
// 	c.JSON(http.StatusOK, result)
// }

// func atoi(str string) int {
// 	i, _ := strconv.Atoi(str)
// 	return i
// }
