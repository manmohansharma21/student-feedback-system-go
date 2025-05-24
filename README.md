# Student Feedback Management System

A web-based feedback system using Go and SQL with a modular and scalable architecture.

## 📦 Folder Structure
```
.
├── cmd
│   └── server
│       └── main.go
├── internal
│   ├── db
│   │   └── db.go
│   ├── handlers
│   │   └── feedback.go
│   └── models
│       └── models.go
├── schema.sql
├── README.md
```

## 🚀 Setup Instructions
1. Install Go (https://golang.org/doc/install)
2. Run:
```bash
cd cmd/server
go mod init student-feedback
go get github.com/gin-gonic/gin
go run main.go
```

## 🔗 API Endpoints

### POST /submit-feedback
**Request Body:**
```json
{
  "student_id": 1,
  "course_id": 2,
  "rating": 5,
  "comments": "Excellent session!"
}
```

### GET /admin/feedback-summary
Returns all feedback entries.

### GET /admin/course-feedback/:id
Returns feedback filtered by course ID.
