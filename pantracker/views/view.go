// Author: Richard Xiong

package views

import (
	"html/template"
	"log"
	"net/http"
)

// View struct is used to generate html and css views
type View struct {
	Template *template.Template
	Master   string
}

// NewView constructor creates a new view type
// which will be used by the controllers package
func NewView(filepath ...string) *View {

	filepaths := append(filepath, "pages/master.html")

	template, err := template.ParseFiles(filepaths...)
	if err != nil {
		log.Println(err)
	}

	return &View{
		Template: template,
		Master:   "master",
	}
}

// RenderTemplate is called to render a template for a specific view
func (v *View) RenderTemplate(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Master, data)
}

// ServeHTTP is implemented by View type from Handler interface
// This used the adapter desing pattern
func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.RenderTemplate(w, nil); err != nil {
		log.Println(err)
	}
}
