package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/sc0rp1us/dotviz-server/controllers"

	"flag"
	"fmt"
)

var bindAddr string

func init() {
	host := flag.String("host", "0.0.0.0", "Set host for binding dotviz-server")
	port := flag.String("port", "1234", "Set port for binding dotviz-server")
	flag.Parse()
	bindAddr = fmt.Sprintf("%v:%v", *host, *port)
}

func main() {
	m := martini.Classic()
	m.Use(martini.Static("public"))
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
	}))

	dotController := new(controllers.DotController)
	m.Get("/dot/:format/:data", dotController.Render)

	dotVizController := new(controllers.DotVizController)
	m.Get("/", dotVizController.Render)
	m.Get("/dotviz", dotVizController.Render)
	m.Get("/dotviz/:data", dotVizController.Render)
	m.Post("/dotviz", dotVizController.New)

	log.Printf("Start listening: %v", bindAddr)
	log.Fatal(http.ListenAndServe(bindAddr, m))

}
