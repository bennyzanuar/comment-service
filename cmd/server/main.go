package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/bennyzanuar/comment-service/internal/transport/http"
)

// App - the struct which contains things like pointers
// to database connections
type App struct{}

// Run - sets up application
func (app *App) Run() error {
	fmt.Println("Setting up application")
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8008", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go rest api")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up application")
		fmt.Println(err)
	}
}
