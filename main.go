package main

import (
	app "github.com/amcereijo/go-country-api/app"
)

func main() {
	app := app.App{}
	app.Intialize(".env")
	app.Run(".env")
}
