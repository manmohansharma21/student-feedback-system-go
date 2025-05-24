package models

type Feedback struct {
    StudentID int    `json:"student_id"`
    CourseID  int    `json:"course_id"`
    Rating    int    `json:"rating"`
    Comments  string `json:"comments"`
}
