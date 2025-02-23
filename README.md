# Task Management API

A RESTful API for task management with role-based access control, supporting Employer and Employee roles.
Built with Go (Gin框架), PostgreSQL, and JWT authentication.

## Features
- **Employee Role:**
    - View assigned tasks
    - Update task status (Pending/In Progress/Completed)
- **Employer Role:**
    - Create and assign tasks
    - View all tasks with filtering and sorting
    - View employee task summaries

## Prerequisites
- Docker
- Docker Compose
- Go 1.21 (optional, for local development without Docker)

## Setup Instructions

### Using Docker Compose (Recommended)
1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd task-management-api
2. **Build and run:**
  ```bash
   docker-compose up --build
  ```
## API Usage
### Login
```curl -X POST http://localhost:8000/login \
-H "Content-Type: application/json" \
-d '{"email": "user@example.com", "password": "password"}'
```
Protected Endpoints
Use the token in the Authorization header for all /api/* endpoints:

Header Format:
```
Authorization: Bearer <jwt-token>
```