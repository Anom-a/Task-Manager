
# 📝 Task Manager API Documentation

This API allows you to **create**, **read**, **update**, and **delete** tasks, with persistent storage using **MongoDB** and the **Mongo Go Driver**.

## 🌐 Base URL

```
http://localhost:8080
```

## 🔌 Endpoints

### 📥 Create a Task

**POST** `/tasks/`

Creates a new task.

#### 🔸 Request Body (JSON)

```json
{
  "id": "1",
  "name": "Learn Go",
  "description": "Study the Mongo driver",
  "due_date": "2025-08-10T00:00:00Z",
  "status": "ongoing"
}
```

> `id` and `name` are required. `due_date` is optional and defaults to 24h from creation.

#### ✅ Response

```json
{
  "id": "1",
  "name": "Learn Go",
  "description": "Study the Mongo driver",
  "due_date": "2025-08-10T00:00:00Z",
  "status": "ongoing"
}
```

### 📄 Get All Tasks

**GET** `/tasks/`

Retrieves a list of all tasks.

#### ✅ Response

```json
[
  {
    "id": "1",
    "name": "Learn Go",
    "description": "Study the Mongo driver",
    "due_date": "2025-08-10T00:00:00Z",
    "status": "ongoing"
  }
]
```

### 🔍 Get a Task by ID

**GET** `/tasks/:id`

Retrieves details for a specific task by ID.

#### ✅ Response

```json
{
  "id": "1",
  "name": "Learn Go",
  "description": "Study the Mongo driver",
  "due_date": "2025-08-10T00:00:00Z",
  "status": "ongoing"
}
```

#### ❌ Error Response

```json
{
  "error": "Task not found"
}
```

### ✏️ Update a Task

**PUT** `/tasks/:id`

Updates an existing task.

#### 🔸 Request Body (JSON)

```json
{
  "name": "Master Go",
  "description": "Read docs",
  "due_date": "2025-08-15T00:00:00Z",
  "status": "done"
}
```

#### ✅ Response

```json
{
  "message": "Task updated successfully"
}
```

#### ❌ Error Response

```json
{
  "error": "Task not found or update failed"
}
```

### ❌ Delete a Task

**DELETE** `/tasks/:id`

Deletes a task by ID.

#### ✅ Response

```json
{
  "message": "Task deleted successfully"
}
```

#### ❌ Error Response

```json
{
  "error": "Task not found"
}
```

## 🛠️ Setup & Configuration

### 📦 Requirements

- Go 1.20+
- MongoDB (local or cloud)
- Go modules installed

### 🔧 MongoDB Connection

In `data/task_service.go`:

```go
const mongoURI = "mongodb://localhost:27017"
```

> Change this if you're using a MongoDB Atlas URI.

### ▶️ Run the App

```bash
go run main.go
```

Visit: [http://localhost:8080](http://localhost:8080)

## 🧪 Testing

Use Postman or curl to send requests to:

- `POST /tasks/` – Add a task
- `GET /tasks/` – Get all tasks
- `GET /tasks/:id` – Get one task
- `PUT /tasks/:id` – Update
- `DELETE /tasks/:id` – Delete

You can also verify in **MongoDB Compass** or with:

```bash
mongo
use taskdb
db.tasks.find().pretty()
```