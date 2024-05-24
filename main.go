package main

import (
	"github.com/dihanto/fiberplate/app"
	"github.com/dihanto/fiberplate/config"
)

func main() {
	db := config.Connect()

	app := app.New(db)

	app.Listen(":3000")
}
