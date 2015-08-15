package controllers

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
)

// DotVizController - Controller structure for all DotViz pages
type DotVizController struct{}

type dotVizResponse struct {
	DotBase64 string
	DotString string
}

func (dvr *dotVizResponse) FromString(dotString string) {
	dvr.DotBase64 = base64.StdEncoding.EncodeToString([]byte(dotString))
	dvr.DotString = dotString
}

func (dvr *dotVizResponse) FromBase64(dotBase64 string) {
	data, _ := base64.StdEncoding.DecodeString(dotBase64)
	dvr.DotString = string(data)
	dvr.DotBase64 = dotBase64
}

// Render - Method for rendering main html page
func (dotviz *DotVizController) Render(r render.Render, params martini.Params) {
	data := params["data"]
	dotvizresp := new(dotVizResponse)

	if data != "" {
		dotvizresp.FromBase64(data)
	} else {
		// dotvizresp.FromBase64("TXlTdHJpbmdXV1c=")
		// dotvizresp.FromBase64("Z3JhcGggewphIC0tIGIgYSAtLSBiICBiIC0tIGEgIHIgLS0gZwp9Cg==")
		dotvizresp.FromString("graph { a -- b }")
	}

	r.HTML(200, "dotviz", dotvizresp)
}

// New - Method for convertation POST DOT plaintext into base64
// and Redirect for rendering
func (dotviz *DotVizController) New(r render.Render, req *http.Request, res http.ResponseWriter) {
	dotString := req.FormValue("text")
	dotBase64 := base64.StdEncoding.EncodeToString([]byte(dotString))
	r.Redirect(fmt.Sprintf("/dotviz/%v", dotBase64))
}
