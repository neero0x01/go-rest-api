package main

import "fmt"

// App - the struct which contains things like pointers to db connection
type App struct {}

func (app *App) Run() error {
	fmt.Println("Setting up the App")
	return nil
}

func main() {
	fmt.Println("GO REST API!")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up the REST API")
		fmt.Println(err)
	}
}