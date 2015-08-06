package main

// import "github.com/sc0rp1us/godotviz"

import "github.com/go-martini/martini"
import "encoding/base64"

func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		return "Hello world!!!!!!13333"
	})

	m.Get("/viz/:data/:format", func(params martini.Params) string {
		result, _ := base64.StdEncoding.DecodeString(params["data"])
		return "asdasd " + params["data"] + " " + params["format"] + " " + string(result)
	})

	m.Run()
}

// func main() {
// 	result := godotviz.DotRender("graph {a -- b a -- b  b -- a }", "png")
// 	os.Stdout.Write(result)
// }
