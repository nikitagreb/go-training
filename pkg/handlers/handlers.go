package handlers

import (
	"github.com/nikitagreb/go-training/pkg/render"
	"net/http"
)

// Home is the about page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
