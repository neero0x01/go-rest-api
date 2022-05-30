package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/neero0x01/go-rest-api/internal/transport/http"
)

// App - the struct which contains things like pointers to db connection
type App struct {}

func (app *App) Run() error {
	fmt.Println("Setting up the App")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

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