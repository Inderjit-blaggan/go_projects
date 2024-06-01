# Go Todo Application

Welcome to the Go Todo Application repository! This project is a simple todo list application built using the [Fiber](https://gofiber.io/) web framework for Go.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Installation

To get started with this project, you'll need to have Go installed on your machine. You can download it from [golang.org](https://golang.org/dl/).

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/go-todo-app.git
    cd go-todo-app
    ```

2. Install the required dependencies:

    ```sh
    go get github.com/gofiber/fiber/v2
    go get github.com/cosmtrek/air@latest
    ```

3. (Optional) Install `air` for live reloading during development:

    ```sh
    go install github.com/cosmtrek/air@latest
    ```

4. Run the application:

    ```sh
    go run main.go
    ```

    Or if you have `air` installed, run:

    ```sh
    air
    ```

## Usage

Once the application is running, you can interact with the API using tools like `curl`, Postman, or your web browser.

### Example Requests

- **Get all todos**:

    ```sh
    curl http://localhost:3000/todos
    ```

- **Create a new todo**:

    ```sh
    curl -X POST -H "Content-Type: application/json" -d '{"completed":false,"body":"Learn Go Fiber"}' http://localhost:3000/todos
    ```

- **Get a specific todo by ID**:

    ```sh
    curl http://localhost:3000/todos/1
    ```

- **Update a todo**:

    ```sh
    curl -X PUT -H "Content-Type: application/json" -d '{"completed":true,"body":"Learn Go Fiber Updated"}' http://localhost:3000/todos/1
    ```

- **Delete a todo**:

    ```sh
    curl -X DELETE http://localhost:3000/todos/1
    ```

## API Endpoints

### GET /todos

Retrieves a list of all todo items.

#### Response

```json
[
  {
    "id": 1,
    "completed": false,
    "body": "Learn Go"
  },
  {
    "id": 2,
    "completed": true,
    "body": "Learn Fiber"
  }
]
