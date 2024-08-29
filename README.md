# GoTodo API

GoTodo is a simple task management API built using Go, Gin for routing, and GORM for database interactions. This project allows you to create, read, update, and delete (CRUD) tasks in a PostgreSQL database.

## Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Running the Project](#running-the-project)
- [Database Schema](#database-schema)
- [Contributing](#contributing)
- [License](#license)


## Installation

To get started with GoTodo, you need to have Go installed on your machine. You can download and install Go from [here](https://golang.org/dl/).

### Clone the repository:
```bash
git clone https://github.com/yourusername/gotodo.git
cd gotodo
```
### Install dependencies:
```bash
go mod tidy
```

This will install the necessary Go packages defined in `go.mod`.

### Set up PostgreSQL:

Ensure you have PostgreSQL installed and running. Create a new database:

```PostgreSQL
CREATE DATABASE gotodo;
```

### Create the .env file:
Create a `.env` file in the root directory with the following content:

```makefile
HOST=localhost
PORT=5432
USER=postgres
PASSWORD=yourpassword
DB_NAME=gotodo
```
Replace `yourpassword` with your PostgreSQL user password.

## Configuration

Database Configuration: The API uses environment variables for database connection settings. Ensure that your .env file contains the correct information for your PostgreSQL instance.

TimeZone: The database connection string includes TimeZone=Africa/Lagos. Modify it in the db.go file if you need to use a different timezone.

## Usage

### API Endpoints
Create a Task

POST /todos
Body:
```json
{
    "title": "Task Title",
    "description": "Task Description",
    "completed": false
}
```

Response:
201 Created: Task created successfully.
400 Bad Request: Invalid JSON or validation error.

Get All Tasks

GET /todos
    Response:
        - 200 OK: Returns a list of all tasks.
        - 500 Internal Server Error: Failed to retrieve tasks.

Get a Task by ID

    GET /todos/:id
    Response:
        - 200 OK: Returns the task with the specified ID.
        - 404 Not Found: Task with the specified ID does not exist.
        - 500 Internal Server Error: Failed to retrieve the task.

Update a Task

    PUT `/todos/:id`
        Body:
```json
    {
    "title": "Updated Title",
    "description": "Updated Description",
    "completed": true
    }
```

    Response:
        - 200 OK: Task updated successfully.
        - 400 Bad Request: Invalid JSON or task ID.
        - 500 Internal Server Error: Failed to update the task.

Delete a Task

#### DELETE /todos/:id
    Response:
        - 200 OK: Task deleted successfully.
        - 404 Not Found: Task with the specified ID does not exist.
        - 500 Internal Server Error: Failed to delete the task.

## Running the Project
Build and Run:

You can build and run the project using the following command:


```bash
go run cmd/main.go
```

Access the API:
The API will be available at http://localhost:8080. You can use tools like curl, Postman, or any other HTTP client to interact with the API.


## Database Schema


The PostgreSQL table used in this project is defined as:

```PostgreSQL
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN
);
```

## Contributing
Contributions are welcome! Please open an issue or submit a pull request if you find bugs or want to add new features.

## License
No license as of know.