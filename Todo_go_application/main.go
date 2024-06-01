package main
import ( 
	"fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// This struct represents a todo item with three fields: ID, Completed, and Body. 
// The JSON tags specify how these fields should be named in JSON.
type Todo struct {
	ID          int     `json:"id"`
	Completed   bool    `json:"completed"`
	Body        string  `json:"body"`
}

func main(){
	fmt.Println("Starting the go todo app")
	// Create a new Fiber instance
    app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	// This line creates a new instance of the Todo struct.
	todos := []Todo{}

	// The code snippet you provided is an example of defining a route in a Fiber web application in Go. 
	// This particular route handles HTTP GET requests to the root URL ("/") and returns a JSON response with a status code of 200 and a message "hello world".
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	// Create a Todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		// &Todo{} initializes a Todo with its default zero values: ID is 0, Completed is false, and Body is an empty string "".
		todo := &Todo{}  

		// c.BodyParser(todo) reads the JSON from the request body and fills the todo struct with the corresponding data.
		// If parsing fails, an error is returned, and the handler terminates.
		if err := c.BodyParser(todo); err != nil {
			return err
		}
		
		// This line checks if the Body field of the todo struct is empty.
		// If todo.Body is an empty string, it returns a 400 Bad Request status with a JSON error message indicating that the todo body is required.
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{ "error": "Todo body is required"})
		}

		// This line sets the ID of the new todo item.
		// len(todos) + 1 gives the next ID value based on the current number of todos. This assumes todos is a slice that holds all existing todos.
		todo.ID = len(todos) + 1
		// This line appends the new todo item to the todos slice.
		// *todo dereferences the pointer to todo, adding the actual struct to the slice.
		todos = append(todos,*todo)

		return c.Status(201).JSON(todo)
	})
	
	// Update a Todo
	app.Patch("/api/todos/:id",func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo:= range todos{
			if fmt.Sprint(todo.ID) == id{
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{ "error": "Todo not found"})
	})

	// Delete a todo
	app.Delete("/api/todos/:id",func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo:= range todos{
			if fmt.Sprint(todo.ID) == id{
				todos = append(todos[:i],todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{ "success": "true"})
			}
		}
		return c.Status(404).JSON(fiber.Map{ "error": "Todo not found"})
	})

    // Start the server on port 4000
    log.Fatal(app.Listen(":"+PORT))
}