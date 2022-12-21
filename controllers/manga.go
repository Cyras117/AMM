package controllers

import (
	"html/template"
	"net/http"
)

var tmp = template.Must(template.ParseGlob("templates/*.html"))

func IndexHanddler(w http.ResponseWriter, r *http.Request) {
	tmp.ExecuteTemplate(w, "index", nil)
}
