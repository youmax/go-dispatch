package main

import (
	app "upay/app"
	"upay/routes"
)

func main() {
	app := app.CreateApplication()
	routes.SetRoutes(app)
	app.Run()
}
