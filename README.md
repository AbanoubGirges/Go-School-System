# Go School System

A school management REST API built with Go, Gin, GORM, and PostgreSQL.

## Overview

This project provides a backend service for managing users, classes, students, and attendance records.

Features:
- User registration and login with JWT-based authentication
- Role-based access control for admin endpoints
- Class creation, update, retrieval, and deletion
- Student creation, retrieval, and attendance tracking
- PostgreSQL database integration

## Tech stack

- Go 1.26.3
- Gin web framework
- GORM ORM
- PostgreSQL (via `gorm.io/driver/postgres`)
- JWT authentication
- Docker-ready build

## Project structure

- `cmd/main.go` - application entry point
- `internal/config` - environment and database configuration
- `internal/controllers` - HTTP controller logic
- `internal/repo` - data access layer
- `internal/models` - GORM models
- `internal/dto` - request DTOs and model converters
- `internal/routes` - route definitions
- `internal/middleware` - JWT authentication and authorization

## Environment

Create a `.env` file at the project root with the following values:

```env
PORT=8080
DATABASE_URL=postgres://user:password@localhost:5432/database_name
JWT_SECRET=your_jwt_secret
```

## Run locally

1. Download dependencies:

```powershell
go mod download
```

2. Start the server:

```powershell
go run cmd/main.go
```

3. The API listens on `http://localhost:<PORT>`.

## Docker

Build the Docker image:

```powershell
docker build -t go-school-system .
```

Run the container:

```powershell
docker run -p 8080:8080 --env-file .env go-school-system
```

## API Reference

Base URL: `/api/v1`

### Authentication

#### Register

- `POST /api/v1/register`
- Content type: `multipart/form-data`
- Body fields:
  - `name` (string, required)
  - `phoneNumber` (string, required)
  - `password` (string, required)
  - `email` (string, required)
  - `image` (file, required)

Response:
- `201 Created` on success
- `400` when payload is invalid
- `500` when upload or creation fails

#### Login

- `POST /api/v1/login`
- Content type: `application/json`
- Body:
  - `phoneNumber` (string, required)
  - `password` (string, required)

Response:
- `200 OK` with JSON containing `user` and `token`
- `401` for invalid credentials
- `404` if the user is not found

### Profile

#### Edit profile

- `PATCH /api/v1/profile`
- Authorization: `Bearer <token>`
- Body: JSON object with update fields

Response:
- `200 OK` on success
- `401` if token is missing or invalid

### Classes

All class endpoints require admin authorization unless otherwise noted.

#### Create class

- `POST /api/v1/classes`
- Authorization: `Bearer <token>`
- Body:
  - `name` (string)
  - `grade` (string)

Response:
- `201 Created`

#### Update class

- `PUT /api/v1/classes/:id`
- Authorization: `Bearer <token>`
- Path parameter: class ID

Response:
- `200 OK`
- `400` for invalid ID

#### Get all classes

- `GET /api/v1/classes`
- Authorization: `Bearer <token>`

Response:
- `200 OK` with array of classes

#### Get class by ID

- `GET /api/v1/classes/:id`
- Authorization: `Bearer <token>`

Response:
- `200 OK` with class object
- `404` if class not found

#### Delete class

- `DELETE /api/v1/classes/:id`
- Authorization: `Bearer <token>`

Response:
- `200 OK` on success

#### Get current class (non-admin)

- `GET /api/v1/classes`
- Authorization: `Bearer <token>`

This endpoint returns the class associated with the authenticated user.

### Students

#### Create student

- `POST /api/v1/students`
- Authorization: `Bearer <token>`
- Body JSON:
  - `name` (string, required)
  - `age` (int, required)
  - `class_id` (UUID, required)
  - `phone_number` (string, required)
  - `email` (string, required)
  - `location` (string, required)
  - `coordinates` (string, required)
  - `birthdate` (string, required)

Response:
- `201 Created`

#### Get all students

- `GET /api/v1/students?classId=<uuid>`
- Authorization: admin token

Response:
- `200 OK` with student list

#### Get student by ID

- `GET /api/v1/students/:id`
- Authorization: admin token

Response:
- `200 OK` with student object

#### Create attendance

- `POST /api/v1/students/attendance`
- Authorization: admin token
- Body JSON:
  - `studentId` (string, required)
  - `date` (RFC3339 timestamp, required)
  - `present` (boolean)

Response:
- `201 Created` with attendance record

#### Get attendance by student

- `GET /api/v1/students/:id/attendance`
- Authorization: admin token

Response:
- `200 OK` with attendance records

## Data models

### User

- `id` (UUID)
- `name` (string)
- `phoneNumber` (string)
- `password` (string, hashed)
- `email` (string)
- `role` (string)
- `class` (UUID)
- `status` (string)

### Class

- `id` (UUID)
- `name` (string)
- `grade` (string)

### Student

- `id` (UUID)
- `name` (string)
- `phoneNumber` (string)
- `email` (string)
- `class` (UUID)
- `location` (string)
- `coordinates` (string)
- `age` (uint)
- `birthdate` (string)

### Attendance

- `id` (UUID)
- `studentId` (UUID)
- `date` (timestamp)
- `present` (boolean)

## Notes

- The service uses `Authorization: Bearer <token>` for protected routes.
- The `register` endpoint expects an `image` file upload.
- The app auto-migrates models on startup.
