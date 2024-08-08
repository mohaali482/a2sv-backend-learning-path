# Task API Documentation

# Tasks

## Get All Tasks

### Request

`GET /tasks`

### Response

**Status Code:** 200 OK

**Response Body:**

```json
[
  {
    "id": "63e4d8b9f03d3b4b981f5e1c",
    "user_id": "63e4d8b9f03d3b4b981f5e1d",
    "title": "Task 1",
    "description": "This is the description for Task 1",
    "datetime": "2023-02-08T12:34:56.789Z",
    "done": false
  },
  {
    "id": "63e4d8b9f03d3b4b981f5e1e",
    "user_id": "63e4d8b9f03d3b4b981f5e1d",
    "title": "Task 2",
    "description": "This is the description for Task 2",
    "datetime": "2023-02-09T10:11:12.345Z",
    "done": true
  }
]
```

## Get Task by ID

### Request

`GET /tasks/:id`

**Path Parameters:** - `id` (string, required): The ID of the task to retrieve.

### Response

**Status Code:** 200 OK

**Response Body:**

```json
{
  "id": "63e4d8b9f03d3b4b981f5e1c",
  "user_id": "63e4d8b9f03d3b4b981f5e1d",
  "title": "Task 1",
  "description": "This is the description for Task 1",
  "datetime": "2023-02-08T12:34:56.789Z",
  "done": false
}
```

**Status Code:** 400 Bad Request

**Response Body:**

```json
{
  "error": "invalid id"
}
```

**Status Code:** 404 Not Found

**Response Body:**

```json
{
  "error": "task not found"
}
```

## Update Task

### Request

`PUT /tasks/:id`

**Path Parameters:** - `id` (string, required): The ID of the task to update.

**Request Body:**

```json
{
  "user_id": "63e4d8b9f03d3b4b981f5e1d",
  "title": "Updated Task 1",
  "description": "This is the updated description for Task 1",
  "datetime": "2023-02-08T12:34:56.789Z",
  "done": true
}
```

### Response

**Status Code:** 200 OK

**Response Body:**

```json
{
  "id": "63e4d8b9f03d3b4b981f5e1c",
  "user_id": "63e4d8b9f03d3b4b981f5e1d",
  "title": "Updated Task 1",
  "description": "This is the updated description for Task 1",
  "datetime": "2023-02-08T12:34:56.789Z",
  "done": true
}
```

**Status Code:** 400 Bad Request

**Response Body:**

```json
{
  "error": "invalid id"
}
```

**Status Code:** 404 Not Found

**Response Body:**

```json
{
  "error": "task not found"
}
```

## Delete Task

### Request

`DELETE /tasks/:id`

**Path Parameters:** - `id` (string, required): The ID of the task to delete.

### Response

**Status Code:** 200 OK

**Response Body:**

```json
{
  "message": "task deleted"
}
```

**Status Code:** 400 Bad Request

**Response Body:**

```json
{
  "error": "invalid id"
}
```

**Status Code:** 404 Not Found

**Response Body:**

```json
{
  "error": "task not found"
}
```

## Create Task

### Request

`POST /tasks`

**Request Body:**

```json
{
  "id": "63e4d8b9f03d3b4b981f5e1c",
  "user_id": "63e4d8b9f03d3b4b981f5e1d",
  "title": "New Task",
  "description": "This is a new task",
  "datetime": "2023-02-08T12:34:56.789Z",
  "done": true
}
```

### Response

**Status Code:** 201 Created

**Response Body:**

```json
{
  "id": "63e4d8b9f03d3b4b981f5e1f",
  "user_id": "63e4d8b9f03d3b4b981f5e1d",
  "title": "New Task",
  "description": "This is a new task",
  "datetime": "2023-02-10T08:15:22.567Z",
  "done": false
}
```

**Status Code:** 400 Bad Request

**Response Body:**

```json
{
  "error": "Invalid request body"
}
```

# Authentication

## Login

### Request

`POST /login`

**Request Body:**

```json
{
  "username": "user123",
  "password": "password123"
}
```

### Response

**Status Code:** 202 Accepted

**Response Body:**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjY2YjQ3ZWY1ZTk1ZDJlYTQzZTk2NzIyMCIsInVzZXJuYW1lIjoiaGVsbG8ifQ.zNoJB8Dq1OB8hp31bvVJXnT0mIvdPZgvRu6H3HHa0Yg"
}
```

**Status Code:** 400 Bad Request

**Response Body:**

```json
{
  "error": "invalid credentials"
}
```

## Promote

### Request

`POST /promote`

**Request Body:**

```json
{
  "username": "user123"
}
```

### Response

**Status Code:** 202 Accepted

**Response Body:**

```json
{
  "message": "user promoted succesfuly"
}
```

**Status Code:** 400 Bad Request

**Response Body:**

```json
{
  "error": "user is already promoted"
}
```

**Status Code:** 400 Bad Request

**Response Body:**

```json
{
  "error": "user not found"
}
```

## Register

### Request

`POST /register`

**Request Body:**

```json
{
  "username": "user123",
  "password": "password123"
}
```

### Response

**Status Code:** 202 Accepted

**Response Body:**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjY2YjQ3ZWY1ZTk1ZDJlYTQzZTk2NzIyMCIsInVzZXJuYW1lIjoiaGVsbG8ifQ.zNoJB8Dq1OB8hp31bvVJXnT0mIvdPZgvRu6H3HHa0Yg"
}
```

**Status Code:** 400 Bad Request

**Response Body:**

```json
{
  "error": "user is already promoted"
}
```
