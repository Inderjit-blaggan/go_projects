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
    git clone git@github.com:Inderjit-blaggan/go_projects.git
    cd go_projects/Todo_go_application
    ```

2. Install the required dependencies:

    ```sh
    go get github.com/gofiber/fiber/v2
    go get github.com/joho/godotenv
    ```

3. (Optional) Install `air` for live reloading during development:

    ```sh
    go install github.com/cosmtrek/air@latest
    ```
4. Set the custom port inside a `.env` file:
   - I have used 5000 as shown below
   
    ```sh
    PORT=5000
    ```

5. Run the application:

    ```sh
    go run main.go
    ```

    Or if you have `air` installed, run:

    ```sh
    air
    ```
![Screen Shot 2024-06-02 at 05 07 20 AM](https://github.com/Inderjit-blaggan/go_projects/assets/73047852/4f5f7d91-80fa-4ca7-8848-9d202a2d1b7a)

## Usage

Once the application is running, you can interact with the API using tools like `curl`, Postman, or your web browser.

### Example Requests

- **Get all todos**:

    ```sh
    curl http://localhost:PORT/api/todos
    ```

- **Create a new todo**:

    ```sh
    curl -X POST -H "Content-Type: application/json" -d '{"completed":false,"body":"Learn Go Fiber"}' http://localhost:PORT/api/todos
    ```

- **Update a todo**:

    ```sh
    curl -X PATCH -H "Content-Type: application/json" -d '{"completed":true,"body":"Learn Go Fiber Updated"}' http://localhost:PORT/api/todos/id
    ```

- **Delete a todo**:

    ```sh
    curl -X DELETE http://localhost:PORT/api/todos/id
    ```

## API Endpoints

### GET /todos

