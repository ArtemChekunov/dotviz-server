package main

// import "github.com/sc0rp1us/godotviz"

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/sc0rp1us/godotviz"
)
import "encoding/base64"

func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		return "Hello world!!!!!!13333"
	})

	m.Get("/viz/:data/:format", func(res http.ResponseWriter, params martini.Params) {
		res.Header().Set("Content-Type", "image/png")
		graph, _ := base64.StdEncoding.DecodeString(params["data"])
		// return "asdasd " + params["data"] + " " + params["format"] + " " + string(result)
		result := godotviz.DotRender(string(graph), "png")
		res.Write(result)
	})

	m.Run()
}

// func main() {
// 	result := godotviz.DotRender("graph {a -- b a -- b  b -- a }", "png")
// 	os.Stdout.Write(result)
// }
// Stdout.Write(result)
// }
