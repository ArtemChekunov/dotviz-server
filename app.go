package main

import (
	"github.com/go-martini/martini"
	"github.com/sc0rp1us/dotviz-server/controllers"
)

func main() {
	m := martini.Classic()

	dotController := new(controllers.DotController)
	m.Get("/dot/:format/:data", dotController.Render)

	m.Run()
}
