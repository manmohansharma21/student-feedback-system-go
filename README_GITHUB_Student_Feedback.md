
# Student Feedback Management System

A modular web-based backend system written in **Go (Golang)** with **SQL schema** support for managing course feedback by students. Built using the Gin web framework with a clean project structure.

## 🔧 Features

- Submit course feedback
- Admin view of all feedback
- Filter feedback by course ID
- Scalable modular folder layout

---

## 🗂️ Folder Structure

```
.
├── cmd
│   └── server
│       └── main.go              # Entry point
├── internal
│   ├── db
│   │   └── db.go                # PostgreSQL DB connection handler
│   ├── handlers
│   │   └── feedback.go          # Route handlers
│   └── models
│       └── models.go            # Feedback struct definition
├── migrations                  # SQL migration scripts
│   ├── 0001_create_feedback_table.up.sql
│   └── 0001_create_feedback_table.down.sql
├── schema.sql                  # (optional) flat schema
├── .env                        # DB credentials and environment
├── README.md                   # Setup & usage guide
```

---

## 🚀 Getting Started

1. Install Go: https://golang.org/doc/install  
2. Clone the repository

```bash
git clone https://github.com/manmohansharma21/student-feedback-system-go.git
cd student-feedback-system-go/cmd/server
```

3. Initialize Go module and install dependencies

```bash
go mod init student-feedback
go get github.com/gin-gonic/gin github.com/lib/pq github.com/joho/godotenv
```

4. Run the application

```bash
go run main.go
```

---

## 🔗 API Endpoints

### POST `/submit-feedback`

**Request Body**
```json
{
  "student_id": 1,
  "course_id": 2,
  "rating": 5,
  "comments": "Great course!"
}
```

### GET `/admin/feedback-summary`
Returns all submitted feedback.

### GET `/admin/course-feedback/:id`
Returns feedback filtered by course ID.

---

## 🧠 Tech Stack

- Go (Golang)
- Gin framework
- PostgreSQL (real DB connection)
- JSON APIs
- `.env` environment loading

---

## 👤 Author

**Manmohan Sharma**  
GitHub: [@manmohansharma21](https://github.com/manmohansharma21)

---

## 📜 License

MIT License - feel free to reuse with attribution.

---

## 🗃️ Database Configuration (.env)

Create a `.env` file in your project root with the following content:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=admin
DB_NAME=feedback_db
```

This will be loaded automatically by the app using `github.com/joho/godotenv`.

---

## 🔄 Database Migrations (Using golang-migrate)

Migration scripts are stored under:

```
/migrations
├── 0001_create_feedback_table.up.sql
├── 0001_create_feedback_table.down.sql
```

### 📥 Install CLI (if not already):

```bash
brew install golang-migrate
# or for Linux
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin
```

### ▶️ Apply Migrations:

```bash
migrate -path ./migrations -database "postgres://postgres:admin@localhost:5432/feedback_db?sslmode=disable" up
```

### ⬅️ Rollback Migrations:

```bash
migrate -path ./migrations -database "postgres://postgres:admin@localhost:5432/feedback_db?sslmode=disable" down
```
