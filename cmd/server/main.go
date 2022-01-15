package main

import "fmt"

// App - the struct which contains things like pointers
// to database connections
type App struct{}

// Run - sets up application
func (app *App) Run() error {
	fmt.Println("Setting up application")
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
