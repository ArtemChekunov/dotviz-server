package controllers

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/sc0rp1us/godotviz"
)

// DotController - Controller structure for all DOT functions
type DotController struct{}

// Render - Method for rendering DOT from base64 to image
func (dot *DotController) Render(res http.ResponseWriter, params martini.Params) {
	format := params["format"]
	data := params["data"]

	err := dot.checkFormat(format)
	if err != nil {
		log.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintln(err)))
	}

	graph, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: add error hendler
	result := godotviz.DotRender(string(graph), format)

	res.Header().Set("Content-Type", "image/"+format)
	res.Write(result)

}

// checkFormat - Check available formats for rendering
func (dot *DotController) checkFormat(ext string) (err error) {
	boolsupportedFormats := map[string]bool{
		"canon":     true,
		"cmap":      true,
		"cmapx":     true,
		"cmapx_np":  true,
		"dot":       true,
		"eps":       true,
		"fig":       true,
		"gd":        true,
		"gd2":       true,
		"gif":       true,
		"gv":        true,
		"imap":      true,
		"imap_np":   true,
		"ismap":     true,
		"jpe":       true,
		"jpeg":      true,
		"jpg":       true,
		"pdf":       true,
		"pic":       true,
		"plain":     true,
		"plain-ext": true,
		"png":       true,
		"pov":       true,
		"ps":        true,
		"ps2":       true,
		"svg":       true,
		"svgz":      true,
		"tk":        true,
		"vml":       true,
		"vmlz":      true,
		"vrml":      true,
		"wbmp":      true,
		"x11":       true,
		"xdot":      true,
		"xdot1.2":   true,
		"xdot1.4":   true,
		"xlib":      true,
	}

	if boolsupportedFormats[ext] != true {
		return fmt.Errorf("Error: extention %s is not supported.", ext)
	}

	return
}
