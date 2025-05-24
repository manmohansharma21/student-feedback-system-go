CREATE DATABASE feedback_db;

\c feedback_db

CREATE TABLE feedback (
    id SERIAL PRIMARY KEY,
    student_id INT,
    course_id INT,
    rating INT CHECK (rating >= 1 AND rating <= 5),
    comments TEXT,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
